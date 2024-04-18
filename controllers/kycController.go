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
	Latitude     string `json:"latitude" `
	Longitude    string `json:"longitude"`
}

type CreateKYCWorkingAreaRequest struct {
	AreaName   string                `json:"area_name" binding:"required"`
	Activities []CreateActivityItems `json:"activities" binding:"required"`
}

type CreateActivityItems struct {
	ActivityName string   `json:"activity_name" binding:"required"`
	Items        []string `json:"items"`
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

	// Create address
	address := models.Address{
		UserID:       userID,
		Province:     body.Address.Province,
		District:     body.Address.District,
		Municipality: body.Address.Municipality,
		WardNumber:   body.Address.WardNumber,
		Longitude:    body.Address.Longitude,
		Latitude:     body.Address.Latitude,
	}

	// Create working area
	workingArea := models.WorkingArea{
		UserID:   userID,
		AreaName: body.WorkingArea.AreaName,
	}

	// Create activities
	var activities []models.Activity
	for _, activityReq := range body.WorkingArea.Activities {
		activity := models.Activity{ActivityName: activityReq.ActivityName}

		// Create activity items
		var items []models.ActivityItem
		for _, item := range activityReq.Items {
			items = append(items, models.ActivityItem{Name: item})
		}
		activity.Items = items

		activities = append(activities, activity)
	}

	// Create service
	service := models.Service{
		UserID:      userID,
		ServiceName: models.ServiceType(body.Service.ServiceName),
	}

	// Create KYC
	kyc := models.Kyc{
		UserID:         userID,
		FullName:       body.FullName,
		MobileNumber:   body.MobileNumber,
		FirmRegistered: body.FirmRegistered,
		Address:        address,
		WorkingArea:    workingArea,
		Service:        service,
	}

	// Associate activities with working area
	kyc.WorkingArea.Activities = activities

	// Create KYC record
	result := initializers.DB.Create(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create KYC"})
		return
	}

	// Update user's KYC status
	initializers.DB.Model(&models.User{}).Where("id = ?", userID).Update("is_kyc", true)

	c.JSON(http.StatusOK, gin.H{"message": "Kyc Successfully Completed", "kyc": kyc})

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
	userRole, exists := c.Get("role")
	if !exists || (userRole != "ADMIN" && userRole != "USER") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID := c.Param("id")

	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// If the user is not an admin, check if the requested user ID matches their own
	if userRole != "ADMIN" {
		sub, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		userIDFromToken, ok := sub.(uint)
		if !ok || userIDFromToken != uint(userIDUint) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
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

	c.JSON(http.StatusOK, gin.H{"message": "User Successfully fetched", "kyc": kyc})
}

// @Summary Update KYC by User ID
// @Description Update KYC (Know Your Customer) record by User ID.
// @Tags KYC
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Param body body UpdateKYCRequest true "KYC details"
// @Success 200 "KYC updated successfully"
// @Failure 400 "Invalid user ID or failed to read body"
// @Failure 404 "KYC not found for the given user ID"
// @Router /user/kyc/update/{id} [put]
func UpdateKycByUserID(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || (userRole != "ADMIN" && userRole != "USER") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var userIDUint uint
	if userRole == "USER" {
		userIDFromTokenRaw, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		userIDUint = userIDFromTokenRaw.(uint)
	} else {
		userIDStr := c.Param("id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		userIDUint = uint(userID)
	}

	var body UpdateKYCRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var kyc models.Kyc
	result := initializers.DB.Where("user_id = ?", userIDUint).First(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "KYC not found for the given user ID"})
		return
	}

	if userRole == "USER" && kyc.UserID != userIDUint {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	kyc.FullName = body.FullName
	kyc.MobileNumber = body.MobileNumber
	kyc.FirmRegistered = body.FirmRegistered
	if body.Address.Province != "" {
		kyc.Address.Province = body.Address.Province
	}
	if body.Address.District != "" {
		kyc.Address.District = body.Address.District
	}
	if body.Address.Municipality != "" {
		kyc.Address.Municipality = body.Address.Municipality
	}
	if body.Address.WardNumber != "" {
		kyc.Address.WardNumber = body.Address.WardNumber
	}
	if body.Address.Latitude != "" {
		kyc.Address.Latitude = body.Address.Latitude
	}
	if body.Address.Longitude != "" {
		kyc.Address.Longitude = body.Address.Longitude
	}
	if body.WorkingArea.AreaName != "" {
		kyc.WorkingArea.AreaName = body.WorkingArea.AreaName
	}
	var activities []models.Activity
	for _, activityReq := range body.WorkingArea.Activities {
		activity := models.Activity{ActivityName: activityReq.ActivityName}

		if len(activityReq.Items) > 0 {
			var items []models.ActivityItem
			for _, item := range activityReq.Items {
				items = append(items, models.ActivityItem{Name: item})
			}
			activity.Items = items
		}

		activities = append(activities, activity)
	}
	kyc.WorkingArea.AreaName = body.WorkingArea.AreaName
	kyc.WorkingArea.Activities = activities

	if body.Service.ServiceName != "" {
		kyc.Service.ServiceName = models.ServiceType(body.Service.ServiceName)
	}

	result = initializers.DB.Save(&kyc)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update KYC"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "KYC updated successfully", "kyc": kyc})
}

type UpdateKYCRequest struct {
	FullName       string                      `json:"full_name"`
	MobileNumber   string                      `json:"mobile_number"`
	FirmRegistered bool                        `json:"firm_registered"`
	Address        UpdateKYCAddressRequest     `json:"address"`
	WorkingArea    UpdateKYCWorkingAreaRequest `json:"working_area"`
	Service        UpdateKYCServiceRequest     `json:"service"`
}

type UpdateKYCAddressRequest struct {
	Province     string `json:"province"`
	District     string `json:"district"`
	Municipality string `json:"municipality"`
	WardNumber   string `json:"ward_number"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

type UpdateKYCWorkingAreaRequest struct {
	AreaName   string                `json:"area_name"`
	Activities []UpdateActivityItems `json:"activities"`
}

type UpdateActivityItems struct {
	ActivityName string   `json:"activity_name"`
	Items        []string `json:"items"`
}

type UpdateKYCServiceRequest struct {
	ServiceName string `json:"service_name"`
}
