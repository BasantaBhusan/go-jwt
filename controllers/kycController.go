package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
// @Param id path int true "User ID"
// @Param body body CreateKYCRequest true "KYC details"
// @Success 200 "KYC created successfully"
// @Failure 400 "Failed to read body or create KYC"
// @Router /user/kyc/{id} [post]
func Createkyc(c *gin.Context) {

	var body CreateKYCRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Extract user ID from path parameter
	userID := c.Param("id")

	fmt.Println("user ko id k ho van", userID)

	// Parse user ID to uint
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if !body.IsKyc {
		c.JSON(http.StatusOK, gin.H{"message": "KYC creation skipped"})
		return
	}

	address := models.Address{
		UserID:       uint(userIDUint),
		Province:     body.Address.Province,
		District:     body.Address.District,
		Municipality: body.Address.Municipality,
		WardNumber:   body.Address.WardNumber,
	}

	workingArea := models.WorkingArea{
		UserID:   uint(userIDUint),
		AreaName: body.WorkingArea.AreaName,
	}

	var activities []models.Activity
	for _, activityName := range body.WorkingArea.Activities {
		activities = append(activities, models.Activity{ActivityName: activityName})
	}

	service := models.Service{

		UserID: uint(userIDUint),

		ServiceName: models.ServiceType(body.Service.ServiceName),
	}

	kyc := models.Kyc{
		UserID:         uint(userIDUint),
		FullName:       body.FullName,
		MobileNumber:   body.MobileNumber,
		FirmRegistered: body.FirmRegistered,
		Address:        address,
		WorkingArea:    workingArea,
	}

	kyc.WorkingArea.Activities = activities
	kyc.Service = service

	result := initializers.DB.Create(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create KYC"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kyc": kyc})
}

// @Summary Get KYC by User ID
// @Description Retrieve KYC (Know Your Customer) record by User ID.
// @Tags KYC
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} models.Kyc "KYC information"
// @Failure 400 "Invalid user ID"
// @Failure 404 "KYC not found for the given user ID"
// @Router /user/kyc/{id} [get]
func GetKycByUserID(c *gin.Context) {
	userID := c.Param("id")

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var kyc models.Kyc
	// result := initializers.DB.Where("user_id = ?", userIDUint).First(&kyc)
	result := initializers.DB.
		Joins("JOIN addresses ON kycs.user_id = addresses.user_id").
		Joins("JOIN working_areas ON kycs.user_id = working_areas.user_id").
		Where("kycs.user_id = ?", userIDUint).
		First(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "KYC not found for the given user ID"})
		return
	}

	c.JSON(http.StatusOK, kyc)
}
