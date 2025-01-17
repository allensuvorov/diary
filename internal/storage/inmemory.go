package storage

import "github.com/allensuvorov/diary/internal/domain/entities"

type noteInMemoryStorage struct {
	notes map[int]map[int]entities.Note
}

func NewNoteInMemoryStorage() *noteInMemoryStorage {
	return &noteInMemoryStorage{
		notes: make(map[int]map[int]entities.Note),
	}
}
