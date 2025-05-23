 package controllers

import (
	"net/http"

	"gimnasio-app/config"
	"gimnasio-app/models"

	"github.com/gin-gonic/gin"
)

// CreateActivity crea una nueva actividad
func CreateActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&activity)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la actividad"})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

// GetActivities obtiene todas las actividades
func GetActivities(c *gin.Context) {
	var activities []models.Activity
	filter := c.Query("filter")

	query := config.DB.Model(&models.Activity{})
	if filter != "" {
		query = query.Where("title LIKE ? OR category LIKE ?", "%"+filter+"%", "%"+filter+"%")
	}

	if err := query.Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las actividades"})
		return
	}

	c.JSON(http.StatusOK, activities)
}

// GetActivityByID obtiene una actividad por su ID
func GetActivityByID(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := config.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// UpdateActivity actualiza una actividad existente
func UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := config.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&activity)
	c.JSON(http.StatusOK, activity)
}

// DeleteActivity elimina una actividad
func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := config.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	config.DB.Delete(&activity)
	c.JSON(http.StatusOK, gin.H{"message": "Actividad eliminada correctamente"})
}
