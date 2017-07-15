package xmlrpc

// XMLRPCError represents an error that was caused either by some violation
// of the XML-RPC specification or by an assumption made by this implementation
// turning out to be wrong.
type XMLRPCError struct {
	message string
}

// Error returns the errors string representation.
func (e XMLRPCError) Error() string {
	return e.message
}
