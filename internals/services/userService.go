package services

import (
	"fmt"
	"log"

	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/mappers"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/models"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/repositories"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/types/dto"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetById(id uint) (models.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (s *UserService) Create(user *dto.CreateUserDto) (models.User, error) {
	userModel := mappers.GetUserDtoFromCreateUserDto(user)

	createdUser, err := s.userRepository.Create(&userModel)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return models.User{}, err
	}

	return *createdUser, nil
}

func (s *UserService) Update(user dto.UpdateUserDto) (models.User, error) {
	userModel := mappers.UpdateUserDtoToUserModel(&user)

	oldUser, err := s.userRepository.FindById(user.ID)
	if err != nil {
		log.Printf("error finding user: %+v\n", err)
		return models.User{}, err
	}

	updatedUser, err := s.userRepository.Update(oldUser, &userModel)
	if err != nil {
		log.Printf("error updating user: %+v\n", err)
		return models.User{}, err
	}

	return *updatedUser, nil
}

func (s *UserService) Delete(id uint) error {
	log.Printf("deleting user with id: %d\n", id)
	err := s.userRepository.Delete(id)
	if err != nil {
		log.Printf("error deleting user: %+v\n", err)
		return err
	}
	return nil
}
