package xmlrpc

import "strconv"

// Fault represents an error that occurred on the remote during
// an XML-RPC call.
type Fault struct {
	message string
	code    int
}

// Error returns the faults string representation.
func (e Fault) Error() string {
	return "XML-RPC fault (" + strconv.Itoa(e.Code()) + "): " + e.String()
}

// Code returns the code associated with the fault.
func (e Fault) Code() int {
	return e.code
}

// String returns the message associated with the fault.
func (e Fault) String() string {
	return e.message
}
