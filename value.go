package xmlrpc

import "time"

// Kind represents the specific kind of value that a Value wraps.
// The zero Kind represents an invalid Value.
type Kind uint

const (
	Invalid Kind = iota
	Array
	Base64
	Bool
	DateTime
	Double
	Int
	String
	Struct
)

// Value is a wrapper around an actual XML-RPC value.
type Value interface {
	// Values returns the value's underlying value, as a slice.
	// It panics if the value's Kind is not Array.
	Values() []Value

	// Bytes returns the value's underlying value, as a slice of byte.
	// It panics if the value's Kind is not Base64.
	Bytes() []byte

	// Bool returns the value's underlying value, as a bool.
	// It panics if the value's Kind is not Bool.
	Bool() bool

	// Time returns the value's underlying value, as a time.Time.
	// It panics if the value's Kind is not DateTime.
	Time() time.Time

	// Double returns the value's underlying value, as a float64.
	// It panics if the value's Kind is not Double.
	Double() float64

	// Int returns the value's underlying value, as an int.
	// It panics if the value's Kind is not Int.
	Int() int

	// String returns the value's underlying value, as a string.
	// It panics if the value's Kind is not String.
	String() string

	// Members returns the value's underlying value, as a slice of Member.
	// It panics if the value's Kind is not Struct.
	Members() []Member

	// Kind returns the specific type of this value.
	Kind() Kind
}
