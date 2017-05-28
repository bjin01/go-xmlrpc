package xmlrpc

import "encoding/xml"

type parameters struct {
	XMLName   xml.Name    `xml:"params"`
	ParamTags []parameter `xml:"param"`
}
