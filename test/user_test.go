package test

import (
	"context"
	"log"
	"testing"

	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/repositories"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/services"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/types/dto"
	"github.com/iacopoghilardi/golang-backend-boilerplate/test/producers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/gorm"
)

type UserTestSuite struct {
	suite.Suite
	db          *gorm.DB
	pgContainer *postgres.PostgresContainer
	service     *services.UserService
	ctx         context.Context
}

func (s *UserTestSuite) SetupSuite() {
	s.ctx = context.Background()

	pgContainer, db, err := producers.SetupGenericSuite(s.ctx)
	if err != nil {
		s.T().Fatalf("Failed to setup generic suite: %v", err)
	}
	s.pgContainer = pgContainer
	s.db = db
	s.service = services.NewUserService(repositories.NewUserRepository(s.db))
}

func (s *UserTestSuite) TearDownSuite() {
	producers.TearDownGenericSuite(s.ctx, s.db, s.pgContainer)
}

func (s *UserTestSuite) TestCreateUser() {
	CreateUserDto := dto.CreateUserDto{
		Email:     "test@test.com",
		Password:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	user, err := s.service.Create(&CreateUserDto)
	if err != nil {
		log.Fatal("Failed to create user: ", err)
	}

	assert.NotNil(s.T(), user)
	assert.Equal(s.T(), user.Email, CreateUserDto.Email)
	assert.Equal(s.T(), user.FirstName, CreateUserDto.FirstName)
	assert.Equal(s.T(), user.LastName, CreateUserDto.LastName)
}

func (s *UserTestSuite) TestFindAll() {
	producers.DeleteAllUsers(s.db)
	producers.CreateTestUser(s.db)
	users, err := s.service.GetAll()
	if err != nil {
		log.Fatal("Failed to find all users: ", err)
	}
	assert.NotNil(s.T(), users)
	assert.Greater(s.T(), len(users), 0)
	assert.Equal(s.T(), users[0].Email, producers.TestUserEmail)
	assert.Equal(s.T(), len(users), 1)
}

func (s *UserTestSuite) TestFindById() {
	producers.DeleteAllUsers(s.db)
	user := producers.CreateTestUser(s.db)
	foundUser, err := s.service.GetById(user.ID)
	if err != nil {
		log.Fatal("Failed to find user by id: ", err)
	}
	assert.NotNil(s.T(), foundUser)
	assert.Equal(s.T(), foundUser.Email, producers.TestUserEmail)
}

func (s *UserTestSuite) TestDelete() {
	producers.DeleteAllUsers(s.db)
	user := producers.CreateTestUser(s.db)
	err := s.service.Delete(user.ID)
	if err != nil {
		log.Fatal("Failed to delete user: ", err)
	}
	assert.Nil(s.T(), err)
}

func (s *UserTestSuite) TestUpdate() {
	producers.DeleteAllUsers(s.db)
	user := producers.CreateTestUser(s.db)
	updatedUser, err := s.service.Update(dto.UpdateUserDto{
		ID:        user.ID,
		Email:     "test2@test.com",
		FirstName: "test2",
		LastName:  "test2",
	})
	if err != nil {
		log.Fatal("Failed to update user: ", err)
	}
	assert.NotNil(s.T(), updatedUser)
	assert.Equal(s.T(), updatedUser.Email, "test2@test.com")
	assert.Equal(s.T(), updatedUser.FirstName, "test2")
	assert.Equal(s.T(), updatedUser.LastName, "test2")
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
