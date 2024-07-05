package utils

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitCloudinary() *cloudinary.Cloudinary {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary, %v", err)
	}

	return cld
}

func UploadToCloudinary(filePath string) (string, error) {
	cld := InitCloudinary()
	ctx := context.Background()

	uploadResult, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return uploadResult.SecureURL, nil
}
