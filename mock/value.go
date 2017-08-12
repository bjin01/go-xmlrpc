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

func (m *Value) Values() []xmlrpc.Value   { return m.ValuesMock() }
func (m *Value) Bytes() []byte            { return m.BytesMock() }
func (m *Value) Bool() bool               { return m.BoolMock() }
func (m *Value) Time() time.Time          { return m.TimeMock() }
func (m *Value) Double() float64          { return m.DoubleMock() }
func (m *Value) Int() int                 { return m.IntMock() }
func (m *Value) String() string           { return m.StringMock() }
func (m *Value) Members() []xmlrpc.Member { return m.MembersMock() }
func (m *Value) Kind() xmlrpc.Kind        { return m.KindMock() }

func ValueReturningValues(actual ...xmlrpc.Value) *Value {
	return &Value{
		KindMock:   func() xmlrpc.Kind { return xmlrpc.Array },
		ValuesMock: func() []xmlrpc.Value { return actual },
	}
}

func ValueReturningBytes(actual []byte) *Value {
	return &Value{
		KindMock:  func() xmlrpc.Kind { return xmlrpc.Base64 },
		BytesMock: func() []byte { return actual },
	}
}

func ValueReturningBool(actual bool) *Value {
	return &Value{
		KindMock: func() xmlrpc.Kind { return xmlrpc.Bool },
		BoolMock: func() bool { return actual },
	}
}

func ValueReturningTime(actual time.Time) *Value {
	return &Value{
		KindMock: func() xmlrpc.Kind { return xmlrpc.DateTime },
		TimeMock: func() time.Time { return actual },
	}
}

func ValueReturningDouble(actual float64) *Value {
	return &Value{
		KindMock:   func() xmlrpc.Kind { return xmlrpc.Double },
		DoubleMock: func() float64 { return actual },
	}
}

func ValueReturningInt(actual int) *Value {
	return &Value{
		KindMock: func() xmlrpc.Kind { return xmlrpc.Int },
		IntMock:  func() int { return actual },
	}
}

func ValueReturningString(actual string) *Value {
	return &Value{
		KindMock:   func() xmlrpc.Kind { return xmlrpc.String },
		StringMock: func() string { return actual },
	}
}

func ValueReturningMembers(actual map[string]xmlrpc.Value) *Value {
	m := make([]xmlrpc.Member, 0)
	for n, v := range actual {
		m = append(m, &Member{
			NameMock:  func() string { return n },
			ValueMock: func() xmlrpc.Value { return v },
		})
	}
	return &Value{
		KindMock:    func() xmlrpc.Kind { return xmlrpc.Struct },
		MembersMock: func() []xmlrpc.Member { return m },
	}
}
