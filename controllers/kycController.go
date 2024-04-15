package controllers

import (
	"net/http"
	"strconv"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
// @Success 200 "KYC created successfully"
// @Failure 400 "Failed to read body or create KYC"
// @Router /user/kyc/create [post]
func Createkyc(c *gin.Context) {

	var body CreateKYCRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID := user.(models.User).ID

	if !body.IsKyc {
		c.JSON(http.StatusOK, gin.H{"message": "KYC creation skipped"})
		return
	}

	address := models.Address{
		UserID:       userID,
		Province:     body.Address.Province,
		District:     body.Address.District,
		Municipality: body.Address.Municipality,
		WardNumber:   body.Address.WardNumber,
	}

	workingArea := models.WorkingArea{
		UserID:   userID,
		AreaName: body.WorkingArea.AreaName,
	}

	var activities []models.Activity
	for _, activityName := range body.WorkingArea.Activities {
		activities = append(activities, models.Activity{ActivityName: activityName})
	}

	service := models.Service{

		UserID: userID,

		ServiceName: models.ServiceType(body.Service.ServiceName),
	}

	kyc := models.Kyc{
		UserID:         userID,
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

	// updateUser := models.User{ID: userID}
	initializers.DB.Model(&models.User{}).Where("id = ?", userID).Update("is_kyc", true)
	// initializers.DB.Model(&updateUser).Update("is_kyc", true)

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

	result := initializers.DB.Preload("Address").
		Preload("WorkingArea", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Activities")
		}).
		Preload("Service").
		Where("user_id = ?", userIDUint).
		First(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "KYC not found for the given user ID"})
		return
	}

	c.JSON(http.StatusOK, kyc)
}
