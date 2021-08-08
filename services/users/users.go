package users

import (
	"errors"

	"github.com/sajir-dev/go-crowdfire/data"
	"github.com/sajir-dev/go-crowdfire/services/users/contract"
)

func Create(req *contract.CreateUser) (*contract.UserModel, error) {

	if len(req.Email) < 5 || len(req.Password) < 5 {
		return nil, errors.New("could not verify email or password")
	}

	for _, user := range data.UsersMap {
		if user.Email == req.Email {
			return nil, errors.New("email already registered")
		}
	}

	data.UsersMap[req.Email] = &contract.UserModel{
		Id:       req.Email,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	return data.UsersMap[req.Email], nil
}

func Login(req *contract.LoginReq) (*contract.UserModel, error) {
	for _, user := range data.UsersMap {
		if user.Email == req.Email && user.Password == req.Password {
			return user, nil
		}
	}
	return nil, errors.New("invalid credentials")
}
