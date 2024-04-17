package controllers

import (
	"net/http"
	"strconv"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
)

// CreateActivity creates a new activity for a working area.
func CreateActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var activity models.Activity
	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	result := initializers.DB.Create(&activity)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create activity"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// GetActivity retrieves activity by ID.
func GetActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	activityID, _ := strconv.ParseUint(id, 10, 64)

	var activity models.Activity
	result := initializers.DB.First(&activity, activityID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// UpdateActivity updates an existing activity.
func UpdateActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	activityID, _ := strconv.ParseUint(id, 10, 64)

	var activity models.Activity
	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	result := initializers.DB.Model(&models.Activity{}).Where("id = ?", activityID).Updates(&activity)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}

// DeleteActivity deletes an activity by ID.
func DeleteActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	activityID, _ := strconv.ParseUint(id, 10, 64)

	result := initializers.DB.Delete(&models.Activity{}, activityID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
