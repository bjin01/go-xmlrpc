package mock

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	"reflect"
	"testing"
)

type mockClient struct {
	CallMock func(methodName string, args ...interface{}) (xmlrpc.Value, error)

	Testing *testing.T

	expectedMethodName    func(t *testing.T, actual string)
	expectedArgumentCount func(t *testing.T, actual int)
	expectedArguments     map[int]func(t *testing.T, actual interface{})
}

func (m *mockClient) Call(methodName string, args ...interface{}) (v xmlrpc.Value, err error) {
	if m.expectedMethodName != nil {
		m.expectedMethodName(m.Testing, methodName)
	}

	if m.expectedArgumentCount != nil {
		m.expectedArgumentCount(m.Testing, len(args))
	}

	if m.expectedArguments != nil {
		for i, f := range m.expectedArguments {
			f(m.Testing, args[i])
		}
	}

	if m.Testing.Failed() {
		m.Testing.FailNow()
	}

	return m.CallMock(methodName, args...)
}

func (m *mockClient) ExpectMethodName(expected string) {
	m.expectedMethodName = func(t *testing.T, actual string) {
		if actual != expected {
			t.Errorf("methodName == %q, want %q", actual, expected)
		}
	}
}

func (m *mockClient) ExpectArgumentCount(expected int) {
	m.expectedArgumentCount = func(t *testing.T, actual int) {
		if actual != expected {
			t.Errorf("len(args) == %v, want %v", actual, expected)
		}
	}
}

func (m *mockClient) ExpectArgument(index int, kind reflect.Kind, expected interface{}) {
	if m.expectedArguments == nil {
		m.expectedArguments = make(map[int]func(t *testing.T, actual interface{}))
	}

	m.expectedArguments[index] = func(t *testing.T, actual interface{}) {
		if reflect.ValueOf(actual).Kind() != kind || actual != expected {
			t.Errorf("args[%d] == %q, want %q", index, actual, expected)
		}
	}
}
