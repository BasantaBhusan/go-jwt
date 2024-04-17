package controllers

import (
	"net/http"
	"strconv"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
)

// CreateWorkingArea creates a new working area for a user.
func CreateWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var workingArea models.WorkingArea
	if err := c.BindJSON(&workingArea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	result := initializers.DB.Create(&workingArea)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create working area"})
		return
	}

	c.JSON(http.StatusOK, workingArea)
}

// GetWorkingArea retrieves working area by ID.
func GetWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	workingAreaID, _ := strconv.ParseUint(id, 10, 64)

	var workingArea models.WorkingArea
	result := initializers.DB.First(&workingArea, workingAreaID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Working area not found"})
		return
	}

	c.JSON(http.StatusOK, workingArea)
}

// UpdateWorkingArea updates an existing working area.
func UpdateWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	workingAreaID, _ := strconv.ParseUint(id, 10, 64)

	var workingArea models.WorkingArea
	if err := c.BindJSON(&workingArea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	result := initializers.DB.Model(&models.WorkingArea{}).Where("id = ?", workingAreaID).Updates(&workingArea)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update working area"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Working area updated successfully"})
}

// DeleteWorkingArea deletes a working area by ID.
func DeleteWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	workingAreaID, _ := strconv.ParseUint(id, 10, 64)

	result := initializers.DB.Delete(&models.WorkingArea{}, workingAreaID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete working area"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Working area deleted successfully"})
}
