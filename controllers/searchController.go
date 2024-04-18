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
func SearchByEmail(c *gin.Context) {
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

// GlobalSearch handles global search requests across all models
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
// @Param query query string true "Search query"
// @Success 200 {object} SearchResult "Search result"
// @Router /search/address [get]
func AddressSearch(c *gin.Context) {
	query := c.Query("query")

	// If the query is empty, return an error
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search query"})
		return
	}

	// Define conditions for the query based on provided query parameter
	conditions := "province ILIKE ? OR district ILIKE ? OR municipality ILIKE ? OR ward_number ILIKE ?"
	searchQuery := "%" + strings.ToLower(query) + "%"

	// Search for the addresses
	var addresses []models.Address
	if err := initializers.DB.Where(conditions, searchQuery, searchQuery, searchQuery, searchQuery).Find(&addresses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Addresses not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search addresses"})
			return
		}
	}

	// Initialize a slice to hold the results
	var results []SearchResult

	// Iterate over each address to retrieve associated data
	for _, address := range addresses {
		// Find the working area based on the address information
		var workingArea models.WorkingArea
		if err := initializers.DB.
			Joins("JOIN kycs ON working_areas.kyc_id = kycs.id").
			Joins("JOIN addresses ON kycs.id = addresses.kyc_id").
			Where("addresses.user_id = ? AND addresses.kyc_id = ?", address.UserID, address.KycID).
			Find(&workingArea).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Working area not found for the provided address"})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve working area"})
				return
			}
		}

		// Retrieve activities associated with the working area
		var activities []models.Activity
		if err := initializers.DB.Where("working_area_id = ?", workingArea.ID).Find(&activities).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
			return
		}

		// Retrieve services associated with the activities
		var services []models.Service
		if err := initializers.DB.
			Joins("JOIN activities ON activities.working_area_id = ?", workingArea.ID).
			Joins("JOIN working_areas ON activities.working_area_id = working_areas.id").
			Where("services.kyc_id = working_areas.kyc_id").
			Find(&services).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
			return
		}

		// Append the results to the slice
		result := SearchResult{
			Address:            address,
			WorkingArea:        workingArea,
			Activities:         activities,
			AssociatedServices: services,
		}
		results = append(results, result)
	}

	// Construct the response
	c.JSON(http.StatusOK, results)
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

// AllAddressSearch handles search requests based on the address model
// and returns associated working area, activities, and services
// @Summary Perform a search based on the address model
// @Description Search based on the address model and return associated working area, activities, and services
// @Tags Search
// @Accept json
// @Produce json
// @Param province path string true "Province"
// @Param district path string true "District"
// @Param municipality path string true "Municipality"
// @Param ward_number path string true "Ward Number"
// @Success 200 {object} SearchResult "Search result"
// @Router /search/all/address/{province}/{district}/{municipality}/{ward_number} [get]
func AllAddressSearch(c *gin.Context) {
	province := c.Param("province")
	district := c.Param("district")
	municipality := c.Param("municipality")
	wardNumber := c.Param("ward_number")

	// Search for the address
	var address models.Address
	if err := initializers.DB.Where("LOWER(province) = ? AND LOWER(district) = ? AND LOWER(municipality) = ? AND LOWER(ward_number) = ?",
		strings.ToLower(province), strings.ToLower(district), strings.ToLower(municipality), strings.ToLower(wardNumber)).
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
	if err := initializers.DB.
		Joins("JOIN kycs ON working_areas.kyc_id = kycs.id").
		Joins("JOIN addresses ON kycs.id = addresses.kyc_id").
		Where("addresses.user_id = ? AND addresses.kyc_id = ?", address.UserID, address.KycID).
		Find(&workingArea).Error; err != nil {
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
	if err := initializers.DB.
		Joins("JOIN activities ON activities.working_area_id = ?", workingArea.ID).
		Joins("JOIN working_areas ON activities.working_area_id = working_areas.id").
		Where("services.kyc_id = working_areas.kyc_id").
		Find(&services).Error; err != nil {
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
