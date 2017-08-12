package mock

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	"time"
)

type mockValue struct {
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

func (m *mockValue) Values() []xmlrpc.Value   { return m.ValuesMock() }
func (m *mockValue) Bytes() []byte            { return m.BytesMock() }
func (m *mockValue) Bool() bool               { return m.BoolMock() }
func (m *mockValue) Time() time.Time          { return m.TimeMock() }
func (m *mockValue) Double() float64          { return m.DoubleMock() }
func (m *mockValue) Int() int                 { return m.IntMock() }
func (m *mockValue) String() string           { return m.StringMock() }
func (m *mockValue) Members() []xmlrpc.Member { return m.MembersMock() }
func (m *mockValue) Kind() xmlrpc.Kind        { return m.KindMock() }

func MockValueReturningValues(actual ...xmlrpc.Value) *mockValue {
	return &mockValue{
		KindMock:   func() xmlrpc.Kind { return xmlrpc.Array },
		ValuesMock: func() []xmlrpc.Value { return actual },
	}
}

func MockValueReturningBytes(actual []byte) *mockValue {
	return &mockValue{
		KindMock:  func() xmlrpc.Kind { return xmlrpc.Base64 },
		BytesMock: func() []byte { return actual },
	}
}

func MockValueReturningBool(actual bool) *mockValue {
	return &mockValue{
		KindMock: func() xmlrpc.Kind { return xmlrpc.Bool },
		BoolMock: func() bool { return actual },
	}
}

func MockValueReturningTime(actual time.Time) *mockValue {
	return &mockValue{
		KindMock: func() xmlrpc.Kind { return xmlrpc.DateTime },
		TimeMock: func() time.Time { return actual },
	}
}

func MockValueReturningDouble(actual float64) *mockValue {
	return &mockValue{
		KindMock:   func() xmlrpc.Kind { return xmlrpc.Double },
		DoubleMock: func() float64 { return actual },
	}
}

func MockValueReturningInt(actual int) *mockValue {
	return &mockValue{
		KindMock: func() xmlrpc.Kind { return xmlrpc.Int },
		IntMock:  func() int { return actual },
	}
}

func MockValueReturningString(actual string) *mockValue {
	return &mockValue{
		KindMock:   func() xmlrpc.Kind { return xmlrpc.String },
		StringMock: func() string { return actual },
	}
}

func MockValueReturningMembers(actual map[string]xmlrpc.Value) *mockValue {
	members := make([]xmlrpc.Member, 0)
	for name, value := range actual {
		members = append(members, &mockMember{
			NameMock:  func() string { return name },
			ValueMock: func() xmlrpc.Value { return value },
		})
	}
	return &mockValue{
		KindMock:    func() xmlrpc.Kind { return xmlrpc.Struct },
		MembersMock: func() []xmlrpc.Member { return members },
	}
}
