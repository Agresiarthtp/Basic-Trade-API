package entity

import (
	"Basic-Trade-API/domain/dto"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

var validate = validator.New()

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

func NewAdmin(dto dto.AdminDTO) (*Admin, error) {
	adminUUID, err := uuid.Parse(dto.UUID)
	if err != nil {
		return nil, err
	}

	a := &Admin{
		ID:        0,
		UUID:      adminUUID,
		Name:      dto.Name,
		Email:     dto.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = a.ValidateAdminStruct()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func NewProduct(dto dto.ProductDTO) (*Product, error) {
	adminID, err := uuid.Parse(dto.AdminID)
	if err != nil {
		return nil, err
	}

	productUUID, err := uuid.Parse(dto.UUID)
	if err != nil {
		return nil, err
	}

	a := &Product{
		ID:        0,
		UUID:      productUUID,
		Name:      dto.Name,
		ImageURL:  dto.ImageURL,
		AdminID:   adminID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = a.ValidateProductStruct()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func NewVariant(dto dto.VariantDTO) (*Variant, error) {
	variantId, err := uuid.Parse(dto.UUID)
	if err != nil {
		return nil, err
	}

	productId, err := uuid.Parse(dto.ProductID)
	if err != nil {
		return nil, err
	}

	a := &Variant{
		ID:          0,
		UUID:        variantId,
		VariantName: dto.VariantName,
		Quantity:    dto.Quantity,
		ProductID:   productId,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	err = a.ValidateVariantStruct()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Admin) ValidateAdminStruct() error {
	return validate.Struct(a)
}

func (a *Product) ValidateProductStruct() error {
	return validate.Struct(a)
}

func (a *Variant) ValidateVariantStruct() error {
	return validate.Struct(a)
}
