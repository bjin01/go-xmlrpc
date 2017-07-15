package xmlrpc

// Member represents a member in a struct.
type Member interface {
	Name() string
	Value() Value
}
