package db

type ValueType = uint8

const (
	UintValueType       ValueType = iota + 1 // Unsigned Varint
	IntValueType                             // Signed Varint with
	FloatValueType                           // 64 bit float FloatType
	StringValueType                          // UTF-8 string
	StringListValueType                      // UTF-8 String list
)

type Value interface {
	Type() ValueType
	Data() []byte
	Encode(dst interface{}) error
}
