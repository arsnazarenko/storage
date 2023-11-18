package impl 

import (
	db "github.com/arsnazarenko/storage/internal/interface"
)

var _ db.Storage = (*Storage)(nil)

type Storage struct {
}

// AtomicInc implements internal.Storage.
func (*Storage) AtomicInc(key db.Key) error {

	panic("unimplemented")
}

// AtomicPop implements internal.Storage.
func (*Storage) AtomicPop(key db.Key) error {
	panic("unimplemented")
}

// AtomicPut implements internal.Storage.
func (*Storage) AtomicPut(key db.Key) error {
	panic("unimplemented")
}

// Get implements internal.Storage.
func (*Storage) Get(key db.Key) (db.Value, error) {
	panic("unimplemented")
}

// Put implements internal.Storage.
func (*Storage) Put(key db.Key, value db.Value) error {
	panic("unimplemented")
}

// Scan implements internal.Storage.
func (*Storage) Scan(begin db.Key, end db.Key) (db.Iterator, error) {
	panic("unimplemented")
}

// WriteBatch implements internal.Storage.
func (*Storage) WriteBatch(b *db.Batch) error {
	panic("unimplemented")
}
