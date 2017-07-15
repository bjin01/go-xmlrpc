package xmlrpc

// Member represents a member in a struct.
type Member interface {
	// Name returns the name of the struct's member.
	Name() string

	// Value returns the value of the struct's member.
	Value() Value
}
