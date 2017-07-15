package xmlrpc

import (
	"encoding/xml"
	"strconv"
	"time"
)

type value struct {
	XMLName        xml.Name `xml:"value"`
	ArrayValueTags *[]value `xml:"array>data>value,omitempty"`
	Base64         []byte   `xml:"base64,omitempty"`
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
	Nil    string     `xml:"nil,omitempty"`
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
	return v.Base64
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

func (v value) AsNil() interface{} {
	return nil
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
