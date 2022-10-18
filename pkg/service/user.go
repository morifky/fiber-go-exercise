package service

import (
	"fiber-go-exercise/pkg/models"
	"fiber-go-exercise/utils"

	"go.uber.org/zap"
)

type UserService struct {
	repository *models.UserRepository
}

func New(repository *models.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) FindAllUsers() (*[]models.User, error) {
	return s.repository.FindAll()
}

func (s *UserService) FindUserByEmail(email string) (*models.User, error) {
	u, err := s.repository.FindByEmail(email)
	if err != nil {
		zap.S().Warn("Unable to find user, error: ", err)
		return nil, err
	}
	return u, nil
}

func (s *UserService) FindUserByID(id uint32) (*models.User, error) {
	u, err := s.repository.FindOne(id)
	if err != nil {
		zap.S().Warn("Unable to find user, error: ", err)
		return nil, err
	}
	return u, nil
}

func (s *UserService) CreateUser(u *models.User) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)

	if err != nil {
		zap.S().Warn("Unable to hash password, error: ", err)
	}
	_, err = s.repository.Save(u)
	if err != nil {
		zap.S().Warn("Unable to save user, error: ", err)
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id uint32) error {
	_, err := s.repository.DeleteOne(id)
	if err != nil {
		zap.S().Warn("Unable to delete user, error: ", err)
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(id uint32, u *models.User) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		zap.S().Warn("Unable to hash password, error: ", err)
	}
	_, err = s.repository.UpdateOne(id, u)

	if err != nil {
		zap.S().Warn("Unable to update user, error: ", err)
		return err
	}
	return nil
}
