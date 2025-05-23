package controllers

import (
	"net/http"

	"gimnasio-app/config"
	"gimnasio-app/models"

	"github.com/gin-gonic/gin"
)

// CreateEnrollment crea una nueva inscripción
func CreateEnrollment(c *gin.Context) {
	var enrollment models.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si hay cupo disponible
	var activity models.Activity
	if err := config.DB.First(&activity, enrollment.ActivityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	// Contar inscripciones activas
	var count int64
	config.DB.Model(&models.Enrollment{}).Where("activity_id = ? AND status = ?", enrollment.ActivityID, "active").Count(&count)
	if count >= int64(activity.Capacity) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No hay cupos disponibles"})
		return
	}

	// Verificar si el usuario ya está inscrito
	var existingEnrollment models.Enrollment
	if err := config.DB.Where("user_id = ? AND activity_id = ? AND status = ?", enrollment.UserID, enrollment.ActivityID, "active").First(&existingEnrollment).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario ya está inscrito en esta actividad"})
		return
	}

	enrollment.Status = "active"
	result := config.DB.Create(&enrollment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la inscripción"})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}

// GetUserEnrollments obtiene todas las inscripciones de un usuario
func GetUserEnrollments(c *gin.Context) {
	userID := c.Param("userID")
	var enrollments []models.Enrollment

	if err := config.DB.Preload("Activity").Where("user_id = ?", userID).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las inscripciones"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

// CancelEnrollment cancela una inscripción
func CancelEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enrollment models.Enrollment

	if err := config.DB.First(&enrollment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inscripción no encontrada"})
		return
	}

	enrollment.Status = "cancelled"
	config.DB.Save(&enrollment)
	c.JSON(http.StatusOK, gin.H{"message": "Inscripción cancelada correctamente"})
}
