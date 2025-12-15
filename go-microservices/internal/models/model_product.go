package models

type Product struct {
	ProductID string  `gorm:"primaryKey" json:"product_id"`
	VendorID  string  `json:"vendor_id"`
	Name      string  `json:"name"`
	Price     float32 `gorm:"type:numeric" json:"price"`
}
