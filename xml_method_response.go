package xmlrpc

import "encoding/xml"

type methodResponse struct {
	XMLName   xml.Name   `xml:"methodResponse"`
	ParamsTag parameters `xml:"params"`
}
