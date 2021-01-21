package storage

import (
	"errors"
	"http-go/models"
)

var ErrUserExists = errors.New("user alreeady exists")
var ErrUserDoesNotExist = errors.New("user does not exist")

type Users map[string]*models.User

type Adder interface {
	Add(user *models.User) error
}

type Updater interface {
	Update(user *models.User) (*models.User, error)
}

type Getter interface {
	Get(userName string) (*models.User, error)
}

type GetterAll interface {
	GetAll() Users
}

func (users Users) Add(user *models.User) error {
	if _, ok := users[user.UserName]; ok {
		return ErrUserExists
	}
	users[user.UserName] = user

	return nil
}

func (users Users) Update(user *models.User) (*models.User, error) {
	if _, ok := users[user.UserName]; !ok {
		return nil, ErrUserDoesNotExist
	}

	users[user.UserName] = user

	return users[user.UserName], nil
}

func (users Users) Get(username string) (*models.User, error) {
	if _, ok := users[username]; ok {
		return users[username], nil
	}

	return nil, ErrUserDoesNotExist
}

func (users Users) GetAll() Users {
	return users
}
