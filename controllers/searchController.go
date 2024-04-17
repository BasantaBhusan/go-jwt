package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Search handles search requests
// @Summary Perform a search
// @Description Search for users by email
// @Tags Search
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} models.User "List of users matching the search query"
// @Router /search [get]
func Search(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search query"})
		return
	}

	var users []models.User
	if err := initializers.DB.Where("email LIKE ?", "%"+query+"%").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Search handles global search requests across all models
// @Summary Perform a global search across all models
// @Description Perform a global search across all models
// @Tags Search
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Router /search/all [get]
func GlobalSearch(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search query"})
		return
	}

	var results []models.Kyc

	if err := searchModels(query).Preload("Address").
		Preload("WorkingArea").Preload("WorkingArea.Activities").
		Preload("Service").Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform global search"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func searchModels(query string) *gorm.DB {
	return initializers.DB.
		Joins("JOIN users ON kycs.user_id = users.id").
		Joins("LEFT JOIN services ON kycs.id = services.kyc_id").
		Joins("LEFT JOIN addresses ON kycs.id = addresses.kyc_id").
		Joins("LEFT JOIN working_areas ON kycs.id = working_areas.kyc_id").
		Joins("LEFT JOIN activities ON working_areas.id = activities.working_area_id").
		Where("LOWER(kycs.full_name) LIKE ? OR LOWER(users.email) LIKE ? OR LOWER(working_areas.area_name) LIKE ? OR LOWER(activities.activity_name) LIKE ? OR LOWER(services.service_name) LIKE ? OR LOWER(addresses.province) LIKE ? OR LOWER(addresses.district) LIKE ? OR LOWER(addresses.municipality) LIKE ? OR LOWER(addresses.ward_number) LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")
}

// AdvancedSearch handles advanced global search requests across all models
// @Summary Perform an advanced global search across all models
// @Description Perform an advanced global search across all models based on the provided query string
// @Tags Search
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} models.Kyc "List of results matching the advanced search query"
// @Router /search/advanced [get]
func AdvancedSearch(c *gin.Context) {
	query := strings.ToLower(c.Query("q"))

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search query"})
		return
	}

	var results []models.Kyc

	// Split the query into terms
	terms := strings.Split(query, " ")

	// Build the dynamic query
	dbQuery := buildAdvancedQuery(terms)

	// Execute the query
	if err := dbQuery.Preload("Address").
		Preload("WorkingArea").Preload("WorkingArea.Activities").
		Preload("Service").Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform advanced search"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// buildAdvancedQuery builds a dynamic query based on the provided search terms
func buildAdvancedQuery(terms []string) *gorm.DB {
	dbQuery := initializers.DB.
		Joins("JOIN users ON kycs.user_id = users.id").
		Joins("LEFT JOIN services ON kycs.id = services.kyc_id").
		Joins("LEFT JOIN addresses ON kycs.id = addresses.kyc_id").
		Joins("LEFT JOIN working_areas ON kycs.id = working_areas.kyc_id").
		Joins("LEFT JOIN activities ON working_areas.id = activities.working_area_id")

	for _, term := range terms {
		// Add conditions for each term
		term = "%" + term + "%"
		dbQuery = dbQuery.Where("LOWER(kycs.full_name) LIKE ? OR LOWER(users.email) LIKE ? OR LOWER(working_areas.area_name) LIKE ? OR LOWER(activities.activity_name) LIKE ? OR LOWER(services.service_name) LIKE ? OR LOWER(addresses.province) LIKE ? OR LOWER(addresses.district) LIKE ? OR LOWER(addresses.municipality) LIKE ? OR LOWER(addresses.ward_number) LIKE ?",
			term, term, term, term, term, term, term, term, term)
	}

	return dbQuery
}

// AddressSearch handles search requests based on the address model
// and returns associated working area, activities, and services
// @Summary Perform a search based on the address model
// @Description Search based on the address model and return associated working area, activities, and services
// @Tags Search
// @Accept json
// @Produce json
// @Param body body AddressSearchRequest true "Search request"
// @Success 200 {object} SearchResult "Search result"
// @Router /search/address [post]
func AddressSearch(c *gin.Context) {
	var reqBody AddressSearchRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	addressQuery := reqBody.Address

	// Search for the address
	var address models.Address
	if err := initializers.DB.Where("LOWER(province) LIKE ? OR LOWER(district) LIKE ? OR LOWER(municipality) LIKE ? OR LOWER(ward_number) LIKE ?",
		"%"+strings.ToLower(addressQuery)+"%", "%"+strings.ToLower(addressQuery)+"%", "%"+strings.ToLower(addressQuery)+"%", "%"+strings.ToLower(addressQuery)+"%").
		First(&address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search address"})
		}
		return
	}

	// Find the working area based on the address information
	var workingArea models.WorkingArea
	if err := initializers.DB.Where("area_name LIKE ?", "%"+strings.ToLower(addressQuery)+"%").First(&workingArea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Working area not found for the provided address"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve working area"})
		}
		return
	}

	// Retrieve activities associated with the working area
	var activities []models.Activity
	if err := initializers.DB.Where("working_area_id = ?", workingArea.ID).Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
		return
	}

	// Retrieve services associated with the activities
	var services []models.Service
	if err := initializers.DB.Joins("JOIN activities ON services.activity_id = activities.id").
		Where("activities.id IN (?)", activities).Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}

	// Construct the response
	result := SearchResult{
		Address:            address,
		WorkingArea:        workingArea,
		Activities:         activities,
		AssociatedServices: services,
	}

	c.JSON(http.StatusOK, result)
}

// SearchResult represents the search result containing address, working area, activities, and associated services
type SearchResult struct {
	Address            models.Address     `json:"address"`
	WorkingArea        models.WorkingArea `json:"working_area"`
	Activities         []models.Activity  `json:"activities"`
	AssociatedServices []models.Service   `json:"associated_services"`
}

// AddressSearchRequest represents the request body for AddressSearch
type AddressSearchRequest struct {
	Address string `json:"address" binding:"required"`
}
