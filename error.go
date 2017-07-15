package xmlrpc

// Error represents an error that was caused either by some violation
// of the XML-RPC specification or by an assumption made by this implementation
// turning out to be wrong.
type Error struct {
	message string
}

// Error returns the errors string representation.
func (e Error) Error() string {
	return e.message
}
