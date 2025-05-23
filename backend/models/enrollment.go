package models

import (
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID     uint     `gorm:"not null" json:"user_id"`
	ActivityID uint     `gorm:"not null" json:"activity_id"`
	User       User     `gorm:"foreignKey:UserID" json:"user"`
	Activity   Activity `gorm:"foreignKey:ActivityID" json:"activity"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Enrollment) TableName() string {
	return "enrollments"
}
