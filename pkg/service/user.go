package service

import (
	"fiber-go-exercise/pkg/models"
	"fiber-go-exercise/utils"

	"go.uber.org/zap"
)

type UserService struct {
	userRepository *models.UserRepository
}

func (us *UserService) FindAllUsers() (*[]models.User, error) {
	return us.userRepository.FindAll()
}

func (us *UserService) FindUserByEmail(email string) (*models.User, error) {
	u, err := us.userRepository.FindByEmail(email)
	if err != nil {
		zap.S().Warn("Unable to find user, error: ", err)
		return nil, err
	}
	return u, nil
}

func (us *UserService) FindUserByID(id uint32) (*models.User, error) {
	u, err := us.userRepository.FindOne(id)
	if err != nil {
		zap.S().Warn("Unable to find user, error: ", err)
		return nil, err
	}
	return u, nil
}

func (us *UserService) CreateUser(u *models.User) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)

	if err != nil {
		zap.S().Warn("Unable to hash password, error: ", err)
	}
	err = us.userRepository.Save(u)
	if err != nil {
		zap.S().Warn("Unable to create a new user, error: ", err)
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(id uint32) error {
	_, err := us.userRepository.DeleteOne(id)
	if err != nil {
		zap.S().Warn("Unable to delete user, error: ", err)
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(id uint32, u *models.User) (*models.User, error) {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		zap.S().Warn("Unable to hash password, error: ", err)
		return nil, err
	}
	err = us.userRepository.UpdateOne(id, u)

	if err != nil {
		zap.S().Warn("Unable to update user, error: ", err)
		return nil, err
	}

	user, err := us.userRepository.FindOne(id)
	if err != nil {
		zap.S().Warn("Unable to find updated user data, error: ", err)
		return nil, err
	}
	return user, nil
}
