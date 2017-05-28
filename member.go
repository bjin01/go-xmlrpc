package xmlrpc

type Member interface {
	Name() string
	Value() Value
}
