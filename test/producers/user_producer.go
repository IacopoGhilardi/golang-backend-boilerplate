package producers

import (
	"log"

	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/models"
	"gorm.io/gorm"
)

var TestUserEmail = "test@example.com"

func CreateTestUser(db *gorm.DB) *models.User {
	user := &models.User{
		FirstName: "Test",
		LastName:  "User",
		Email:     TestUserEmail,
		Password:  "password123",
	}

	if err := db.Create(user).Error; err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	return user
}

func DeleteTestUser(db *gorm.DB) {
	db.Where("email = ?", TestUserEmail).Delete(&models.User{})
}

func DeleteAllUsers(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}
