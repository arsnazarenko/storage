package impl

import (
	"github.com/arsnazarenko/storage/db"
)

var _ db.Storage = (*Storage)(nil)

type Storage struct {
}

// AtomicInc implements db.Storage.
func (*Storage) AtomicInc(key db.Key) error {
	panic("unimplemented")
}

// AtomicPop implements db.Storage.
func (*Storage) AtomicPop(key db.Key) error {
	panic("unimplemented")
}

// AtomicPut implements db.Storage.
func (*Storage) AtomicPut(key db.Key) error {
	panic("unimplemented")
}

// Delete implements db.Storage.
func (*Storage) Delete(key db.Key) error {
	panic("unimplemented")
}

// Get implements db.Storage.
func (*Storage) Get(key db.Key) (db.Value, error) {
	panic("unimplemented")
}

// Put implements db.Storage.
func (*Storage) Put(key db.Key, value db.Value) error {
	panic("unimplemented")
}

// Scan implements db.Storage.
func (*Storage) Scan(begin db.Key, end db.Key) (db.Iterator, error) {
	panic("unimplemented")
}

// Write implements db.Storage.
func (*Storage) Write(b *db.Batch) error {
	panic("unimplemented")
}

