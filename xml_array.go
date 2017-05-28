package xmlrpc

import "encoding/xml"

type array struct {
	XMLName xml.Name `xml:"array"`
	DataTag data     `xml:"data,omitempty"`
}
