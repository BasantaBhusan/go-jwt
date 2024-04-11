package controllers

import (
	"fmt"
	"net/http"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
)

type CreateKYCRequest struct {
	FullName       string                      `json:"full_name" binding:"required"`
	MobileNumber   string                      `json:"mobile_number" binding:"required"`
	FirmRegistered bool                        `json:"firm_registered"`
	Address        CreateKYCAddressRequest     `json:"address" binding:"required"`
	WorkingArea    CreateKYCWorkingAreaRequest `json:"working_area" binding:"required"`
	Service        CreateKYCServiceRequest     `json:"service" binding:"required"`
	IsKyc          bool                        `json:"is_kyc"`
}

type CreateKYCAddressRequest struct {
	Province     string `json:"province" binding:"required"`
	District     string `json:"district" binding:"required"`
	Municipality string `json:"municipality"`
	WardNumber   string `json:"ward_number" binding:"required"`
}

type CreateKYCWorkingAreaRequest struct {
	AreaName   string   `json:"area_name" binding:"required"`
	Activities []string `json:"activities" binding:"required"`
}

type CreateKYCServiceRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
}

// @Summary Create KYC
// @Description Create KYC (Know Your Customer) record.
// @Tags KYC
// @Accept json
// @Produce json
// @Param body body CreateKYCRequest true "KYC details"
// @Success 200 {object} gin.H "KYC created successfully"
// @Failure 400 {object} gin.H "Failed to read body or create KYC"
// @Router /user/kyc [post]
func Createkyc(c *gin.Context) {
	var body CreateKYCRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	fmt.Println(body.IsKyc)
	fmt.Println(body)
	if !body.IsKyc {
		// If isKyc flag is false, do not create KYC
		c.JSON(http.StatusOK, gin.H{"message": "KYC creation skipped"})
		return
	}

	address := models.Address{
		Province:     body.Address.Province,
		District:     body.Address.District,
		Municipality: body.Address.Municipality,
		WardNumber:   body.Address.WardNumber,
	}

	workingArea := models.WorkingArea{
		AreaName: body.WorkingArea.AreaName,
	}

	// Convert []string to []models.Activities
	var activities []models.Activity
	for _, activityName := range body.WorkingArea.Activities {
		activities = append(activities, models.Activity{ActivityName: activityName})
	}

	service := models.Service{
		ServiceName: models.ServiceType(body.Service.ServiceName), // Convert string to models.ServiceType
	}

	kyc := models.Kyc{
		FullName:       body.FullName,
		MobileNumber:   body.MobileNumber,
		FirmRegistered: body.FirmRegistered,
		Address:        address,
		WorkingArea:    workingArea,
	}

	// Assign activities to WorkingArea
	kyc.WorkingArea.Activities = activities

	// Assign service to KYC
	kyc.Service = service

	result := initializers.DB.Create(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create KYC"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kyc": kyc})
}
