package services

import (
	"storeApiRest/models"
	"storeApiRest/repositories"
)

func AuthenticateUserService(email, password string) (*models.AuthResponse, error) {
	auth, err := repositories.AuthenticateUser(email, password)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
