package redcon

// Type of RESP
type Type byte

// Various RESP kinds
const (
	Integer = ':'
	String  = '+'
	Bulk    = '$'
	Array   = '*'
	Error   = '-'
)

type RESP struct {
	Type  Type
	Raw   []byte
	Data  []byte
	Count int
}

// ForEach iterates over each Array element
func (r RESP) ForEach(iter func(resp RESP) bool) { _ = "STUB: not implemented"; return }

func (r RESP) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (r RESP) String() string { _ = "STUB: not implemented"; return "" }

func (r RESP) Int() int64 { _ = "STUB: not implemented"; return 0 }

func (r RESP) Float() float64 { _ = "STUB: not implemented"; return 0 }

// Map returns a key/value map of an Array.
// The receiver RESP must be an Array with an equal number of values, where
// the value of the key is followed by the key.
// Example: key1,value1,key2,value2,key3,value3
func (r RESP) Map() map[string]RESP { _ = "STUB: not implemented"; return nil }

func (r RESP) MapGet(key string) RESP { _ = "STUB: not implemented"; return *new(RESP) }

func (r RESP) Exists() bool {
	_ = "STUB: not implemented"

	// ReadNextRESP returns the next resp in b and returns the number of bytes the
	// took up the result.
	return false
}

func ReadNextRESP(b []byte) (n int, resp RESP) { _ = "STUB: not implemented"; return 0, *new(RESP) }

// no data to read

// invalid kind

// read to end of line

// not enough data

//, missing CR character

// Integer

//, invalid integer

//, invalid integer

// invalid integer

// String, Error

// Bulk

// invalid number of bytes

// not enough data

// invalid end of line

// Array

// invalid number of elements

// Kind is the kind of command
type Kind int

const (
	// Redis is returned for Redis protocol commands
	Redis Kind = iota
	// Tile38 is returnd for Tile38 native protocol commands
	Tile38
	// Telnet is returnd for plain telnet commands
	Telnet
)

var errInvalidMessage = &errProtocol{"invalid message"}

// ReadNextCommand reads the next command from the provided packet. It's
// possible that the packet contains multiple commands, or zero commands
// when the packet is incomplete.
// 'argsbuf' is an optional reusable buffer and it can be nil.
// 'complete' indicates that a command was read. false means no more commands.
// 'args' are the output arguments for the command.
// 'kind' is the type of command that was read.
// 'leftover' is any remaining unused bytes which belong to the next command.
// 'err' is returned when a protocol error was encountered.
func ReadNextCommand(packet []byte, argsbuf [][]byte) (
	complete bool, args [][]byte, kind Kind, leftover []byte, err error,
) {
	_ = "STUB: not implemented"
	return false, nil, *new(Kind), nil, nil
}

// standard redis command

// done reading

func readTile38Command(packet []byte, argsbuf [][]byte) (
	complete bool, args [][]byte, kind Kind, leftover []byte, err error,
) {
	_ = "STUB: not implemented"
	return false, nil, *new(Kind), nil, nil
}

// The native protocol cannot understand json boundaries so it assumes that
// a json element must be at the end of the line.

// Setting a string value that is contained inside double quotes.
// This is only because of the boundary issues of the native protocol.

func readTelnetCommand(packet []byte, argsbuf [][]byte) (
	complete bool, args [][]byte, kind Kind, leftover []byte, err error,
) {
	_ = "STUB: not implemented"
	// just a plain text command
	return false, nil, *new(Kind), nil, nil
}

// appendPrefix will append a "$3\r\n" style redis prefix for a message.
func appendPrefix(b []byte, c byte, n int64) []byte { _ = "STUB: not implemented"; return nil }

// AppendUint appends a Redis protocol uint64 to the input bytes.
func AppendUint(b []byte, n uint64) []byte { _ = "STUB: not implemented"; return nil }

// AppendInt appends a Redis protocol int64 to the input bytes.
func AppendInt(b []byte, n int64) []byte { _ = "STUB: not implemented"; return nil }

// AppendArray appends a Redis protocol array to the input bytes.
func AppendArray(b []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

// AppendBulk appends a Redis protocol bulk byte slice to the input bytes.
func AppendBulk(b []byte, bulk []byte) []byte { _ = "STUB: not implemented"; return nil }

// AppendBulkString appends a Redis protocol bulk string to the input bytes.
func AppendBulkString(b []byte, bulk string) []byte { _ = "STUB: not implemented"; return nil }

// AppendString appends a Redis protocol string to the input bytes.
func AppendString(b []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

// AppendError appends a Redis protocol error to the input bytes.
func AppendError(b []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

// AppendOK appends a Redis protocol OK to the input bytes.
func AppendOK(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func stripNewlines(s string) string { _ = "STUB: not implemented"; return "" }

// AppendTile38 appends a Tile38 message to the input bytes.
func AppendTile38(b []byte, data []byte) []byte { _ = "STUB: not implemented"; return nil }

// AppendNull appends a Redis protocol null to the input bytes.
func AppendNull(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// AppendBulkFloat appends a float64, as bulk bytes.
func AppendBulkFloat(dst []byte, f float64) []byte { _ = "STUB: not implemented"; return nil }

// AppendBulkInt appends an int64, as bulk bytes.
func AppendBulkInt(dst []byte, x int64) []byte { _ = "STUB: not implemented"; return nil }

// AppendBulkUint appends an uint64, as bulk bytes.
func AppendBulkUint(dst []byte, x uint64) []byte { _ = "STUB: not implemented"; return nil }

func prefixERRIfNeeded(msg string) string { _ = "STUB: not implemented"; return "" }

// SimpleString is for representing a non-bulk representation of a string
// from an *Any call.
type SimpleString string

// SimpleInt is for representing a non-bulk representation of a int
// from an *Any call.
type SimpleInt int

// SimpleError is for representing an error without adding the "ERR" prefix
// from an *Any call.
type SimpleError error

// Marshaler is the interface implemented by types that
// can marshal themselves into a Redis response type from an *Any call.
// The return value is not check for validity.
type Marshaler interface {
	MarshalRESP() []byte
}

// AppendAny appends any type to valid Redis type.
//
//	nil             -> null
//	error           -> error (adds "ERR " when first word is not uppercase)
//	string          -> bulk-string
//	numbers         -> bulk-string
//	[]byte          -> bulk-string
//	bool            -> bulk-string ("0" or "1")
//	slice           -> array
//	map             -> array with key/value pairs
//	SimpleString    -> string
//	SimpleInt       -> integer
//	Marshaler       -> raw bytes
//	everything-else -> bulk-string representation using fmt.Sprint()
func AppendAny(b []byte, v interface{}) []byte { _ = "STUB: not implemented"; return nil }

type strKeyItem struct {
	key   string
	value interface{}
}
