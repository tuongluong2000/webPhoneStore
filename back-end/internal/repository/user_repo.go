package repository

import "github.com/tuongluong2000/webPhoneStore/back-end/internal/model"

func FindUserByEmailAndPass(email string, password string) (*model.User, error) {
	// giả lập DB
	if email == "admin@gmail.com" && password == "123456" {
		return &model.User{
			ID:       1,
			Email:    "admin@gmail.com",
			Password: "123456",
		}, nil
	}

	return nil, nil
}
