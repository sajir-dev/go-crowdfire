package users

import (
	"errors"

	"github.com/sajir-dev/go-crowdfire/domain"
	"github.com/sajir-dev/go-crowdfire/services/users/contract"
)

func Create(req *contract.CreateUser) (*contract.UserModel, error) {

	if len(req.Email) < 5 || len(req.Password) < 5 {
		return nil, errors.New("invalid email or password")
	}

	if err := domain.CreateUser(req); err != nil {
		return nil, err
	}

	res, err := domain.Login(&contract.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Login(req *contract.LoginReq) (*contract.UserModel, error) {
	res, err := domain.Login(req)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return res, nil
}
