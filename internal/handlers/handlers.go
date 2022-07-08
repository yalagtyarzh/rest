package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/yalagtyarzh/rest/internal/helpers"
	"github.com/yalagtyarzh/rest/internal/models"
	"github.com/yalagtyarzh/rest/internal/repository"
	"github.com/yalagtyarzh/rest/internal/validation"
)

//Repository is the repository type
type Repository struct {
	DB repository.DatabaseRepo
}

//Repo the repository used by the handlers
var Repo *Repository

//NewRepo creates a new repository
func NewRepo(repo repository.DatabaseRepo) *Repository {
	return &Repository{DB: repo}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// CreateAccount creates new account via json
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.UnmarshalUser(r)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	err = Repo.DB.InsertUser(user)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "User created!", []models.User{user})
}

// GetBalance gets user balance by url value
func GetBalance(w http.ResponseWriter, r *http.Request) {
	required := []string{"user_id"}

	v := r.URL.Query()
	values, err := validation.ValidateURLValues(v, required...)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(values["user_id"])
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	user, err := Repo.DB.GetUserByID(id)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "Got user!", []models.User{user})
}

// TransferMoney transfer balance from first user to second
func TransferMoney(w http.ResponseWriter, r *http.Request) {
	users, err := helpers.UnmarshalUsers(r)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	if len(users) != 2 {
		helpers.ThrowError(w, http.StatusBadRequest, errors.New("invalid amount of users"))
		return
	}

	sender, err := Repo.DB.GetUserByID(users[0].ID)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	receiver, err := Repo.DB.GetUserByID(users[1].ID)
	if err != nil {
		helpers.ThrowError(w, http.StatusBadRequest, err)
		return
	}

	if sender.Balance < users[0].Balance {
		helpers.ThrowError(w, http.StatusBadRequest, errors.New("too much money to send"))
		return
	}

	sender.Balance -= users[0].Balance
	receiver.Balance += users[0].Balance

	sender, err = Repo.DB.UpdateUser(sender)
	if err != nil {
		helpers.ThrowError(w, http.StatusInternalServerError, err)
		return
	}

	receiver, err = Repo.DB.UpdateUser(receiver)
	if err != nil {
		helpers.ThrowError(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteResponse(w, http.StatusOK, "Money transferred!", []models.User{sender, receiver})
}
