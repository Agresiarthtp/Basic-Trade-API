package models

import (
	"Basic-Trade-API/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	UUID      string    `gorm:"type:char(36);not null" json:"uuid""`
	Name      string    `gorm:"type:varchar(255);not null" json:"name" valid:"required~Name is required"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email" valid:"email~Invalid email format"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password" valid:"required~Password is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Admin) BeforeCreate(trx *gorm.DB) (err error) {
	_, errorCreate := govalidator.ValidateStruct(admin)
	if errorCreate != nil {
		err = errorCreate
		return
	}

	// create admin/regist do hashing pass
	admin.Password = helpers.HashPass(admin.Password)
	err = nil
	return
}
