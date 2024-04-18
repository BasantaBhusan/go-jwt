package controllers

import (
	"net/http"
	"strconv"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
)

type CreateWorkingAreaRequest struct {
	AreaName   string                `json:"area_name" binding:"required"`
	Activities []CreateActivityItems `json:"activities" binding:"required"`
}

type UpdateWorkingAreaRequest struct {
	AreaName   string                `json:"area_name"`
	Activities []UpdateActivityItems `json:"activities"`
}

func CreateWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req CreateWorkingAreaRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var workingArea models.WorkingArea
	workingArea.AreaName = req.AreaName

	for _, activityReq := range req.Activities {
		activity := models.Activity{ActivityName: activityReq.ActivityName}

		if len(activityReq.Items) > 0 {
			for _, item := range activityReq.Items {
				activity.Items = append(activity.Items, models.ActivityItem{Name: item})
			}
		}

		workingArea.Activities = append(workingArea.Activities, activity)
	}

	result := initializers.DB.Create(&workingArea)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create working area"})
		return
	}

	c.JSON(http.StatusOK, workingArea)
}

func GetWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	workingAreaID, _ := strconv.ParseUint(id, 10, 64)

	var workingArea models.WorkingArea
	result := initializers.DB.Preload("Activities").First(&workingArea, workingAreaID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Working area not found"})
		return
	}

	c.JSON(http.StatusOK, workingArea)
}

func UpdateWorkingArea(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	workingAreaID, _ := strconv.ParseUint(id, 10, 64)

	var req UpdateWorkingAreaRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var workingArea models.WorkingArea
	workingArea.AreaName = req.AreaName

	for _, activityReq := range req.Activities {
		activity := models.Activity{ActivityName: activityReq.ActivityName}

		if len(activityReq.Items) > 0 {
			for _, item := range activityReq.Items {
				activity.Items = append(activity.Items, models.ActivityItem{Name: item})
			}
		}

		workingArea.Activities = append(workingArea.Activities, activity)
	}

	result := initializers.DB.Model(&models.WorkingArea{}).Where("id = ?", workingAreaID).Updates(&workingArea)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update working area"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Working area updated successfully"})
}
