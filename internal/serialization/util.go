package serialization

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/bits"
)

var (
    BufWriteError = errors.New("an error occurred while writing to the buffer")
    BufReadError = errors.New("an error occurred while reading from the buffer")
    ZeroStrLenError = errors.New("string length must be greater then 0")
    ZeroSliceLenError = errors.New("slice length must be greater then 0")
)

func SerializeUint(v uint64, buf *bytes.Buffer) error {
    var (
        bitsLen int = 0
        bytesLen int = 0
        remainder int  = 0
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
        curByte := byte((v >>(7*i)) & 0x7f | 0x80)
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
    n, err := buf.Write([]byte(s)); if err != nil || n < slen {
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
        slen int
        err error
    )
    for _, s := range(v) {
        if slen = len(s); slen == 0 {
            return fmt.Errorf("error in SerializeStringList: %w", ZeroStrLenError) 
        }
        if err = SerializeUint(uint64(slen), buf); err != nil {
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
            return BufReadError
		}
		res |= uint64((curByte & 0x7F)) << (7 * i)
		if (curByte & 0x80) == 0 {
			break
		}
		i++
	}
	return nil
}


func DeserialzeInt(reader *bytes.Reader, dst *int64) error {
    var unsigned uint64
    if err := DeserializeUint(reader, &unsigned); err != nil {
        return err
    }
    *dst = int64((unsigned >> 1) ^ -(unsigned & 0x1))
    return nil
}



func DeserializeString(reader *bytes.Reader, dst *string) error {
    var bLen uint64
    if err := DeserializeUint(reader, &bLen); err != nil {
        return err
    }
	strBytes := make([]byte, bLen, bLen)
    if n, err := reader.Read(strBytes); err != nil || n < int(bLen) {
        return BufReadError
    }
    *dst = string(strBytes)
	return nil
}

func DeserializeFloat(reader *bytes.Reader, dst *float64) error {
	var f float64
    if err := binary.Read(reader, binary.LittleEndian, &f); err != nil {
        return BufReadError
    }
    *dst = f 
	return nil
}

func DeserializeStringList(reader *bytes.Reader, dst *[]string) error {
    var (
        s string
        sliceLen uint64
    )
    if err := DeserializeUint(reader, &sliceLen); err != nil {
        return err
    }

}
