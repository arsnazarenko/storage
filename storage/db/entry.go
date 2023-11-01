package db

type Entry struct {
    key Key
    value Value
    deleted bool
}

func (e *Entry) Value() Value {
    return e.value
}

func (e *Entry) Key() Key {
    return e.key
}

func (e *Entry) IsDeleted() bool {
    return e.deleted
}
