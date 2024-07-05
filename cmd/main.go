package main

import (
	"Basic-Trade-API/controllers"
	"Basic-Trade-API/middlewares"
	"Basic-Trade-API/models"
	"Basic-Trade-API/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Load environment variables from .env file
	load := godotenv.Load()
	if load != nil {
		log.Fatalf("Error loading .env file: %v", load)
	}

	// Call configuration to connect database
	env, err := config.InitialAllConfig()
	if err != nil {
		log.Panic("error configuration", err)
	}

	// Sql connection
	databaseSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", env.MySqlName, env.MySqlPass, env.MySqlHost, env.MySqlPort, env.MySqlDbName)
	dbConnection, err := gorm.Open(mysql.Open(databaseSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("error failed to connect to database, %v", err)
	}

	dbConnection.AutoMigrate(&models.Admin{}, &models.Product{}, &models.Product{})

	// create controller
	r := gin.Default()

	adminController := &controllers.AuthController{DB: dbConnection}
	productController := &controllers.ProductController{DB: dbConnection}
	variantController := &controllers.VariantController{DB: dbConnection}

	r.POST("/auth/register", adminController.RegisterAdmin)
	r.POST("auth//login", adminController.LoginAdmin)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.POST("/products", productController.CreateProduct)                // Create Product
		authorized.GET("/products", productController.GetProduct)                    // Get All Product
		authorized.GET("/products/:productUUID", productController.GetProduct)       // Get Product by UUID
		authorized.PUT("/products/:productUUID", productController.UpdateProduct)    // Update Product by UUID
		authorized.DELETE("/products/:productUUID", productController.DeleteProduct) // Delete Product by UUID

		authorized.POST("/variants", variantController.CreateVariant)                // Create Variant
		authorized.GET("/variants", variantController.GetVariant)                    // Get all Variant
		authorized.GET("/variants/:variantUUID", variantController.GetVariant)       // Get Variant by UUID
		authorized.PUT("/variants/:variantUUID", variantController.UpdateVariant)    // Update Variant by UUID
		authorized.DELETE("/variants/:variantUUID", variantController.DeleteVariant) //Delete Variant by UUID
	}
	r.Run("[::]:8080")
}
