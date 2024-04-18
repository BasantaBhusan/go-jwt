package controllers

import (
	"net/http"
	"strconv"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
)

type CreateActivityRequest struct {
	ActivityName string   `json:"activity_name" binding:"required"`
	Items        []string `json:"items"`
}

type UpdateActivityRequest struct {
	ActivityName string   `json:"activity_name"`
	Items        []string `json:"items"`
}

func CreateActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req CreateActivityRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var activity models.Activity
	activity.ActivityName = req.ActivityName

	if len(req.Items) > 0 {
		for _, item := range req.Items {
			activity.Items = append(activity.Items, models.ActivityItem{Name: item})
		}
	}

	result := initializers.DB.Create(&activity)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create activity"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	activityID, _ := strconv.ParseUint(id, 10, 64)

	var activity models.Activity
	result := initializers.DB.Preload("Items").First(&activity, activityID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func UpdateActivity(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	activityID, _ := strconv.ParseUint(id, 10, 64)

	var req UpdateActivityRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var activity models.Activity
	activity.ActivityName = req.ActivityName

	if len(req.Items) > 0 {
		var items []models.ActivityItem
		for _, item := range req.Items {
			items = append(items, models.ActivityItem{Name: item})
		}
		activity.Items = items
	}

	result := initializers.DB.Model(&models.Activity{}).Where("id = ?", activityID).Updates(&activity)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}
