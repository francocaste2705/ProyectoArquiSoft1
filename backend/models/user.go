package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`    // El "-" evita que se muestre en JSON
	Role     string `gorm:"not null" json:"role"` // "admin" o "socio"
}

// TableName especifica el nombre de la tabla en la base de datos
func (User) TableName() string {
	return "users"
}
