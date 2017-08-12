package mock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
	"reflect"
	"testing"
	"time"
)

var _ = Describe("mock Client", func() {

	Context("nil Testing", func() {
		var (
			client *mock.Client
		)

		BeforeEach(func() {
			client = mock.NewClient(nil)
		})

		Context("setting func", func() {
			It("calls the CallMock", func() {
				value := mock.NewValue()

				client.CallMock = func(methodName string, args ...interface{}) (xmlrpc.Value, error) {
					Expect(methodName).To(Equal("test-method"))
					Expect(args).To(HaveLen(3))
					Expect(args).To(ConsistOf("foo", 42, true))

					return value, nil
				}

				v, err := client.Call("test-method", "foo", 42, true)

				Expect(err).ToNot(HaveOccurred())
				Expect(v).To(BeIdenticalTo(value))
			})
		})

		Context("using method", func() {
			Context("WithValue", func() {
				It("returns the Value", func() {
					value := mock.NewValue()

					client.WithValue(value)

					v, err := client.Call("test-method")

					Expect(err).ToNot(HaveOccurred())
					Expect(v).To(BeIdenticalTo(value))
				})
			})

			Context("WithError", func() {
				It("returns the error", func() {
					client.WithError(errors.New("test-error"))

					v, err := client.Call("test-method")

					Expect(err).To(HaveOccurred())
					Expect(v).To(BeNil())
				})
			})
		})

		It("won't assert expectations about the methodName", func() {
			client.ExpectMethodName("test-method-not-expected")

			client.WithValue(nil)

			Expect(client.Call("test-method")).To(BeNil())
		})

		It("won't assert expectations about the number of args", func() {
			client.ExpectArgumentCount(3)

			client.WithValue(nil)

			Expect(client.Call("test-method")).To(BeNil())
		})

		It("won't assert expectations about a specific arg", func() {
			client.ExpectArgument(0, reflect.String, "test-arg-not-expected")

			client.WithValue(nil)

			Expect(client.Call("test-method")).To(BeNil())
		})
	})

	Context("non-nil Testing", func() {
		var (
			client *mock.Client
		)

		BeforeEach(func() {
			client = mock.NewClient(&testing.T{})
		})

		Context("will assert expectations about the methodName", func() {
			Context("methodName matches", func() {
				It("will not modify Testing", func() {
					client.ExpectMethodName("test-method")

					client.WithValue(nil)

					Expect(client.Call("test-method")).To(BeNil())
					Expect(client.Testing.Failed()).To(BeFalse())
				})
			})

			Context("methodName doesn't match", func() {
				It("will modifiy Testing", func() {
					client.ExpectMethodName("test-method-not-expected")

					client.WithValue(nil)

					Expect(client.Call("test-method")).To(BeNil())
					Expect(client.Testing.Failed()).To(BeTrue())
				})
			})
		})

		Context("will assert expectations about the number of args", func() {
			Context("number of args matches", func() {
				It("will not modifiy Testing", func() {
					client.ExpectArgumentCount(3)

					client.WithValue(nil)

					Expect(client.Call("test-method", "foo", 42, true)).To(BeNil())
					Expect(client.Testing.Failed()).To(BeFalse())
				})
			})

			Context("number of args doesn't match", func() {
				It("will modifiy Testing", func() {
					client.ExpectArgumentCount(7)

					client.WithValue(nil)

					Expect(client.Call("test-method", "foo", 42, true)).To(BeNil())
					Expect(client.Testing.Failed()).To(BeTrue())
				})
			})
		})

		Context("will assert expectations about specific args", func() {
			Context("specific args match", func() {
				It("will not modifiy Testing", func() {
					client.ExpectArgument(0, reflect.String, "foo")
					client.ExpectArgument(1, reflect.Int, 42)
					client.ExpectArgument(2, reflect.Bool, true)

					client.WithValue(nil)

					Expect(client.Call("test-method", "foo", 42, true)).To(BeNil())
					Expect(client.Testing.Failed()).To(BeFalse())
				})
			})

			Context("specific args don't match", func() {
				It("will modify Testing", func() {
					client.ExpectArgument(0, reflect.String, "foo")
					client.ExpectArgument(1, reflect.Int, 42)
					client.ExpectArgument(2, reflect.Bool, true)

					client.WithValue(nil)

					Expect(client.Call("test-method", 8.25, time.Now())).To(BeNil())
					Expect(client.Testing.Failed()).To(BeTrue())
				})
			})
		})
	})
})
