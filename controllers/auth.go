package controller

import (
	"github.com/sajir-dev/go-crowdfire/services/users"
	"github.com/sajir-dev/go-crowdfire/services/users/contract"
)

func LoginController(email string, password string) (*contract.UserModel, error) {
	user, err := users.Login(&contract.LoginReq{
		Email:    email,
		Password: password,
	})

	return user, err
}
