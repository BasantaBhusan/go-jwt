// controllers/search.go

package controllers

import (
	"net/http"
	"strings"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform search"})
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
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search query"})
		return
	}

	var results []models.Kyc

	// Perform case-insensitive search query across all models
	if err := initializers.DB.
		Joins("JOIN users ON kycs.user_id = users.id").
		Joins("LEFT JOIN services ON kycs.id = services.kyc_id").
		Joins("LEFT JOIN addresses ON kycs.id = addresses.kyc_id").
		Where("LOWER(kycs.full_name) LIKE ? OR LOWER(users.email) LIKE ? OR LOWER(services.service_name) LIKE ? OR LOWER(addresses.province) LIKE ? OR LOWER(addresses.district) LIKE ? OR LOWER(addresses.municipality) LIKE ? OR LOWER(addresses.ward_number) LIKE ?",
			"%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%").
		Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform search"})
		return
	}

	for i := range results {
		err := initializers.DB.Model(&results[i]).Association("Address").Find(&results[i].Address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load addresses"})
			return
		}

		err = initializers.DB.Model(&results[i]).Association("WorkingArea").Find(&results[i].WorkingArea)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load working areas"})
			return
		}

		err = initializers.DB.Model(&results[i].WorkingArea).Association("Activities").Find(&results[i].WorkingArea.Activities)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load activities"})
			return
		}

		err = initializers.DB.Model(&results[i]).Association("Service").Find(&results[i].Service)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load services"})
			return
		}
	}

	c.JSON(http.StatusOK, results)
}
