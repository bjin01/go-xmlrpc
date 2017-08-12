# mock

    import "github.com/SHyx0rmZ/go-xmlrpc/mock"


## Usage

#### type Client

```go
type Client struct {
	// CallMock will be invoked every time Call() is called.
	CallMock func(methodName string, args ...interface{}) (xmlrpc.Value, error)

	// Testing, if non-nill, will be used to assert that expectations about
	// the call hold true.
	Testing *testing.T
}
```

Client is a mock object for xmlrpc.Client. It can be used to fake XML-RPC calls
in tests. If the client contains a non-nil instance of *testing.T, it will
assert that any expectations about the call, that were made, hold true.

#### func  NewClient

```go
func NewClient(t *testing.T) *Client
```
NewClient returns a new mock object for xmlrpc.Client. If t is non-nil, it will
be used to assert that expectations about the call hold true. To get any use out
of the mock Client, you must set CallMock directly or call either of WithValue()
or WithError().

#### func (*Client) Call

```go
func (m *Client) Call(methodName string, args ...interface{}) (v xmlrpc.Value, err error)
```
Call invokes CallMock and returns its result. If the Testing is non-nil, it will
also assert that any expectations about the call hold true.

#### func (*Client) ExpectArgument

```go
func (m *Client) ExpectArgument(index int, kind reflect.Kind, expected interface{})
```
ExpectArgument will assert that Call() is called with an arg of a specific kind
and an expected value at a specific index in the args slice when Call() is
called and Testing is non-nil.

#### func (*Client) ExpectArgumentCount

```go
func (m *Client) ExpectArgumentCount(expected int)
```
ExpectArgumentCount will assert that Call() is called with the expected number
of args when Call() is called and Testing is non-nil.

#### func (*Client) ExpectMethodName

```go
func (m *Client) ExpectMethodName(expected string)
```
ExpectMethodName will assert that Call() is called with the expected methodName
when Call() is called and Testing is non-nil.

#### func (*Client) WithError

```go
func (m *Client) WithError(err error) *Client
```
WithError assigns a func to CallMock. The func in CallMock will return (nil,
err).

#### func (*Client) WithValue

```go
func (m *Client) WithValue(v xmlrpc.Value) *Client
```
WithValue assigns a func to CallMock. The func in CallMock will return (v, nil).

#### type Member

```go
type Member struct {
	// NameMock will be invoked every time Name() is called.
	NameMock func() string
	// ValueMock will be invoked every time Value() is called.
	ValueMock func() xmlrpc.Value
}
```

Member is a mock object for xmlrpc.Member.

#### func  NewMember

```go
func NewMember() *Member
```
NewMember returns a new mock object for xmlrpc.Member, which will act as if it
did have an empty name and contained an invalid xmlrpc.Value.

#### func (*Member) Name

```go
func (m *Member) Name() string
```
Name invokes NameMock and returns its result.

#### func (*Member) Value

```go
func (m *Member) Value() xmlrpc.Value
```
Value invokes ValueMock and returns its result.

#### func (*Member) WithName

```go
func (m *Member) WithName(v string) *Member
```
WithName assigns a func to NameMock. The func in NameMock will return v.

#### func (*Member) WithValue

```go
func (m *Member) WithValue(v xmlrpc.Value) *Member
```
WithValue assigns a func to ValueMock. The func in ValueMock will return v.

#### type Value

```go
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
```

Value is a mock object for xmlrpc.Value.

#### func  NewValue

```go
func NewValue() *Value
```
NewValue returns a new mock object for xmlrpc.Value, which will act as if it
were of an invalid kind.

#### func (*Value) Bool

```go
func (m *Value) Bool() bool
```
Bool invokes BoolMock and returns its result.

#### func (*Value) Bytes

```go
func (m *Value) Bytes() []byte
```
Bytes invokes BytesMock and returns its result.

#### func (*Value) Double

```go
func (m *Value) Double() float64
```
Double invokes DoubleMock and returns its result.

#### func (*Value) Int

```go
func (m *Value) Int() int
```
Int invokes IntMock and returns its result.

#### func (*Value) Kind

```go
func (m *Value) Kind() xmlrpc.Kind
```
Kind invokes KindMock and returns its result.

#### func (*Value) Members

```go
func (m *Value) Members() []xmlrpc.Member
```
Members invokes MembersMock and returns its result.

#### func (*Value) String

```go
func (m *Value) String() string
```
String invokes StringMock and returns its result.

#### func (*Value) Time

```go
func (m *Value) Time() time.Time
```
Time invokes TimeMock and returns its result.

#### func (*Value) Values

```go
func (m *Value) Values() []xmlrpc.Value
```
Values invokes ValuesMock and returns its result.

#### func (*Value) WithBool

```go
func (m *Value) WithBool(v bool) *Value
```
WithBool assigns funcs to KindMock and BoolMock. The func in KindMock will
return xmlrpc.Bool, whereas the func in BoolMock will return v.

#### func (*Value) WithBytes

```go
func (m *Value) WithBytes(v []byte) *Value
```
WithBytes assigns funcs to KindMock and BytesMock. The func in KindMock will
return xmlrpc.Base64, whereas the func in BytesMock will return v.

#### func (*Value) WithDouble

```go
func (m *Value) WithDouble(v float64) *Value
```
WithDouble assigns funcs to KindMock and DoubleMock. The func in KindMock will
return xmlrpc.Double, whereas the func in DoubleMock will return v.

#### func (*Value) WithInt

```go
func (m *Value) WithInt(v int) *Value
```
WithInt assigns funcs to KindMock and IntMock. The func in KindMock will return
xmlrpc.Int, whereas the func in IntMock will return v.

#### func (*Value) WithMembers

```go
func (m *Value) WithMembers(v map[string]xmlrpc.Value) *Value
```
WithMembers assigns funcs to KindMock and MembersMock. The func in KindMock will
return xmlrpc.Struct, whereas the func in MembersMock will return v.

#### func (*Value) WithString

```go
func (m *Value) WithString(v string) *Value
```
WithString assigns funcs to KindMock and StringMock. The func in KindMock will
return xmlrpc.String, whereas the func in StringMock will return v.

#### func (*Value) WithTime

```go
func (m *Value) WithTime(v time.Time) *Value
```
WithTime assigns funcs to KindMock and TimeMock. The func in KindMock will
return xmlrpc.DateTime, whereas the func in TimeMock will return v.

#### func (*Value) WithValues

```go
func (m *Value) WithValues(v ...xmlrpc.Value) *Value
```
WithValues assigns funcs to KindMock and ValuesMock. The func in KindMock will
return xmlrpc.Array, whereas the func in ValuesMock will return v.
