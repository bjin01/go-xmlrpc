package xmlrpc

import "encoding/xml"

type member struct {
	XMLName  xml.Name `xml:"member"`
	NameTag  string   `xml:"name"`
	ValueTag value    `xml:"value"`
}

func (m member) Name() string {
	return m.NameTag
}

func (m member) Value() Value {
	return m.ValueTag
}
