package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Instructor  string `gorm:"not null" json:"instructor"`
	Duration    int    `gorm:"not null" json:"duration"` // duración en minutos
	Image       string `json:"image"`
	Status      string `gorm:"not null" json:"status"` // "active" o "inactive"
	DayOfWeek   string `gorm:"not null" json:"day_of_week"`
	StartTime   string `gorm:"not null" json:"start_time"` // formato "HH:mm"
	EndTime     string `gorm:"not null" json:"end_time"`   // formato "HH:mm"
	Capacity    int    `gorm:"not null" json:"capacity"`
	Category    string `gorm:"not null" json:"category"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Activity) TableName() string {
	return "activities"
}
