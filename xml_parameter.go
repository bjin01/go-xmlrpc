package xmlrpc

import "encoding/xml"

type parameter struct {
	XMLName  xml.Name `xml:"param"`
	ValueTag value    `xml:"value"`
}
