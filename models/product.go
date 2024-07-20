package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID      string    `gorm:"type:char(36);not null" json:"uuid"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name" valid:"required~Name is required"`
	ImageURL  string    `gorm:"type:varchar(255);not null" json:"image_url" valid:"required~Image URL is required,url~Invalid URL format"`
	AdminID   []Admin   `gorm:"type:char(36);not null" json:"admin_id" valid:"required~Admin ID is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (product *Product) BeforeCreate(trx *gorm.DB) (err error) {
	_, errorCreate := govalidator.ValidateStruct(product)
	if errorCreate != nil {
		err = errorCreate
		return
	}
	return
}
