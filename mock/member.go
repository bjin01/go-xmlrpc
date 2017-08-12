package mock

import "github.com/SHyx0rmZ/go-xmlrpc"

type Member struct {
	NameMock  func() string
	ValueMock func() xmlrpc.Value
}

func NewMember() *Member {
	m := &Member{}
	m.WithValue(NewValue())
	return m
}

func (m *Member) Name() string {
	return m.NameMock()
}

func (m *Member) Value() xmlrpc.Value {
	return m.ValueMock()
}

func (m *Member) WithName(value string) *Member {
	m.NameMock = func() string {
		return value
	}
	return m
}

func (m *Member) WithValue(value xmlrpc.Value) *Member {
	m.ValueMock = func() xmlrpc.Value {
		return value
	}
	return m
}
