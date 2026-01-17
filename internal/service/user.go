package service

import (
	"errors"

	v1 "github.com/yuolrui/gin-base/internal/controller/v1"
	"github.com/yuolrui/gin-base/internal/repository"
)

func GetUserByID(id string) (*repository.User, error) {
	user, err := repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(req v1.CreateUserRequest) (*repository.User, error) {
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("name and email required")
	}

	user := &repository.User{
		Name:  req.Name,
		Email: req.Email,
	}
	err := repository.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type CreateUserRequest struct {
	Name  string
	Email string
}
