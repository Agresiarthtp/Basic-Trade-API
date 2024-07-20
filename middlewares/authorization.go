package middlewares

import (
	"Basic-Trade-API/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
)

type ProductAuth struct {
	DB *gorm.DB
}

func (d *ProductAuth) ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productUUID := ctx.Param("productUUID")

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		var getProduct models.Product
		err := d.DB.Where("id = ?", productUUID).First(&getProduct).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}

		if getProduct.ID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
