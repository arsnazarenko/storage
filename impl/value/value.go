package internal

import (
	"bytes"
    "github.com/arsnazarenko/storage/db"
)

type Value struct {

    vtype db.ValueType
    data []byte
} 

func (v *Value) Type() db.ValueType           { return v.vtype }
func (v *Value) Encode(dst interface{}) error { return nil }

func FromUint(v uint64) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteByte()
	serialization.SerializeUint(v, buf)

	return &Value{
		valueType: db.UintValueType,
		data:      buf.Bytes(),
	}

}

func FromInt(v int64) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		valueType: db.IntValueType,
		data:      buf.Bytes(),
	}
}

func FromFloat(v float64) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		valueType: db.FloatValueType,
		data:      buf.Bytes(),
	}
}

func FromString(v string) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		valueType: db.StringValueType,
		data:      buf.Bytes(),
	}
}

func FromStringList(v []string) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		valueType: db.StringListValueType,
		data:      buf.Bytes(),
	}
}

func FromObject(v serialization.Serializable) db.Value {
	buf := bytes.NewBuffer(make([]byte, 0))

	return &Value{
		valueType: db.ObjectValueType,
		data:      buf.Bytes(),
	}

}
