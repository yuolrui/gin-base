package service

import (
	"errors"

	"github.com/yuolrui/gin-base/internal/model"
	"github.com/yuolrui/gin-base/internal/repository"
)

func GetUserByID(id string) (*model.UserRes, error) {
	user, err := repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user.ToResponse(), nil
}

func CreateUser(req model.CreateUserReq) (*model.User, error) {
	if req.Username == "" || req.Email == "" {
		return nil, errors.New("name and email required")
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
	}
	err := repository.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
