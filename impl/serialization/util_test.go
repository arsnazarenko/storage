package serialization

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func genStringSlice(length int) []string {
	s := make([]string, length, length)
	for i := 0; i < length; i++ {
		s[i] = "some_repeated_string"
	}
	return s
}

func TestSerializeUint(t *testing.T) {

	tests := []struct {
		name  string
		value uint64
	}{
		{name: "ZeroUnsigned", value: 0},
		{name: "MaxUnsigned", value: math.MaxUint64},
		{name: "SomeUnsigned", value: 123123123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out uint64
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeUint(tt.value, buf)
			require.NoError(t, err)
			require.NoError(t, err)
			err = DeserializeUint(bytes.NewReader(buf.Bytes()), &out)
			require.NoError(t, err)
			require.Equal(t, out, tt.value)
		})
	}
}
func TestSerializeInt(t *testing.T) {
	tests := []struct {
		name  string
		value int64
	}{
		{name: "MinSigned", value: math.MinInt64},
		{name: "MaxSigned", value: math.MaxInt64},
		{name: "ZeroSigned", value: 0},
		{name: "NegativeSigned", value: -1788723},
		{name: "PositiveSigned", value: 1883939},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out int64
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeInt(tt.value, buf)
			require.NoError(t, err)
			err = DeserializeInt(bytes.NewReader(buf.Bytes()), &out)
			require.NoError(t, err)
			require.Equal(t, out, tt.value)
		})
	}
}

func TestSerializeFloat(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{name: "MinFloat", value: -math.MaxFloat64},
		{name: "MaxFloat", value: math.MaxFloat64},
		{name: "ZeroFloat", value: 0},
		{name: "NegativeFloat", value: -1788.123723},
		{name: "PositiveFloat", value: 18831000.991939},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out float64
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeFloat(tt.value, buf)
			require.NoError(t, err)
			err = DeserializeFloat(bytes.NewReader(buf.Bytes()), &out)
			require.NoError(t, err)
			require.Equal(t, out, tt.value)
		})
	}
}

func TestSerializeString(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "UTF-8 stirng", value: "Проверка строки!!!!!!!!!!!!!!"},
		{name: "Some stirng", value: "asdajhgdajkshgdakjshgdkjahgsdka"},
		{name: "Very long string", value: strings.Repeat("some_very_long_string", 100000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out string
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeString(tt.value, buf)
			require.NoError(t, err)
			err = DeserializeString(bytes.NewReader(buf.Bytes()), &out)
			require.NoError(t, err)
			require.Equal(t, out, tt.value)
		})
	}
}

func TestSerializeStringList(t *testing.T) {
	tests := []struct {
		name  string
		value []string
	}{
		{name: "Some string slice", value: []string{"adasda", "adasdasd", "adasdasd", "asdasdasd"}},
		{name: "One element long string slice", value: []string{strings.Repeat("somestringthatsrepeat", 100000)}},
		{name: "Big string slice", value: genStringSlice(1000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out []string
			buf := bytes.NewBuffer(make([]byte, 0))
			err := SerializeStringList(tt.value, buf)
			require.NoError(t, err)
			err = DeserializeStringList(bytes.NewReader(buf.Bytes()), &out)
			require.NoError(t, err)
			require.Equal(t, out, tt.value)
		})
	}
}
