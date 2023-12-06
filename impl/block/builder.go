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

func (b *BlockBuilder) Add(key []byte, value []byte) bool {
    return false
}

func (b *BlockBuilder) IsEmpty() bool {
    return (len(b.offsets) == 0)
}

func (b *BlockBuilder) Build() Block {
    if b.IsEmpty() {
        panic("Block should not be empty")
    }
    return Block{
    	data:    b.data,
    	offsets: b.offsets,
    }
}
