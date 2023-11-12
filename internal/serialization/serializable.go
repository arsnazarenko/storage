package serialization

import "bytes"

type Serialisable interface {
	Serialize(w *bytes.Buffer) error
	Deserialize(data []byte) error
}

