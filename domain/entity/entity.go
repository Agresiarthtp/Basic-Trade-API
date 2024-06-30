package entity

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
	ProductID   uuid.UUID `gorm:"type:char(36);not null" json:"product_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Product struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID      uuid.UUID `gorm:"type:char(36);not null" json:"uuid"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	ImageURL  string    `gorm:"type:varchar(255);not null" json:"image_url"`
	AdminID   uuid.UUID `gorm:"type:char(36);not null" json:"admin_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Admin struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID      uuid.UUID `gorm:"type:char(36);not null" json:"uuid"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (variant *Variant) BeforeCreate(trx *gorm.DB) (err error) {
	_, errorCreate := govalidator.ValidateStruct(variant)
	if errorCreate != nil {
		err = errorCreate
	}
	return
}

func (admin *Admin) BeforeCreate(trx *gorm.DB) (err error) {
	_, errorCreate := govalidator.ValidateStruct(admin)
	if errorCreate != nil {
		err = errorCreate
		return
	}
	return
}

func (product *Product) BeforeCreate(trx *gorm.DB) (err error) {
	_, errorCreate := govalidator.ValidateStruct(product)
	if errorCreate != nil {
		err = errorCreate
		return
	}
	return
}
