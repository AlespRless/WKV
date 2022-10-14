package storage

type Key interface{}

type Value interface{}

type Storage interface {
	GetIfPresent(Key) (Value, bool)

	Put(Key, Value)

	Delete(Key)

	DeleteAll()
}
