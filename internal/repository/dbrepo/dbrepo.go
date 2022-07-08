package dbrepo

import (
	"sync"

	"github.com/yalagtyarzh/rest/internal/models"
)

// MemStorage is a memory storage struct
type MemStorage struct {
	mutex *sync.Mutex
	users map[int]models.User
}

// NewMemStorage returns new memory storage
func NewMemStorage() *MemStorage {
	return &MemStorage{
		mutex: &sync.Mutex{},
		users: make(map[int]models.User),
	}
}
