package storage

type noteInMemoryStorage struct {
	notes map[int]string
}

func NewNoteInMemoryStorage() *noteInMemoryStorage {
	return &noteInMemoryStorage{
		notes: make(map[int]string),
	}
}
