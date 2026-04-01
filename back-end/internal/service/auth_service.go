package service

import (
	"errors"

	"github.com/tuongluong2000/webPhoneStore/back-end/internal/repository"
	"github.com/tuongluong2000/webPhoneStore/back-end/internal/utils"
)

func Login(email, password string) (string, error) {
	_, err := repository.FindUserByEmailAndPass(email, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	return utils.GenerateJWT(email), nil
}
