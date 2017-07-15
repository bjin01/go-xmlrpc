package xmlrpc

import "time"

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

type Value interface {
	AsArray() []Value
	AsBytes() []byte
	AsBool() bool
	AsTime() time.Time
	AsDouble() float64
	AsInt() int
	AsString() string
	AsStruct() []Member
	Kind() Kind
}
