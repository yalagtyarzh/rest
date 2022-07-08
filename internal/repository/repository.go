package repository

import "github.com/yalagtyarzh/rest/internal/models"

// DatabaseRepo says what methods DB shall implement
type DatabaseRepo interface {
	InsertUser(u models.User) error
	UpdateUser(u models.User) (models.User, error)
	GetUserByID(userID int) (models.User, error)
}
