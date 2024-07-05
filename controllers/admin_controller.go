package controllers

import (
	"Basic-Trade-API/helpers"
	"Basic-Trade-API/models"
	"Basic-Trade-API/pkg/config"
	"Basic-Trade-API/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	DB *gorm.DB
}

func (ac *AuthController) RegisterAdmin(c *gin.Context) {
	var input models.Admin

	//  check status or validate json of request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hashing password
	input.Password = helpers.HashPass(input.Password)

	err := config.DB.Create(&input).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (ac *AuthController) LoginAdmin(c *gin.Context) {
	var input struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	//  check status or validate json of request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admin
	// validate email
	err := config.DB.Where("email = ?", input.Email).First(&admin).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate password between db and input user
	if !helpers.ComparePass(input.Password, admin.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
		return
	}

	// generate JWT Token
	token, err := utils.GenerateToken(admin.Email, admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
