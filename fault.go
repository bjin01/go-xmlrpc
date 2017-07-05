package xmlrpc

type XMLRPCFault struct {
	message string
	code    int
}

func (e XMLRPCFault) Error() string {
	return e.message
}

func (e XMLRPCFault) FaultCode() int {
	return e.code
}

func (e XMLRPCFault) FaultString() string {
	return e.message
}
