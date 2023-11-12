package internal

import (
	"bytes"

	db "github.com/arsnazarenko/storage/internal"
	"github.com/arsnazarenko/storage/internal/serialization"
)

var _ db.Value = (*Value)(nil)


type Value struct {
    valueType db.ValueType
    data []byte
}

func (v *Value) Type() db.ValueType { return v.valueType }
func (v *Value) Data() []byte { return v.data }
func (v *Value) Encode(dst interface{}) error { return nil }


func FromUint(v uint64) db.Value {
    buf := bytes.NewBuffer(make([]byte, 0))
    
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

func FromObject(v serialization.Serialisable) db.Value {
    buf := bytes.NewBuffer(make([]byte, 0))

    return &Value{
    	valueType: db.ObjectValueType,
    	data:      buf.Bytes(),
    }
    
}


