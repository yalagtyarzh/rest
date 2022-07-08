package dbrepo

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/yalagtyarzh/rest/internal/models"
)

// InsertUser inserts new user into DB
func (m *MemStorage) InsertUser(u models.User) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.users[u.ID]
	if ok {
		return errors.New("user already exist")
	}

	if u.Balance < 0 {
		return errors.New("cannot create user with negative money")
	}

	m.users[u.ID] = u
	return nil
}

// UpdateUser updates all user fields
func (m *MemStorage) UpdateUser(u models.User) (models.User, error) {
	if u.Balance < 0 {
		return models.User{}, errors.New("balance cannot be negative")
	}

	user, err := m.GetUserByID(u.ID)
	if err != nil {
		return models.User{}, err
	}

	user.Balance = u.Balance
	m.mutex.Lock()
	m.users[u.ID] = user
	m.mutex.Unlock()
	return user, nil
}

// GetUserByID returns user by ID
func (m *MemStorage) GetUserByID(userID int) (models.User, error) {
	m.mutex.Lock()
	user, _ := m.users[userID]
	m.mutex.Unlock()
	if reflect.DeepEqual(user, models.User{}) {
		return user, fmt.Errorf("user with %d ID not found", userID)
	}
	return user, nil
}
