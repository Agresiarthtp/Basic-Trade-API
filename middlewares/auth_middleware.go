package middlewares

import (
	"Basic-Trade-API/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Memeriksa apakah authHeader memiliki format Bearer <token> menggunakan strings.Split
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		//Jika format header salah, memberikan respons Unauthorized.
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		//  to validate token and get claims
		// Menyimpan userID yang diambil dari claim token ke dalam context untuk digunakan di endpoint selanjutnya
		c.Set("userData", claims.ID)
		c.Next()
	}
}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := utils.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UnAuthorized",
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
