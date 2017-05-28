package xmlrpc

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"reflect"
)

type Client interface {
	Call(methodName string, args ...interface{}) (Value, error)
}

type client struct {
	client   *http.Client
	endpoint string
}

func NewClient(endpoint string) Client {
	return &client{&http.Client{}, endpoint}
}

func (c *client) values(args ...interface{}) ([]value, error) {
	results := make([]value, 0, len(args))

	for _, arg := range args {
		v := reflect.ValueOf(arg)
		switch v.Kind() {
		case reflect.Bool:
			results = append(results, value{Boolean: v.Bool()})
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			results = append(results, value{Int: int(v.Int())})
		case reflect.Float64:
			results = append(results, value{Double: v.Float()})
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			values, err := c.values(arg)
			if err != nil {
				return nil, err
			}

			results = append(results, value{Array: &array{DataTag: data{ValueTags: values}}})
		case reflect.String:
			results = append(results, value{String: v.String()})
		default:
			return nil, &XMLRPCError{"Invalid type " + v.Kind().String()}
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

	buffer := bytes.NewBuffer([]byte{})

	err = xml.NewEncoder(buffer).Encode(methodCall)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Post(c.endpoint, "text/xml; charset=utf-8", buffer)
	if err != nil {
		return nil, err
	}

	var methodResponse methodResponse

	err = xml.NewDecoder(response.Body).Decode(&methodResponse)
	if err != nil {
		return nil, err
	}

	if len(methodResponse.ParamsTag.ParamTags) != 1 {
		return nil, &XMLRPCError{"Invalid amount of return values"}
	}

	return methodResponse.ParamsTag.ParamTags[0].ValueTag, nil
}
