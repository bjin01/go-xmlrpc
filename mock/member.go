package mock

import "github.com/SHyx0rmZ/go-xmlrpc"

type Member struct {
	NameMock  func() string
	ValueMock func() xmlrpc.Value
}

func (m *Member) Name() string {
	return m.NameMock()
}

func (m *Member) Value() xmlrpc.Value {
	return m.ValueMock()
}
