package internal

type Iterator interface {
	First() bool
	Last() bool

	Next() bool
	Prev() bool

	Seek(key Key) bool

	Key() Key
	Value() Value
}
