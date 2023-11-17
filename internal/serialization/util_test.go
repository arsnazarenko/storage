package serialization

import (
	"bytes"
	"math"
	"testing"
)

func noError(t *testing.T, e error) {
	if e != nil {
		t.Errorf("Error occured: %s", e.Error())
	}
}

func TestSerializeUint(t *testing.T) {

	tests := []struct {
		name string
		want uint64
	}{
		{name: "ZeroUnsigned", want: 0},
		{name: "MaxUnsigned", want: math.MaxUint64},
		{name: "SomeUnsigned", want: 123123123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out uint64
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeUint(tt.want, buf)
			noError(t, err)
			err = DeserializeUint(bytes.NewReader(buf.Bytes()), &out)
			noError(t, err)
			if out != tt.want {
				t.Errorf("got %v, want %v", buf, tt.want)
			}
		})
	}
}
func TestSerializeInt(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "MinSigned", want: math.MinInt64},
		{name: "MaxSigned", want: math.MaxInt64},
		{name: "ZeroSigned", want: 123123123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out uint64
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeUint(tt.want, buf)
			noError(t, err)
			err = DeserializeUint(bytes.NewReader(buf.Bytes()), &out)
			noError(t, err)
			if out != tt.want {
				t.Errorf("got %v, want %v", buf, tt.want)
			}
		})
	}
}

func TestSerializeFloat(t *testing.T) {

}

func TestSerializeString(t *testing.T) {

}

func TestSerializeStringList(t *testing.T) {

}
