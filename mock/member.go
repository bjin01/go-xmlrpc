package mock

import "github.com/SHyx0rmZ/go-xmlrpc"

// Member is a mock object for xmlrpc.Member.
type Member struct {
	// NameMock will be invoked every time Name() is called.
	NameMock func() string
	// ValueMock will be invoked every time Value() is called.
	ValueMock func() xmlrpc.Value
}

// NewMember returns a new mock object for xmlrpc.Member, which will act as if
// it did have an empty name and contained an invalid xmlrpc.Value.
func NewMember() *Member {
	m := &Member{}
	m.WithName("")
	m.WithValue(NewValue())
	return m
}

// Name invokes NameMock and returns its result.
func (m *Member) Name() string {
	return m.NameMock()
}

// Value invokes ValueMock and returns its result.
func (m *Member) Value() xmlrpc.Value {
	return m.ValueMock()
}

// WithName assigns a func to NameMock. The func in NameMock will return v.
func (m *Member) WithName(v string) *Member {
	m.NameMock = func() string {
		return v
	}
	return m
}

// WithValue assigns a func to ValueMock. The func in ValueMock will return v.
func (m *Member) WithValue(v xmlrpc.Value) *Member {
	m.ValueMock = func() xmlrpc.Value {
		return v
	}
	return m
}
