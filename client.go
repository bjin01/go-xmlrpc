package xmlrpc

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"time"
)

// A Client is an XML-RPC client.
type Client interface {
	// Call calls a remote method over XML-RPC using the specified arguments.
	// It will return the remote methods result or an error.
	Call(methodName string, args ...interface{}) (Value, error)
}

type client struct {
	client   *http.Client
	endpoint string
}

// NewClient instantiates a new XML-RPC client bound to the specified endpoint.
func NewClient(endpoint string) Client {
	return &client{&http.Client{}, endpoint}
}

func (c *client) values(args ...interface{}) ([]value, error) {
	results := make([]value, 0, len(args))

	for _, arg := range args {
		v := reflect.ValueOf(arg)
		switch v.Kind() {
		case reflect.Bool:
			ptr := new(bool)
			*ptr = v.Bool()
			results = append(results, value{BooleanTag: ptr})
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			ptr := &struct {
				XML []byte `xml:",innerxml"`
			}{}
			ptr.XML = []byte(strconv.Itoa(int(v.Int())))
			results = append(results, value{IntTag: ptr})
		case reflect.Float64:
			ptr := new(float64)
			*ptr = v.Float()
			results = append(results, value{DoubleTag: ptr})
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			if v.Type().Elem().Kind() == reflect.Uint8 {
				ptr := new(string)
				*ptr = base64.StdEncoding.EncodeToString(arg.([]byte))
				results = append(results, value{Base64Tag: ptr})
			} else {
				arguments := make([]interface{}, v.Len())
				for index := 0; index < v.Len(); index++ {
					arguments[index] = v.Index(index).Interface()
				}

				values, err := c.values(arguments...)
				if err != nil {
					return nil, err
				}

				results = append(results, value{ArrayValueTags: &values})
			}
		case reflect.String:
			ptr := new(string)
			*ptr = v.String()
			results = append(results, value{StringTag: ptr})
		case reflect.Map:
			members := make([]member, v.Len())
			keys := make([]string, v.Len())

			for index, key := range v.MapKeys() {
				if key.Kind() != reflect.String {
					return nil, &Error{"Invalid type " + v.Kind().String()}
				}

				keys[index] = key.String()
			}

			sort.Strings(keys)

			for index, key := range keys {
				values, err := c.values(v.MapIndex(reflect.ValueOf(key)).Interface())
				if err != nil {
					return nil, err
				}

				if len(values) != 1 {
					return nil, &Error{"Expected 1 element, got " + strconv.Itoa(len(values))}
				}

				members[index].NameTag = key
				members[index].ValueTag = values[0]
			}

			results = append(results, value{StructTag: &structure{MemberTags: members}})
		case reflect.Struct:
			if v.Type().PkgPath() != "time" || v.Type().Name() != "Time" {
				return nil, &Error{"Invalid type " + v.Kind().String()}
			}

			t := arg.(time.Time)

			results = append(results, value{DateTimeTag: t.Format(time.RFC3339)})
		default:
			return nil, &Error{"Invalid type " + v.Kind().String()}
		}
	}

	return results, nil
}

func (c *client) Call(methodName string, args ...interface{}) (Value, error) {
	methodCall := methodCall{
		MethodTag: methodName,
		ParamsTag: parameters{
			ParamTags: []parameter{},
		},
	}

	values, err := c.values(args...)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		methodCall.ParamsTag.ParamTags = append(methodCall.ParamsTag.ParamTags, parameter{ValueTag: value})
	}

	buffer := bytes.NewBuffer([]byte(`<?xml version="1.0"?>`))

	err = xml.NewEncoder(buffer).Encode(methodCall)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Post(c.endpoint, "text/xml", buffer)
	if err != nil {
		return nil, err
	}

	var methodResponse methodResponse

	err = xml.NewDecoder(response.Body).Decode(&methodResponse)
	if err != nil {
		return nil, err
	}

	if methodResponse.ParamsTag != nil && len(methodResponse.ParamsTag.ParamTags) != 1 {
		return nil, &Error{"Invalid amount of return values"}
	}

	if methodResponse.ParamsTag != nil && methodResponse.FaultTag == nil {
		return methodResponse.ParamsTag.ParamTags[0].ValueTag, nil
	}

	if methodResponse.FaultTag != nil && methodResponse.ParamsTag == nil {
		members := make(map[string]Value)

		for _, member := range methodResponse.FaultTag.ValueTag.StructTag.MemberTags {
			members[member.NameTag] = member.ValueTag
		}

		return nil, &Fault{
			message: members["faultString"].String(),
			code:    members["faultCode"].Int(),
		}
	}

	return nil, &Error{"Invalid amount of return values"}

}
