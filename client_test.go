package xmlrpc_test

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"strconv"
)

var _ = Describe("XMLRPC client", func() {
	var server *ghttp.Server
	var client xmlrpc.Client

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = xmlrpc.NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	It("Can call a method and decode a response into a value", func() {
		request := `<?xml version="1.0"?><methodCall><methodName>foo</methodName><params></params></methodCall>`
		response := `<?xml version="1.0"?><methodResponse><params><param><value><string>bar</string></value></param></params></methodResponse>`

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/"),
				ghttp.VerifyHeaderKV("User-Agent", "Go-http-client/1.1"),
				ghttp.VerifyHeaderKV("Content-Type", "text/xml"),
				ghttp.VerifyHeaderKV("Content-Length", strconv.Itoa(len(request))),
				ghttp.VerifyBody([]byte(request)),
				ghttp.RespondWith(200, []byte(response)),
			),
		)

		val, err := client.Call("foo")

		Expect(err).To(BeNil())
		Expect(val.String()).To(Equal("bar"))
	})

	It("Can decode a fault", func() {
		request := `<?xml version="1.0"?><methodCall><methodName>invalid</methodName><params></params></methodCall>`
		response := `<?xml version="1.0"?><methodResponse><fault><value><struct><member><name>faultCode</name><value><int>42</int></value></member><member><name>faultString</name><value><string>Invalid method name</string></value></member></struct></value></fault></methodResponse>`

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyBody([]byte(request)),
				ghttp.RespondWith(200, []byte(response)),
			),
		)

		val, err := client.Call("invalid")

		Expect(err).ToNot(BeNil())
		Expect(val).To(BeNil())

		fault, typeAssertion := err.(*xmlrpc.XMLRPCFault)

		Expect(typeAssertion).To(BeTrue())
		Expect(fault.FaultCode()).To(Equal(42))
		Expect(fault.FaultString()).To(Equal("Invalid method name"))
	})
})
