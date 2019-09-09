package mock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"time"
)

var _ = Describe("mock Value", func() {
	var (
		value *mock.Value
	)

	BeforeEach(func() {
		value = mock.NewValue()
	})

	Context("if unmodified", func() {
		It("returns an invalid Kind", func() {
			Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
		})
	})

	Context("if modified", func() {
		Context("setting func", func() {
			Context("ValuesMock", func() {
				It("calls the ValuesMock", func() {
					value1 := mock.NewValue()
					value2 := mock.NewValue()

					Expect(value1).ToNot(BeIdenticalTo(value2))

					value.ValuesMock = func() []xmlrpc.Value {
						return []xmlrpc.Value{
							value1,
							value2,
						}
					}

					Expect(value.Values()).To(HaveLen(2))
					Expect(value.Values()).To(ContainElement(value1))
					Expect(value.Values()).To(ContainElement(value2))
				})

				It("returns an invalid kind", func() {
					value.ValuesMock = func() []xmlrpc.Value {
						return []xmlrpc.Value{
							mock.NewValue(),
							mock.NewValue(),
						}
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("BytesMock", func() {
				It("calls the BytesMock", func() {
					bytes := []byte(`test-bytes`)

					value.BytesMock = func() []byte {
						return bytes
					}

					Expect(value.Bytes()).To(Equal(bytes))
				})

				It("returns an invalid Kind", func() {
					value.BytesMock = func() []byte {
						return []byte(`test-bytes`)
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("BoolMock", func() {
				It("calls the BoolMock", func() {
					value.BoolMock = func() bool {
						return true
					}

					Expect(value.Bool()).To(Equal(true))
				})

				It("returns an invalid Kind", func() {
					value.BoolMock = func() bool {
						return true
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("TimeMock", func() {
				It("calls the TimeMock", func() {
					now := time.Now()

					value.TimeMock = func() time.Time {
						return now
					}

					Expect(value.Time()).To(BeIdenticalTo(now))
				})

				It("returns an invalid Kind", func() {
					value.TimeMock = func() time.Time {
						return time.Now()
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("DoubleMock", func() {
				It("calls the DoubleMock", func() {
					value.DoubleMock = func() float64 {
						return 13.37
					}

					Expect(value.Double()).To(Equal(13.37))
				})

				It("returns an invalid Kind", func() {
					value.DoubleMock = func() float64 {
						return 13.37
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("IntMock", func() {
				It("calls the IntMock", func() {
					value.IntMock = func() int {
						return 42
					}

					Expect(value.Int()).To(Equal(42))
				})

				It("returns an invalid Kind", func() {
					value.IntMock = func() int {
						return 42
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("StringMock", func() {
				It("calls the StringMock", func() {
					value.StringMock = func() string {
						return "test-string"
					}

					Expect(value.Text()).To(Equal("test-string"))
				})

				It("returns an invalid Kind", func() {
					value.StringMock = func() string {
						return "test-string"
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})

			Context("MembersMock", func() {
				It("calls the MembersMock", func() {
					members := []xmlrpc.Member{
						mock.NewMember().WithName("foo").WithValue(mock.NewValue()),
						mock.NewMember().WithName("bar").WithValue(mock.NewValue()),
					}

					value.MembersMock = func() []xmlrpc.Member {
						return members
					}

					Expect(value.Members()).To(HaveLen(2))
					Expect(value.Members()).To(Equal(members))
				})

				It("returns an invalid Kind", func() {
					value.MembersMock = func() []xmlrpc.Member {
						return []xmlrpc.Member{
							mock.NewMember().WithName("foo").WithValue(mock.NewValue()),
							mock.NewMember().WithName("bar").WithValue(mock.NewValue()),
						}
					}

					Expect(value.Kind()).To(Equal(xmlrpc.Invalid))
				})
			})
		})

		Context("using method", func() {
			Context("WithValues", func() {
				It("returns Values", func() {
					value1 := mock.NewValue()
					value2 := mock.NewValue()

					Expect(value1).ToNot(BeIdenticalTo(value2))

					value.WithValues(value1, value2)

					Expect(value.Kind()).To(Equal(xmlrpc.Array))
					Expect(value.Values()).To(HaveLen(2))
					Expect(value.Values()).To(ContainElement(value1))
					Expect(value.Values()).To(ContainElement(value2))
				})
			})

			Context("WithBytes", func() {
				It("returns Bytes", func() {
					bytes := []byte(`test-bytes`)

					value.WithBytes(bytes)

					Expect(value.Kind()).To(Equal(xmlrpc.Base64))
					Expect(value.Bytes()).To(Equal(bytes))
				})
			})

			Context("WithBool", func() {
				It("returns Bool", func() {
					value.WithBool(true)

					Expect(value.Kind()).To(Equal(xmlrpc.Bool))
					Expect(value.Bool()).To(Equal(true))
				})
			})

			Context("WithTime", func() {
				It("returns Time", func() {
					now := time.Now()

					value.WithTime(now)

					Expect(value.Kind()).To(Equal(xmlrpc.DateTime))
					Expect(value.Time()).To(BeIdenticalTo(now))
				})
			})

			Context("WithDouble", func() {
				It("returns Double", func() {
					value.WithDouble(13.37)

					Expect(value.Kind()).To(Equal(xmlrpc.Double))
					Expect(value.Double()).To(Equal(13.37))
				})
			})

			Context("WithInt", func() {
				It("returns Int", func() {
					value.WithInt(42)

					Expect(value.Kind()).To(Equal(xmlrpc.Int))
					Expect(value.Int()).To(Equal(42))
				})
			})

			Context("WithString", func() {
				It("returns String", func() {
					value.WithString("test-string")

					Expect(value.Kind()).To(Equal(xmlrpc.String))
					Expect(value.Text()).To(Equal("test-string"))
				})
			})

			Context("WithMembers", func() {
				It("returns Members", func() {
					members := map[string]xmlrpc.Value{
						"foo": mock.NewValue(),
						"bar": mock.NewValue(),
					}

					value.WithMembers(members)

					Expect(value.Kind()).To(Equal(xmlrpc.Struct))
					Expect(value.Members()).To(HaveLen(2))
					Expect(value.Members()).To(ConsistOf(
						SatisfyAll(
							WithTransform((*mock.Member).Name, Equal("foo")),
							WithTransform((*mock.Member).Value, Equal(members["foo"])),
						),
						SatisfyAll(
							WithTransform((*mock.Member).Name, Equal("bar")),
							WithTransform((*mock.Member).Value, Equal(members["bar"])),
						),
					))
				})
			})
		})
	})
})
