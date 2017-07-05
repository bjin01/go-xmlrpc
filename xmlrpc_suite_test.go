package xmlrpc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestXMLRPC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XML-RPC Suite")
}
