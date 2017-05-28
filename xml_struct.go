package xmlrpc

import "encoding/xml"

type structure struct {
	XMLName    xml.Name `xml:"struct"`
	MemberTags []member `xml:"member"`
}
