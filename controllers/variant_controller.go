package controllers

import (
	"Basic-Trade-API/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type VariantController struct {
	DB *gorm.DB
}

func (vc *VariantController) CreateVariant(c *gin.Context) {
	var input models.Variant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := vc.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func (vc *VariantController) GetVariant(c *gin.Context) {
	var variants []models.Variant
	if err := vc.DB.Find(&variants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// add handler to get variant by UUID
	if err := vc.DB.First(&variants, "uuid = ?", c.Param("variantUUID")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": variants})
}

func (vc *VariantController) UpdateVariant(c *gin.Context) {
	var input models.Variant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var variant models.Variant
	if err := vc.DB.First(&variant, input.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}

	if err := vc.DB.Model(&variant).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": variant})
}

func (vc *VariantController) DeleteVariant(c *gin.Context) {
	var variant models.Variant
	if err := vc.DB.First(&variant, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variant not found"})
		return
	}

	if err := vc.DB.Delete(&variant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
