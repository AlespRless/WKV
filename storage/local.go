package storage

type storage struct {
	DB map[Key]Value
}

func New() Storage {
	s := newStorage()
	return s
}

func newStorage() *storage {
	s := &storage{
		DB: map[Key]Value{},
	}
	return s
}

func (s *storage) GetIfPresent(key Key) (Value, bool) {
	if v, has := s.DB[key]; !has {
		return nil, false
	} else {
		return v, true
	}
}

func (s *storage) Put(key Key, value Value) {
	s.DB[key] = value
}

func (s *storage) Delete(key Key) {
	delete(s.DB, key)
}

func (s *storage) DeleteAll() {
	s.DB = map[Key]Value{}
}
