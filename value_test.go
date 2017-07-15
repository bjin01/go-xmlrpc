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

	Context("Encoding doubles", func() {
		It("Can encode 0.0", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><double>0</double></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", 0.0)

			Expect(err).To(BeNil())
		})

		It("Can encode 1337.42", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><double>1337.42</double></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", 1337.42)

			Expect(err).To(BeNil())
		})

		It("Can encode -1337.42", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><double>-1337.42</double></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", -1337.42)

			Expect(err).To(BeNil())
		})
	})

	Context("Decoding doubles", func() {
		It("Can decode 0.0", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><double>0.0</double></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsDouble()).To(Equal(0.0))
		})

		It("Can decode 1337.42", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><double>1337.42</double></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsDouble()).To(Equal(1337.42))
		})

		It("Can decode -1337.42", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><double>-1337.42</double></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsDouble()).To(Equal(-1337.42))
		})
	})

	Context("Encoding integers", func() {
		It("Can encode 0", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><int>0</int></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", 0)

			Expect(err).To(BeNil())
		})

		It("Can encode 1337", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><int>1337</int></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", 1337)

			Expect(err).To(BeNil())
		})

		It("Can encode -1337", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><int>-1337</int></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", -1337)

			Expect(err).To(BeNil())
		})
	})

	Context("Decoding integers", func() {
		Context("Wrapped in <int>", func() {
			It("Can decode 0", func() {
				verifyAndRespond(
					`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
					`<?xml version="1.0"?><methodResponse><params><param><value><int>0</int></value></param></params></methodResponse>`,
				)

				val, err := client.Call("test")

				Expect(err).To(BeNil())
				Expect(val.AsInt()).To(Equal(0))
			})

			It("Can decode 1337", func() {
				verifyAndRespond(
					`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
					`<?xml version="1.0"?><methodResponse><params><param><value><int>1337</int></value></param></params></methodResponse>`,
				)

				val, err := client.Call("test")

				Expect(err).To(BeNil())
				Expect(val.AsInt()).To(Equal(1337))
			})

			It("Can decode -1337", func() {
				verifyAndRespond(
					`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
					`<?xml version="1.0"?><methodResponse><params><param><value><int>-1337</int></value></param></params></methodResponse>`,
				)

				val, err := client.Call("test")

				Expect(err).To(BeNil())
				Expect(val.AsInt()).To(Equal(-1337))
			})
		})

		Context("Wrapped in <i4>", func() {
			It("Can decode 0", func() {
				verifyAndRespond(
					`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
					`<?xml version="1.0"?><methodResponse><params><param><value><i4>0</i4></value></param></params></methodResponse>`,
				)

				val, err := client.Call("test")

				Expect(err).To(BeNil())
				Expect(val.AsInt()).To(Equal(0))
			})

			It("Can decode 1337", func() {
				verifyAndRespond(
					`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
					`<?xml version="1.0"?><methodResponse><params><param><value><i4>1337</i4></value></param></params></methodResponse>`,
				)

				val, err := client.Call("test")

				Expect(err).To(BeNil())
				Expect(val.AsInt()).To(Equal(1337))
			})

			It("Can decode -1337", func() {
				verifyAndRespond(
					`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
					`<?xml version="1.0"?><methodResponse><params><param><value><i4>-1337</i4></value></param></params></methodResponse>`,
				)

				val, err := client.Call("test")

				Expect(err).To(BeNil())
				Expect(val.AsInt()).To(Equal(-1337))
			})
		})
	})

	Context("Encoding strings", func() {
		It("Can encode \"\"", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><string></string></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", "")

			Expect(err).To(BeNil())
		})

		It("Can encode \"Hello, world!\"", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><string>Hello, world!</string></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", "Hello, world!")

			Expect(err).To(BeNil())
		})
	})

	Context("Decoding strings", func() {
		It("Can decode \"\"", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><string></string></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsString()).To(Equal(""))
		})

		It("Can decode \"Hello, world!\"", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><string>Hello, world!</string></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(val.AsString()).To(Equal("Hello, world!"))
		})
	})

	Context("Encoding structs", func() {
		It("Can encode {\"foo\":\"bar\",\"answers\":[42]}", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params><param><value><struct><member><name>answers</name><value><array><data><value><int>42</int></value></data></array></value></member><member><name>foo</name><value><string>bar</string></value></member></struct></value></param></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><boolean>true</boolean></value></param></params></methodResponse>`,
			)

			_, err := client.Call("test", map[string]interface{}{"foo": "bar", "answers": []int{42}})

			Expect(err).To(BeNil())
		})
	})

	Context("Decoding structs", func() {
		It("Can decode {\"foo\":\"bar\",\"answers\":[42]}", func() {
			verifyAndRespond(
				`<?xml version="1.0"?><methodCall><methodName>test</methodName><params></params></methodCall>`,
				`<?xml version="1.0"?><methodResponse><params><param><value><struct><member><name>answers</name><value><array><data><value><int>42</int></value></data></array></value></member><member><name>foo</name><value><string>bar</string></value></member></struct></value></param></params></methodResponse>`,
			)

			val, err := client.Call("test")

			Expect(err).To(BeNil())
			Expect(len(val.AsStruct())).To(Equal(2))

			Expect(val.AsStruct()[0].Name()).To(Equal("answers"))
			Expect(len(val.AsStruct()[0].Value().AsArray())).To(Equal(1))
			Expect(val.AsStruct()[0].Value().AsArray()[0].AsInt()).To(Equal(42))
			Expect(val.AsStruct()[1].Name()).To(Equal("foo"))
			Expect(val.AsStruct()[1].Value().AsString()).To(Equal("bar"))
		})
	})
})
