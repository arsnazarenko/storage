package db

type Key []byte

type Options struct{}

type Batch struct{}

type Storage interface {
	Get(key Key) (Value, error)
	Put(key Key, value Value) error
	Delete(key Key) error
	Scan(begin, end Key) (Iterator, error)
	Write(b *Batch) error

	// operations for values with cpecific types
	AtomicInc(key Key) error
	AtomicPut(key Key) error
	AtomicPop(key Key) error
}
