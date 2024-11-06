package Models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model

	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null" json:"name"`
	Identifier  string `gorm:"unique;not null" json:"identifier"`
	Description string `json:"description"`
	Service     string `gorm:"size:80" json:"service"`
	Olac        bool   `json:"olac"`
}
