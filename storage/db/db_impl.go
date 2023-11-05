package db

var _ Storage = (*StorageImpl)(nil)

type StorageImpl struct{}

func Open(filename string, opt *Options) (db *Storage, err error) {
	return db, err
}

func (s *StorageImpl) Get(key Key) (value Value, err error) {
	return value, err
}

func (s *StorageImpl) Put(key Key, value Value) error {
	return nil
}

func (s *StorageImpl) Scan(begin Key, end Key) (iter Iterator, err error) {
	return iter, nil
}

func (s *StorageImpl) WriteBatch(b Batch) error {
	return nil
}

func (s *StorageImpl) AtomicInc(key Key) error {
	return nil
}

func (s *StorageImpl) AtomicPut(key Key) error {
	return nil

}
func (s *StorageImpl) AtomicPop(key Key) error {
	return nil
}
