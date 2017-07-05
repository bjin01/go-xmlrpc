package xmlrpc

import "encoding/xml"

type methodResponse struct {
	XMLName   xml.Name    `xml:"methodResponse"`
	ParamsTag *parameters `xml:"params,omitempty"`
	FaultTag  *fault      `xml:"fault,omitempty"`
}
