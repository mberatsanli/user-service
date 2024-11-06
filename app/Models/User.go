package Models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model

	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Email       string `gorm:"index" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	Picture     string `json:"picture"`
	CreatedById string `json:"created_by"`

	Permissions []Permission `gorm:"many2many:user_permissions;" json:"permissions"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}
