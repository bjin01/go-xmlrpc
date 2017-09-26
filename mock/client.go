package mock

import (
	"github.com/SHyx0rmZ/go-xmlrpc"
	"reflect"
	"testing"
)

// Client is a mock object for xmlrpc.Client. It can be used to fake XML-RPC
// calls in tests. If the client contains a non-nil instance of *testing.T, it
// will assert that any expectations about the call, that were made, hold true.
type Client struct {
	// CallMock will be invoked every time Call() is called.
	CallMock func(methodName string, args ...interface{}) (xmlrpc.Value, error)

	// Testing, if non-nill, will be used to assert that expectations about
	// the call hold true.
	Testing *testing.T

	expectedMethodName    func(t *testing.T, actual string)
	expectedArgumentCount func(t *testing.T, actual int)
	expectedArguments     map[int]func(t *testing.T, actual interface{})
}

// NewClient returns a new mock object for xmlrpc.Client. If t is non-nil, it
// will be used to assert that expectations about the call hold true. To get
// any use out of the mock Client, you must set CallMock directly or call
// either of WithValue() or WithError().
func NewClient(t *testing.T) *Client {
	return &Client{
		Testing: t,
	}
}

// Call invokes CallMock and returns its result. If the Testing is non-nil, it
// will also assert that any expectations about the call hold true.
func (m *Client) Call(methodName string, args ...interface{}) (v xmlrpc.Value, err error) {
	if m.Testing == nil {
		return m.CallMock(methodName, args...)
	}

	if m.expectedMethodName != nil {
		m.expectedMethodName(m.Testing, methodName)
	}

	if m.expectedArgumentCount != nil {
		m.expectedArgumentCount(m.Testing, len(args))
	}

	if m.expectedArguments != nil {
		for i, f := range m.expectedArguments {
			if len(args) <= i {
				m.Testing.Errorf("args[%d] is missing", i)

				break
			}

			f(m.Testing, args[i])
		}
	}

	return m.CallMock(methodName, args...)
}

// ExpectMethodName will assert that Call() is called with the expected
// methodName when Call() is called and Testing is non-nil.
func (m *Client) ExpectMethodName(expected string) {
	m.expectedMethodName = func(t *testing.T, actual string) {
		if actual != expected {
			t.Errorf("methodName == %q, want %q", actual, expected)
		}
	}
}

// ExpectArgumentCount will assert that Call() is called with the expected
// number of args when Call() is called and Testing is non-nil.
func (m *Client) ExpectArgumentCount(expected int) {
	m.expectedArgumentCount = func(t *testing.T, actual int) {
		if actual != expected {
			t.Errorf("len(args) == %v, want %v", actual, expected)
		}
	}
}

// ExpectArgument will assert that Call() is called with an arg of a specific
// kind and an expected value at a specific index in the args slice when
// Call() is called and Testing is non-nil.
func (m *Client) ExpectArgument(index int, kind reflect.Kind, expected interface{}) {
	if m.expectedArguments == nil {
		m.expectedArguments = make(map[int]func(t *testing.T, actual interface{}))
	}

	m.expectedArguments[index] = func(t *testing.T, actual interface{}) {
		if reflect.ValueOf(actual).Kind() != kind || actual != expected {
			t.Errorf("args[%d] == %q (%T), want %q (%s)", index, actual, actual, expected, kind.String())
		}
	}

	if kind == reflect.Slice {
		m.expectedArguments[index] = func(t *testing.T, actual interface{}) {
			actualValue := reflect.ValueOf(actual)
			expectedValue := reflect.ValueOf(expected)

			if actualValue.Kind() != kind || actualValue.Len() != expectedValue.Len() {
				t.Errorf("args[%d] == %#v, want %#v", index, actual, expected)

				return
			}

			for sliceIndex := 0; sliceIndex < expectedValue.Len(); sliceIndex++ {
				if actualValue.Index(sliceIndex) != expectedValue.Index(sliceIndex) {
					t.Errorf("args[%d] == %#v, want %#v", index, actual, expected)

					return
				}
			}
		}
	}
}

// WithValue assigns a func to CallMock. The func in CallMock will
// return (v, nil).
func (m *Client) WithValue(v xmlrpc.Value) *Client {
	m.CallMock = func(methodName string, args ...interface{}) (xmlrpc.Value, error) {
		return v, nil
	}
	return m
}

// WithError assigns a func to CallMock. The func in CallMock will
// return (nil, err).
func (m *Client) WithError(err error) *Client {
	m.CallMock = func(methodName string, args ...interface{}) (xmlrpc.Value, error) {
		return nil, err
	}
	return m
}
