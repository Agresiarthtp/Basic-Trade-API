package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID      uuid.UUID `gorm:"type:char(36);not null" json:"uuid"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	ImageURL  string    `gorm:"type:varchar(255);not null" json:"image_url"`
	AdminID   []Admin   `gorm:"type:char(36);not null" json:"admin_id"`
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
