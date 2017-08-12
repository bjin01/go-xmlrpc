package mock

import "github.com/SHyx0rmZ/go-xmlrpc"

type mockMember struct {
	NameMock  func() string
	ValueMock func() xmlrpc.Value
}

func (m *mockMember) Name() string {
	return m.NameMock()
}

func (m *mockMember) Value() xmlrpc.Value {
	return m.ValueMock()
}
