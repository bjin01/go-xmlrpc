package xmlrpc

import (
	"encoding/base64"
	"encoding/xml"
	"strconv"
	"time"
)

type value struct {
	XMLName        xml.Name `xml:"value"`
	ArrayValueTags *[]value `xml:"array>data>value,omitempty"`
	Base64Tag      *string  `xml:"base64,omitempty"`
	BooleanTag     *bool    `xml:"boolean,omitempty"`
	DateTimeTag    string   `xml:"dateTime.iso8601,omitempty"`
	DoubleTag      *float64 `xml:"double,omitempty"`
	I4Tag          *struct {
		XML []byte `xml:",innerxml"`
	} `xml:"i4,omitempty"`
	IntTag *struct {
		XML []byte `xml:",innerxml"`
	} `xml:"int,omitempty"`
	StringTag *string    `xml:"string,omitempty"`
	StructTag *structure `xml:"struct,omitempty"`
}

func (v value) Values() []Value {
	if v.ArrayValueTags == nil {
		return []Value{}
	}

	values := make([]Value, 0, len(*v.ArrayValueTags))

	for _, value := range *v.ArrayValueTags {
		values = append(values, value)
	}

	return values
}

func (v value) Bytes() []byte {
	bytes, _ := base64.StdEncoding.DecodeString(*v.Base64Tag)

	return bytes
}

func (v value) Bool() bool {
	return *v.BooleanTag
}

func (v value) Time() time.Time {
	t, err := time.Parse(time.RFC3339, v.DateTimeTag)
	if err != nil {
		return time.Unix(0, 0)
	}

	return t
}

func (v value) Double() float64 {
	return *v.DoubleTag
}

func (v value) Int() int {
	if v.I4Tag != nil {
		i, err := strconv.Atoi(string(v.I4Tag.XML))
		if err == nil {
			return i
		}
	}

	if v.IntTag != nil {
		i, err := strconv.Atoi(string(v.IntTag.XML))
		if err == nil {
			return i
		}
	}

	return 0
}

func (v value) String() string {
	return *v.StringTag
}

func (v value) Members() []Member {
	if v.StructTag == nil {
		return []Member{}
	}

	members := make([]Member, 0, len(v.StructTag.MemberTags))

	for _, member := range v.StructTag.MemberTags {
		members = append(members, member)
	}

	return members
}

func (v value) Kind() Kind {
	kinds := []Kind{}

	if v.ArrayValueTags != nil {
		kinds = append(kinds, Array)
	}

	if v.Base64Tag != nil {
		kinds = append(kinds, Base64)
	}

	if v.BooleanTag != nil {
		kinds = append(kinds, Bool)
	}

	if v.DateTimeTag != "" {
		kinds = append(kinds, DateTime)
	}

	if v.DoubleTag != nil {
		kinds = append(kinds, Double)
	}

	if v.I4Tag != nil {
		kinds = append(kinds, Int)
	}

	if v.IntTag != nil {
		kinds = append(kinds, Int)
	}

	if v.StringTag != nil {
		kinds = append(kinds, String)
	}

	if v.StructTag != nil {
		kinds = append(kinds, Struct)
	}

	if len(kinds) == 1 {
		return kinds[0]
	}

	return Invalid
}
