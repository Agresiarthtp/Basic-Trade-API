package controllers

import (
	"Basic-Trade-API/helpers"
	"Basic-Trade-API/models"
	"Basic-Trade-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	DB *gorm.DB
}

func (ac *AuthController) RegisterAdmin(c *gin.Context) {
	var input models.Admin

	// Check and parse JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new admin with auto-generated ID
	input.Password = helpers.HashPass(input.Password)

	admin := models.Admin{
		Name:     input.Name,
		Email:    input.Email,
		Password: helpers.HashPass(input.Password),
		UUID:     uuid.New(), // Auto-generate a UUID for the new admin
		// Hash the password before storing it
	}

	// Create new admin record in the database
	if err := ac.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (ac *AuthController) LoginAdmin(c *gin.Context) {
	var input struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// Check and parse JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admin
	// Query admin record by email
	if err := ac.DB.Where("email = ?", input.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate password
	if !helpers.ComparePass(input.Password, admin.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
		return
	}

	// Generate JWT Token
	token, err := utils.GenerateToken(admin.Email, admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
