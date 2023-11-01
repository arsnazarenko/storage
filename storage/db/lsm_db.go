package db

var _ Storage = (*LSMStorage)(nil)

type LSMStorage struct { }

func Open(filename string, opt *Options) (db *Storage, err error) {
    return db, err
}

func (l *LSMStorage) Get(key Key) (value Value, err error) {
    return value, err
}

func (l *LSMStorage) Put(key Key, value Value) error {
    return nil
}

func (l *LSMStorage) Scan(begin Key, end Key) (iter Iterator, err error) {
    return iter, nil
}

func (l *LSMStorage) WriteBatch(b Batch) error {
    return nil
}

func (l *LSMStorage) AtomicInc(key Key) error {
    return nil
}

func (l *LSMStorage) AtomicPut(key Key) error {
    return nil

}
func (l *LSMStorage) AtomicPop(key Key) error {
    return nil
}
