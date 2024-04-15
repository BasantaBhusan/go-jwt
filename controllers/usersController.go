package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/BasantaBhusan/go-jwt/initializers"
	"github.com/BasantaBhusan/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Create new user
// @Description Create a new user.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body UserRequest true "User information"
// @Success 200 {object} SuccessResponse "User created Successfully"
// @Router /user/signup [post]
func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary User login
// @Description Log in a user.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User credentials"
// @Success 200 {object} SuccessResponse "User created Successfully"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email =?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		// "token": tokenString,
	})
}

// @Summary Validate user
// @Description Validate User.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 "Ok"
// @Failure 401 "Unauthorized"
// @Router /user/validate [get]
func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// @Summary Logout user
// @Description Clear Cookie.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 "Sucessfully logged out."
// @Router /logout [get]
func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// @Summary Get all users
// @Description Retrieve all users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} UserResponse "List of users"
// @Router /user/all [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	// initializers.DB.Find(&users)
	// initializers.DB.Where("is_kyc = ?", true).Preload("Kyc").Find(&users)
	initializers.DB.Where("is_kyc = ?", true).Preload("Kyc").Preload("Kyc.Address").Preload("Kyc.WorkingArea").Preload("Kyc.WorkingArea.Activities").Preload("Kyc.Service").Find(&users)
	c.JSON(http.StatusOK, users)
}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
