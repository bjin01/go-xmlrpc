package xmlrpc_test

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Value", func() {
	var (
		server           *ghttp.Server
		client           xmlrpc.Client
		verifyAndRespond = func(request, response string) {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyBody([]byte(request)),
					ghttp.RespondWith(200, []byte(response)),
				),
			)
		}
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = xmlrpc.NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Context("Encoding slices", func() {
		It("Can encode two strings", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><array><data><value><string>foo</string></value><value><string>bar</string></value></data></array></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><boolean>true</boolean></param></params></methodResponse>`,
			)

			_, err := client.Call("test", []interface{}{"foo", "bar"})

			Expect(err).To(BeNil())
		})
	})

	Context("Decoding slices", func() {
		It("Can decode two slices", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><array><data><value><string>foo</string></value><value><string>bar</string></value></data></array></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(len(val.AsArray())).To(Equal(2))
			Expect(val.AsArray()[0].AsString()).To(Equal("foo"))
			Expect(val.AsArray()[1].AsString()).To(Equal("bar"))
		})
	})

	Context("Encoding booleans", func() {
		It("Can encode true", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><boolean>true</boolean></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param></param></params></methodResponse>`,
			)

			_, err := client.Call("test", true)

			Expect(err).To(BeNil())
		})

		It("Can encode false", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><boolean>false</boolean></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param></param></params></methodResponse>`,
			)

			_, err := client.Call("test", false)

			Expect(err).To(BeNil())
		})
	})

	Context("Decoding booleans", func() {
		It("Can decode true", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsBool()).To(Equal(true))
		})

		It("Can decode false", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>false</boolean></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsBool()).To(Equal(false))
		})
	})
})
