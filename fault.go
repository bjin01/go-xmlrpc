package xmlrpc

// XMLRPCFault represents an error that occurred on the remote during
// an XML-RPC call.
type XMLRPCFault struct {
	message string
	code    int
}

// Error returns the errors string representation.
func (e XMLRPCFault) Error() string {
	return e.message
}

// FaultCode returns the code associated with the error.
func (e XMLRPCFault) FaultCode() int {
	return e.code
}

// FaultString returns the message associated with the error.
func (e XMLRPCFault) FaultString() string {
	return e.message
}
