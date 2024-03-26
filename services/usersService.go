package services

import (
	"storeApiRest/models"
	"storeApiRest/repositories"
)

func CreateUserService(user models.User) error {
	if err := repositories.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func ReadUsersService() (models.Users, error) {
	users, err := repositories.ReadUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUserService(user models.User, userID int) error {
	if err := repositories.UpdateUser(user, userID); err != nil {
		return err
	}
	return nil
}
