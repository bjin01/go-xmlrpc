package mock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SHyx0rmZ/go-xmlrpc"
	"github.com/SHyx0rmZ/go-xmlrpc/mock"
)

var _ = Describe("Mock member", func() {
	var (
		member *mock.Member
	)

	BeforeEach(func() {
		member = mock.NewMember()
	})

	Context("if unmodified", func() {
		It("returns an empty name", func() {
			Expect(member.Name()).To(Equal(""))
		})

		It("returns an invalid value", func() {
			Expect(member.Value().Kind()).To(Equal(xmlrpc.Invalid))
		})
	})

	Context("if modified", func() {
		It("calls the NameMock", func() {
			member.NameMock = func() string {
				return "test-name"
			}

			Expect(member.Name()).To(Equal("test-name"))
		})

		It("calls the ValueMock", func() {
			value := mock.NewValue()

			member.ValueMock = func() xmlrpc.Value {
				return value
			}

			Expect(member.Value()).To(BeIdenticalTo(value))
		})
	})
})
