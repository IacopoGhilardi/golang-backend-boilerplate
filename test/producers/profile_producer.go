package producers

import (
	"log"
	"time"

	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/db"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/models"
)

func CreateTestProfile(userID uint) *models.Profile {
	profile := &models.Profile{
		FirstName: "Test",
		LastName:  "User",
		BirthDate: time.Now(),
		Avatar:    "avatar.jpg",
		Bio:       "Test bio",
		UserID:    userID,
	}
	if err := db.GetDB().Create(profile).Error; err != nil {
		log.Fatalf("Failed to create profile: %v", err)
	}
	return profile
}
