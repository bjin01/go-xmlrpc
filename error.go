package xmlrpc

type XMLRPCError struct {
	message string
}

func (e XMLRPCError) Error() string {
	return e.message
}
