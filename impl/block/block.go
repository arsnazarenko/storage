package block

import "github.com/arsnazarenko/storage/db"

var _ db.Block = (*Block)(nil)

type Block struct {}
