package dto

import "github.com/google/uuid"

type VariantDTO struct {
	ID          uint64 `json:"id"`
	UUID        string `json:"uuid"`
	VariantName string `json:"variant_name"`
	Quantity    int    `json:"quantity"`
	ProductID   string `json:"product_id"`
}

type ProductDTO struct {
	ID       uint64 `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	AdminID  string `json:"admin_id"`
}

type AdminDTO struct {
	ID    uint64 `json:"id"`
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ConvertStringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func ConvertUUIDToString(u uuid.UUID) string {
	return u.String()
}
