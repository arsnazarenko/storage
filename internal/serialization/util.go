package serialization

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

var (
	BufWriteError     = errors.New("an error occurred while writing to the buffer")
	BufReadError      = errors.New("an error occurred while reading from the buffer")
	ZeroStrLenError   = errors.New("string length must be greater then 0")
	ZeroSliceLenError = errors.New("slice length must be greater then 0")
)

func SerializeUint(v uint64, buf *bytes.Buffer) error {
	var (
		bitsLen   int = 0
		bytesLen  int = 0
		remainder int = 0
	)
	bitsLen = bits.Len64(v)
	if bitsLen == 0 {
		buf.WriteByte(0)
		return nil
	}
	bytesLen, remainder = bitsLen/7, bitsLen%7
	if remainder > 0 {
		bytesLen++
	}
	for i := 0; i < bytesLen; i++ {
		curByte := byte((v>>(7*i))&0x7f | 0x80)
		if i == (bytesLen - 1) {
			curByte &= 0x7f
		}
		// err is always nil
		if err := buf.WriteByte(curByte); err != nil {
			return fmt.Errorf("error in SerializeUint: %w", BufWriteError)
		}
	}
	return nil
}

func SerializeInt(v int64, buf *bytes.Buffer) error {
	unsigned := uint64((v << 1) ^ (v >> (64 - 1)))
	return SerializeUint(unsigned, buf)
}

func SerializeFloat(v float64, buf *bytes.Buffer) error {
	err := binary.Write(buf, binary.LittleEndian, &v)
	if err != nil {
		return fmt.Errorf("error in SerializeInt: %w", BufWriteError)
	}
	return nil
}

func SerializeString(s string, buf *bytes.Buffer) error {
	slen := len(s)
	if slen == 0 {
		return fmt.Errorf("error in SerializeString: %w", ZeroStrLenError)
	}
	// err is always nil
	n, err := buf.Write([]byte(s))
	if err != nil || n < slen {
		return fmt.Errorf("error in SerializeFloat: %w", BufWriteError)
	}
	return nil
}

func SerializeStringList(v []string, buf *bytes.Buffer) error {
	sliceLen := len(v)
	if sliceLen == 0 {
		return fmt.Errorf("error in SerializeStringList: %w", ZeroSliceLenError)
	}
	if err := SerializeUint(uint64(sliceLen), buf); err != nil {
		return fmt.Errorf("error in SerializeStringList: %w", err)
	}
	var (
		blen int
		err  error
	)
	for _, s := range v {
		if blen = len([]byte(s)); blen == 0 {
			return fmt.Errorf("error in SerializeStringList: %w", ZeroStrLenError)
		}
		if err = SerializeUint(uint64(blen), buf); err != nil {
			return fmt.Errorf("error in SerializeStringList: %w", err)
		}
		if err = SerializeString(s, buf); err != nil {
			return fmt.Errorf("error in SerializeStringList: %w", err)
		}
	}
	return nil
}

func SerializeObject(v Serialisable, buf *bytes.Buffer) error {
	if err := v.Serialize(buf); err != nil {
		return err
	}
	return nil
}

func DeserializeUint(reader *bytes.Reader, dst *uint64) error {
	var res uint64 = 0

	i := 0
	for {
		curByte, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("error in DeserializeUint: %w", BufReadError)
		}
		res |= uint64((curByte & 0x7F)) << (7 * i)
		if (curByte & 0x80) == 0 {
			break
		}
		i++
	}
	*dst = res
	return nil
}

func DeserializeInt(reader *bytes.Reader, dst *int64) error {
	var unsigned uint64
	if err := DeserializeUint(reader, &unsigned); err != nil {
		return fmt.Errorf("error in DeserializeInt: %w", err)
	}
	*dst = int64((unsigned >> 1) ^ -(unsigned & 0x1))
	return nil
}

func DeserializeFloat(reader *bytes.Reader, dst *float64) error {
	var f float64
	if err := binary.Read(reader, binary.LittleEndian, &f); err != nil {
		return fmt.Errorf("error in DeserializeFloat: %w", BufReadError)
	}
	*dst = f
	return nil
}

func DeserializeString(reader *bytes.Reader, dst *string) error {
	bLen := reader.Len()
	strBytes := make([]byte, bLen, bLen)
	if n, err := reader.Read(strBytes); err != nil || n < int(bLen) {
		return fmt.Errorf("error in DeserializeString: %w", BufReadError)
	}
	*dst = string(strBytes)
	return nil
}

func DeserializeStringList(reader *bytes.Reader, dst *[]string) error {
	var (
		strLen   uint64
		sliceLen uint64
		strBytes []byte
		resSlice []string
	)
	if err := DeserializeUint(reader, &sliceLen); err != nil {
		return err
	}
	resSlice = make([]string, 0, sliceLen)
	for i := 0; i < int(sliceLen); i++ {
		if err := DeserializeUint(reader, &strLen); err != nil {
			return fmt.Errorf("error in DeserializeStringList: %w", err)
		}
		strBytes = make([]byte, strLen, strLen)
		if n, err := reader.Read(strBytes); n < int(strLen) || err != nil {
			return fmt.Errorf("error in DeserializeStringList: %w", BufReadError)
		}
		resSlice = append(resSlice, string(strBytes))
	}
	*dst = resSlice
	return nil
}
