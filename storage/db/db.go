package db

type Key []byte
type Value []byte

type Options struct {}

type Batch struct {}

type Storage interface {
    Get(key Key) (Value, error)
    Put(key Key, value Value) error
    Scan(begin Key, end Key) (Iterator, error)
    WriteBatch(b Batch) error
    AtomicInc(key Key) error
    AtomicPut(key Key) error
    AtomicPop(key Key) error
}

