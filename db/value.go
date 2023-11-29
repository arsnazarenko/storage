package db

import (
	"errors"
	"fmt"
)

type ValueType = uint8

const moduleName = "db/value"

var (
	ZeroSliceLenError = errors.New("slice length must be greater then 0")
    UndefinedValueType = errors.New("undefined value type")
)

const (
	UintValueType       ValueType = iota + 1 // Unsigned Varint
	IntValueType                             // Signed Varint with
	FloatValueType                           // 64 bit float FloatType
	StringValueType                          // UTF-8 string
	StringListValueType                      // UTF-8 String list
)

type Value []byte

func (v *Value) Type() (ValueType, error) {
	if len(*v) == 0 {
		return 0, fmt.Errorf("%s: %w", moduleName, ZeroSliceLenError)
	}
    vtype := uint8((*v)[0])
    if vtype < 1 || vtype > 5 {
        return 0, fmt.Errorf("%s: %w", moduleName, UndefinedValueType)
    }
    return ValueType(vtype), nil
}

func (v *Value) Decode(dst interface{}) error {
	return nil
}
