package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Variant struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID        string    `gorm:"primary_key;type:char(36);not null" json:"uuid"`
	VariantName string    `json:"variant_name" valid:"required~Variant name is required"`
	Quantity    int       `json:"quantity" valid:"required~Quantity is required,int~Quantity must be an integer"`
	ProductID   []Product `gorm:"type:char(36);not null" json:"product_id" valid:"required~Product ID is required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (variant *Variant) BeforeCreate(trx *gorm.DB) (err error) {
	_, errorCreate := govalidator.ValidateStruct(variant)
	if errorCreate != nil {
		err = errorCreate
	}
	return
}
