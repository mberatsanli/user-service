package Models

import "gorm.io/gorm"

type UserPermission struct {
	gorm.Model

	User         User       `gorm:"foreignKey:id" json:"user"`
	UserId       uint       `gorm:"not null" json:"user_id"`
	Permission   Permission `gorm:"foreignKey:id" json:"permission"`
	PermissionId uint       `gorm:"not null" json:"permission_id"`
}
