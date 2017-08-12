package mock

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	"time"
)

// Value is a mock object for xmlrpc.Value.
type Value struct {
	// ValuesMock will be invoked every time Values() is called.
	ValuesMock func() []xmlrpc.Value

	// BytesMock will be invoked every time Bytes() is called.
	BytesMock func() []byte

	// BoolMock will be invoked every time Bool() is called.
	BoolMock func() bool

	// TimeMock will be invoked every time Time() is called.
	TimeMock func() time.Time

	// DoubleMock will be invoked every time Double() is called.
	DoubleMock func() float64

	// IntMock will be invoked every time Int() is called.
	IntMock func() int

	// StringMock will be invoked every time String() is called.
	StringMock func() string

	// MembersMock will be invoked every time Members() is called.
	MembersMock func() []xmlrpc.Member

	// KindMock will be invoked every time Kind() is called.
	KindMock func() xmlrpc.Kind
}

// NewValue returns a new mock object for xmlrpc.Value, which will act as if
// it were of an invalid kind.
func NewValue() *Value {
	return &Value{
		KindMock: func() xmlrpc.Kind { return xmlrpc.Invalid },
	}
}

// Values invokes ValuesMock and returns its result.
func (m *Value) Values() []xmlrpc.Value { return m.ValuesMock() }

// Bytes invokes BytesMock and returns its result.
func (m *Value) Bytes() []byte { return m.BytesMock() }

// Bool invokes BoolMock and returns its result.
func (m *Value) Bool() bool { return m.BoolMock() }

// Time invokes TimeMock and returns its result.
func (m *Value) Time() time.Time { return m.TimeMock() }

// Double invokes DoubleMock and returns its result.
func (m *Value) Double() float64 { return m.DoubleMock() }

// Int invokes IntMock and returns its result.
func (m *Value) Int() int { return m.IntMock() }

// String invokes StringMock and returns its result.
func (m *Value) String() string { return m.StringMock() }

// Members invokes MembersMock and returns its result.
func (m *Value) Members() []xmlrpc.Member { return m.MembersMock() }

// Kind invokes KindMock and returns its result.
func (m *Value) Kind() xmlrpc.Kind { return m.KindMock() }

// WithValues assigns funcs to KindMock and ValuesMock. The func in KindMock
// will return xmlrpc.Array, whereas the func in ValuesMock will return v.
func (m *Value) WithValues(v ...xmlrpc.Value) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Array }
	m.ValuesMock = func() []xmlrpc.Value { return v }
	return m
}

// WithBytes assigns funcs to KindMock and BytesMock. The func in KindMock
// will return xmlrpc.Base64, whereas the func in BytesMock will return v.
func (m *Value) WithBytes(v []byte) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Base64 }
	m.BytesMock = func() []byte { return v }
	return m
}

// WithBool assigns funcs to KindMock and BoolMock. The func in KindMock will
// return xmlrpc.Bool, whereas the func in BoolMock will return v.
func (m *Value) WithBool(v bool) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Bool }
	m.BoolMock = func() bool { return v }
	return m
}

// WithTime assigns funcs to KindMock and TimeMock. The func in KindMock will
// return xmlrpc.DateTime, whereas the func in TimeMock will return v.
func (m *Value) WithTime(v time.Time) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.DateTime }
	m.TimeMock = func() time.Time { return v }
	return m
}

// WithDouble assigns funcs to KindMock and DoubleMock. The func in KindMock
// will return xmlrpc.Double, whereas the func in DoubleMock will return v.
func (m *Value) WithDouble(v float64) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Double }
	m.DoubleMock = func() float64 { return v }
	return m
}

// WithInt assigns funcs to KindMock and IntMock. The func in KindMock will
// return xmlrpc.Int, whereas the func in IntMock will return v.
func (m *Value) WithInt(v int) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Int }
	m.IntMock = func() int { return v }
	return m
}

// WithString assigns funcs to KindMock and StringMock. The func in KindMock
// will return xmlrpc.String, whereas the func in StringMock will return v.
func (m *Value) WithString(v string) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.String }
	m.StringMock = func() string { return v }
	return m
}

// WithMembers assigns funcs to KindMock and MembersMock. The func in KindMock
// will return xmlrpc.Struct, whereas the func in MembersMock will return v.
func (m *Value) WithMembers(v map[string]xmlrpc.Value) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Struct }
	m.MembersMock = func() []xmlrpc.Member { return m.membersFromMap(v) }
	return m
}

func (Value) membersFromMap(vs map[string]xmlrpc.Value) []xmlrpc.Member {
	m := make([]xmlrpc.Member, 0, len(vs))
	for n, v := range vs {
		m = append(m, NewMember().WithName(n).WithValue(v))
	}
	return m
}
