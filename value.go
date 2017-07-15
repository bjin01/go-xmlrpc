package xmlrpc

import "time"

// Kind represents the specific kind of value that a Value wraps.
// The zero Kind represents an invalid Value.
type Kind uint

const (
	Invalid Kind = iota
	Array
	Base64
	Boolean
	DateTime
	Double
	Integer
	String
	Struct
)

// Value is a wrapper around an actual XML-RPC value.
type Value interface {
	// AsArray returns the value's underlying value, as a slice.
	// It panics if the value's Kind is not Array.
	AsArray() []Value

	// AsBytes returns the value's underlying value, as a slice of byte.
	// It panics if the value's Kind is not Base64.
	AsBytes() []byte

	// AsBool returns the value's underlying value, as a bool.
	// It panics if the value's Kind is not Boolean.
	AsBool() bool

	// AsTime returns the value's underlying value, as a time.Time.
	// It panics if the value's Kind is not DateTime.
	AsTime() time.Time

	// AsDouble returns the value's underlying value, as a float64.
	// It panics if the value's Kind is not Double.
	AsDouble() float64

	// AsInt returns the value's underlying value, as an int.
	// It panics if the value's Kind is not Integer.
	AsInt() int

	// AsString returns the value's underlying value, as a string.
	// It panics if the value's Kind is not String.
	AsString() string

	// AsStruct returns the value's underlying value, as a slice of Member.
	// It panics if the value's Kind is not Struct.
	AsStruct() []Member

	// Kind returns the specific type of this value.
	Kind() Kind
}
