package test

import (
	"context"
	"testing"
	"time"

	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/models"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/repositories"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/services"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/types/dto"
	"github.com/iacopoghilardi/golang-backend-boilerplate/test/producers"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/gorm"
)

type ProfileTestSuite struct {
	suite.Suite
	db          *gorm.DB
	pgContainer *postgres.PostgresContainer
	service     *services.ProfileService
	ctx         context.Context
}

func (s *ProfileTestSuite) SetupSuite() {
	s.ctx = context.Background()

	pgContainer, db, err := producers.SetupGenericSuite(s.ctx)
	if err != nil {
		s.T().Fatalf("Failed to setup generic suite: %v", err)
	}
	s.pgContainer = pgContainer
	s.db = db
	s.service = services.NewProfileService(repositories.NewProfileRepository(s.db))
}

func (s *ProfileTestSuite) TearDownSuite() {
	producers.TearDownGenericSuite(s.ctx, s.db, s.pgContainer)
}

func (s *ProfileTestSuite) TestCreateProfile() {
	user := producers.CreateTestUser(s.db)

	profileDto := dto.CreateProfileDto{
		ProfileDto: dto.ProfileDto{
			FirstName: "Test",
			LastName:  "User",
			BirthDate: time.Now(),
			Avatar:    "avatar.jpg",
			Bio:       "Test bio",
		},
		UserID: user.ID,
	}

	profile, err := s.service.Create(profileDto)
	s.NoError(err)
	s.NotNil(profile)
	s.Equal(profileDto.FirstName, profile.FirstName)
	s.Equal(profileDto.LastName, profile.LastName)
	s.Equal(profileDto.BirthDate, profile.BirthDate)
	s.Equal(profileDto.Avatar, profile.Avatar)
	s.Equal(profileDto.Bio, profile.Bio)
	s.Equal(user.ID, profile.UserID)
}

func (s *ProfileTestSuite) TestUpdateProfile() {
	var user *models.User
	err := s.db.Where("email = ?", producers.TestUserEmail).First(&user).Error
	if err != nil {
		s.T().Fatalf("Failed to find user: %v", err)
	}

	if user.ID == 0 || user == nil {
		user = producers.CreateTestUser(s.db)
	}

	var profile *models.Profile
	err = s.db.Where("user_id = ?", user.ID).First(&profile).Error
	if err != nil {
		s.T().Fatalf("Failed to find profile: %v", err)
	}

	if profile.ID == 0 || profile == nil {
		profile = producers.CreateTestProfile(user.ID)
	}

	profileDto := dto.UpdateProfileDto{
		ID: profile.ID,
		ProfileDto: dto.ProfileDto{
			FirstName: "Updated",
			LastName:  "User",
			BirthDate: time.Now(),
			Avatar:    "avatar.jpg",
			Bio:       "Updated bio",
		},
	}

	updatedProfile, err := s.service.UpdateProfile(profileDto)
	s.NoError(err)
	s.NotNil(updatedProfile)
	s.Equal(profileDto.FirstName, updatedProfile.FirstName)
	s.Equal(profileDto.LastName, updatedProfile.LastName)
	s.Equal(profileDto.BirthDate, updatedProfile.BirthDate)
	s.Equal(profileDto.Avatar, updatedProfile.Avatar)
	s.Equal(profileDto.Bio, updatedProfile.Bio)
}

func TestProfileTestSuite(t *testing.T) {
	suite.Run(t, new(ProfileTestSuite))
}
