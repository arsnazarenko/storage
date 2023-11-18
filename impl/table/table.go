package table

import "github.com/arsnazarenko/storage/db"

var _ db.MemTable = (*MemTable)(nil)

type MemTable struct {}
