package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Schedule    string `gorm:"not null" json:"schedule"` // Formato: "YYYY-MM-DD HH:mm"
	Capacity    int    `gorm:"not null" json:"capacity"`
	Category    string `gorm:"not null" json:"category"`
	Instructor  string `gorm:"not null" json:"instructor"`
	ImageURL    string `json:"image_url"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Activity) TableName() string {
	return "activities"
}
