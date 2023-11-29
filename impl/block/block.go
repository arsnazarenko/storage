package block

import "github.com/arsnazarenko/storage/db"

var _ db.Block = (*Block)(nil)

const BLOCK_SIZE = 2 << 12 //4096

type Block struct{}
