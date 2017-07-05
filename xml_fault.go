package xmlrpc

import "encoding/xml"

type fault struct {
	XMLName  xml.Name `xml:"fault"`
	ValueTag value    `xml:"value"`
}
