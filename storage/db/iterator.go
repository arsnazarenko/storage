package db

type Iterator interface {
    HasNext() bool
    GetNext() *Entry
}


