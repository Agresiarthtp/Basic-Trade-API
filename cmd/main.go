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

	adminController := &controllers.AuthController{DB: dbConnection}
	productController := &controllers.ProductController{DB: dbConnection}
	variantController := &controllers.VariantController{DB: dbConnection}

	// create controller
	router := gin.Default()
	userRouter := router.Group("/auth")
	{
		userRouter.POST("/register", adminController.RegisterAdmin)
		userRouter.POST("/login", adminController.LoginAdmin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", productController.GetProduct) // Get All Product
		productRouter.Use(middlewares.Authentication())

		productRouter.POST("/", productController.CreateProduct)               // Create Product
		productRouter.GET("/:productUUID", productController.GetProduct)       // Get Product by UUID
		productRouter.PUT("/:productUUID", productController.UpdateProduct)    // Update Product by UUID
		productRouter.DELETE("/:productUUID", productController.DeleteProduct) // Delete Product by UUID
	}

	variantRouter := router.Group("/variants")
	{
		variantRouter.POST("/", variantController.CreateVariant)               // Create Variant
		variantRouter.GET("/", variantController.GetVariant)                   // Get all Variant
		variantRouter.GET("/:variantUUID", variantController.GetVariant)       // Get Variant by UUID
		variantRouter.PUT("/:variantUUID", variantController.UpdateVariant)    // Update Variant by UUID
		variantRouter.DELETE("/:variantUUID", variantController.DeleteVariant) //Delete Variant by UUID
		router.Run("[::]:8080")
	}
}
