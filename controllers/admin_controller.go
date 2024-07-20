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

var (
	appJSON = "application/json"
)

func (ac *AuthController) RegisterAdmin(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	Admin := models.Admin{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	// validate if the email already exist
	var existingAdmin models.Admin
	err := ac.DB.Where("email = ?", Admin.Email).First(&existingAdmin).Error
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Bad Request!",
			"message": "Email already in use",
		})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":   "Internal Server Error!",
			"message": "Email already in use",
		})
	}

	// Generate a new UUID
	newUUID := uuid.New()
	Admin.UUID = newUUID.String() // set the generated UUID as the ID

	err = ac.DB.Create(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request!",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Admin,
	})
}

func (ac *AuthController) LoginAdmin(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	Admin := models.Admin{}
	var password string

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	password = Admin.Password

	err := ac.DB.Where("email = ?", Admin.Email).First(&Admin).Error
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized, gin.H{
				"Error":   "Unauthorized",
				"Message": "Invalid Email.",
			})
		return
	}

	comparePass := helpers.ComparePass([]byte(Admin.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error":   "Unauthorized!",
			"Message": "Invalid Password.",
		})
		return
	}

	token := utils.GenerateToken(Admin.Email, Admin.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Merespon dengan token JWT
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
