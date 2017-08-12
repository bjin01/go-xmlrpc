package mock

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	"time"
)

type Value struct {
	ValuesMock  func() []xmlrpc.Value
	BytesMock   func() []byte
	BoolMock    func() bool
	TimeMock    func() time.Time
	DoubleMock  func() float64
	IntMock     func() int
	StringMock  func() string
	MembersMock func() []xmlrpc.Member
	KindMock    func() xmlrpc.Kind
}

func NewValue() *Value {
	return &Value{
		KindMock: func() xmlrpc.Kind { return xmlrpc.Invalid },
	}
}

func (m *Value) Values() []xmlrpc.Value   { return m.ValuesMock() }
func (m *Value) Bytes() []byte            { return m.BytesMock() }
func (m *Value) Bool() bool               { return m.BoolMock() }
func (m *Value) Time() time.Time          { return m.TimeMock() }
func (m *Value) Double() float64          { return m.DoubleMock() }
func (m *Value) Int() int                 { return m.IntMock() }
func (m *Value) String() string           { return m.StringMock() }
func (m *Value) Members() []xmlrpc.Member { return m.MembersMock() }
func (m *Value) Kind() xmlrpc.Kind        { return m.KindMock() }

func (m *Value) WithValues(actual ...xmlrpc.Value) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Array }
	m.ValuesMock = func() []xmlrpc.Value { return actual }
	return m
}

func (m *Value) WithBytes(actual []byte) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Base64 }
	m.BytesMock = func() []byte { return actual }
	return m
}

func (m *Value) WithBool(actual bool) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Bool }
	m.BoolMock = func() bool { return actual }
	return m
}

func (m *Value) WithTime(actual time.Time) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.DateTime }
	m.TimeMock = func() time.Time { return actual }
	return m
}

func (m *Value) WithDouble(actual float64) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Double }
	m.DoubleMock = func() float64 { return actual }
	return m
}

func (m *Value) WithInt(actual int) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Int }
	m.IntMock = func() int { return actual }
	return m
}

func (m *Value) WithString(actual string) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.String }
	m.StringMock = func() string { return actual }
	return m
}

func (m *Value) WithMembers(actual map[string]xmlrpc.Value) *Value {
	m.KindMock = func() xmlrpc.Kind { return xmlrpc.Struct }
	m.MembersMock = func() []xmlrpc.Member { return m.membersFromMap(actual) }
	return m
}

func (Value) membersFromMap(vs map[string]xmlrpc.Value) []xmlrpc.Member {
	m := make([]xmlrpc.Member, len(vs))
	for n, v := range vs {
		m = append(m, &Member{
			NameMock:  func() string { return n },
			ValueMock: func() xmlrpc.Value { return v },
		})
	}
	return m
}
