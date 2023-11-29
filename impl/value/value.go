package internal

import (
	"bytes"
	"github.com/arsnazarenko/storage/db"
)

// Data implements db.Value.
func (v *Value) Data() []byte                 { return v.data }
func (v *Value) Type() db.ValueType           { return v.vtype }
func (v *Value) Encode(dst interface{}) error { return nil }


func FromUint(v uint64) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))
	return &Value{
		vtype: db.UintValueType,
		data:      buf.Bytes(),
	}

}

func FromInt(v int64) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		vtype: db.IntValueType,
		data:      buf.Bytes(),
	}
}

func FromFloat(v float64) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		vtype: db.FloatValueType,
		data:      buf.Bytes(),
	}
}

func FromString(v string) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		vtype: db.StringValueType,
		data:      buf.Bytes(),
	}
}

func FromStringList(v []string) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		vtype: db.StringListValueType,
		data:      buf.Bytes(),
	}
}

