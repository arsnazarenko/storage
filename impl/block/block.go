package block

const BLOCK_SIZE = 2 << 12 //4096

type Block struct{
    data []byte
    offsets []uint16
}
