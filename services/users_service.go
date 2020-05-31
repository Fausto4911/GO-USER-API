package services

import (
	"github.com/fausto4911/GO-USER-API/domain/users"
	"github.com/fausto4911/GO-USER-API/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
