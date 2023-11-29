package block

import (
	"unsafe"
)

type BlockBuilder struct {
	offsets   []uint16
	data      []byte
	blockSize uint
}

const U16_SIZE = uint(unsafe.Sizeof(uint16(0)))


func New(blockSize uint) BlockBuilder {
	return BlockBuilder{
		offsets:   make([]uint16, 0),
		data:      make([]byte, 0),
		blockSize: blockSize,
	}
}

func (b *BlockBuilder) EstimatedSize() uint {
    return uint(len(b.offsets)) * U16_SIZE + uint(len(b.data)) + U16_SIZE 
}
