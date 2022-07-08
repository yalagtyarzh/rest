package helpers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yalagtyarzh/rest/internal/models"
)

// UnmarshalUser unmarshals json user into user's struct
func UnmarshalUser(r *http.Request) (models.User, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return models.User{}, err
	}

	var user models.User

	err = json.Unmarshal(b, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// UnmarshalUsers unmarshals json array of users into slice of users' structs
func UnmarshalUsers(r *http.Request) ([]models.User, error) {
	var users []models.User
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(b, &users)
	if err != nil {
		return users, err
	}

	return users, nil
}
