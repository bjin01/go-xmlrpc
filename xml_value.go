package xmlrpc

import (
	"encoding/xml"
	"strconv"
	"time"
	"encoding/base64"
)

type value struct {
	XMLName        xml.Name `xml:"value"`
	ArrayValueTags *[]value `xml:"array>data>value,omitempty"`
	Base64         *string   `xml:"base64,omitempty"`
	Boolean        *bool    `xml:"boolean,omitempty"`
	DateTime       string   `xml:"dateTime.iso8601,omitempty"`
	Double         *float64 `xml:"double,omitempty"`
	I4             *struct {
		XML []byte `xml:",innerxml"`
	} `xml:"i4,omitempty"`
	Int *struct {
		XML []byte `xml:",innerxml"`
	} `xml:"int,omitempty"`
	String *string    `xml:"string,omitempty"`
	Struct *structure `xml:"struct,omitempty"`
}

func (v value) AsArray() []Value {
	if v.ArrayValueTags == nil {
		return []Value{}
	}

	values := make([]Value, 0, len(*v.ArrayValueTags))

	for _, value := range *v.ArrayValueTags {
		values = append(values, value)
	}

	return values
}

func (v value) AsBytes() []byte {
	bytes, _ := base64.StdEncoding.DecodeString(*v.Base64)

	return bytes
}

func (v value) AsBool() bool {
	return *v.Boolean
}

func (v value) AsTime() time.Time {
	t, err := time.Parse(time.RFC3339, v.DateTime)
	if err != nil {
		return time.Unix(0, 0)
	}

	return t
}

func (v value) AsDouble() float64 {
	return *v.Double
}

func (v value) AsInt() int {
	if v.I4 != nil {
		i, err := strconv.Atoi(string(v.I4.XML))
		if err == nil {
			return i
		}
	}

	if v.Int != nil {
		i, err := strconv.Atoi(string(v.Int.XML))
		if err == nil {
			return i
		}
	}

	return 0
}

func (v value) AsString() string {
	return *v.String
}

func (v value) AsStruct() []Member {
	if v.Struct == nil {
		return []Member{}
	}

	members := make([]Member, 0, len(v.Struct.MemberTags))

	for _, member := range v.Struct.MemberTags {
		members = append(members, member)
	}

	return members
}

func (v value) Kind() Kind {
	kinds := []Kind{}

	if v.ArrayValueTags != nil {
		kinds = append(kinds, Array)
	}

	if v.Base64 != nil {
		kinds = append(kinds, Base64)
	}

	if v.Boolean != nil {
		kinds = append(kinds, Boolean)
	}

	if v.DateTime != "" {
		kinds = append(kinds, DateTime)
	}

	if v.Double != nil {
		kinds = append(kinds, Double)
	}

	if v.I4 != nil {
		kinds = append(kinds, Integer)
	}

	if v.Int != nil {
		kinds = append(kinds, Integer)
	}

	if v.String != nil {
		kinds = append(kinds, String)
	}

	if v.Struct != nil {
		kinds = append(kinds, Struct)
	}

	if len(kinds) == 1 {
		return kinds[0]
	}

	return Invalid
}
