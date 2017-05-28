package xmlrpc

import "encoding/xml"

type data struct {
	XMLName   xml.Name `xml:"data"`
	ValueTags []value  `xml:"value,omitempty"`
}
