package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Variant struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID        uuid.UUID `gorm:"primary_key;type:char(36);not null" json:"uuid"`
	VariantName string    `json:"variant_name"`
	Quantity    int       `json:"quantity"`
	ProductID   []Product `gorm:"type:char(36);not null" json:"product_id"`
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
