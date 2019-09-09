# go-xmlrpc

    import xmlrpc "github.com/SHyx0rmZ/go-xmlrpc"


## Usage

#### type Client

```go
type Client interface {
	// Call calls a remote method over XML-RPC using the specified arguments.
	// It will return the remote methods result or an error.
	//
	// args is list of arguments to be passed to the remote method. Each element
	// of args is a single arg, where arg is either one of:
	//     - bool
	//     - int
	//     - int8
	//     - int16
	//     - int32
	//     - float32
	//     - float64
	//     - string
	//     - []byte
	//     - []arg
	//     - map[string]arg
	//     - time.Time
	Call(methodName string, args ...interface{}) (Value, error)
}
```

A Client is an XML-RPC client.

#### func  NewClient

```go
func NewClient(endpoint string) Client
```
NewClient instantiates a new XML-RPC client bound to the specified endpoint.

#### type Error

```go
type Error struct {
}
```

Error represents an error that was caused either by some violation of the
XML-RPC specification or by an assumption made by this implementation turning
out to be wrong.

#### func (Error) Error

```go
func (e Error) Error() string
```
Error returns the errors string representation.

#### type Fault

```go
type Fault struct {
}
```

Fault represents an error that occurred on the remote during an XML-RPC call.

#### func (Fault) Code

```go
func (e Fault) Code() int
```
Code returns the code associated with the fault.

#### func (Fault) Error

```go
func (e Fault) Error() string
```
Error returns the faults string representation.

#### func (Fault) String

```go
func (e Fault) String() string
```
String returns the message associated with the fault.

#### type Kind

```go
type Kind uint
```

Kind represents the specific kind of value that a Value wraps. The zero Kind
represents an invalid Value.

```go
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
```

#### type Member

```go
type Member interface {
	// Name returns the name of the struct's member.
	Name() string

	// Value returns the value of the struct's member.
	Value() Value
}
```

Member represents a member in a struct.

#### type Value

```go
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

	// Text returns the value's underlying value, as a string.
	// It panics if the value's Kind is not String.
	Text() string

	// Members returns the value's underlying value, as a slice of Member.
	// It panics if the value's Kind is not Struct.
	Members() []Member

	// Kind returns the specific type of this value.
	Kind() Kind
}
```

Value is a wrapper around an actual XML-RPC value.
