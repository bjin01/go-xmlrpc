package xmlrpc

import "time"

type Value interface {
	AsArray() []Value
	AsBytes() []byte
	AsBool() bool
	AsTime() time.Time
	AsDouble() float64
	AsInt() int
	AsString() string
	AsStruct() []Member
}
