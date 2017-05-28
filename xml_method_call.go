package xmlrpc

import "encoding/xml"

type methodCall struct {
	XMLName   xml.Name   `xml:"methodCall"`
	MethodTag string     `xml:"methodName"`
	ParamsTag parameters `xml:"params"`
}
