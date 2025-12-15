package models

type Service struct {
	ServiceID string `gorm:"primaryKey" json:"service_id"`
	Name      string `json:"name"`
	Price     float32 `gorm:"type:numeric" json:"price"`
}
