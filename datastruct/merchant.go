package datastruct

import "time"

type Merchant struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name             string    `gorm:"column:name" json:"name"`
	IsRecommendation int       `gorm:"column:is_recomendation" json:"is_recommendation"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type MerchantProduct struct {
	ID      uint64  `gorm:"column:id" json:"id"`
	Images  string  `gorm:"column:images" json:"image"`
	Name    string  `gorm:"column:name" json:"name"`
	Price   float64 `gorm:"column:price" json:"price"`
	Status  string  `gorm:"column:status" json:"status"`
	Clicked uint64  `gorm:"column:clicked" json:"clicked"`
}

type MerchantVisitor struct {
	Today     uint64 `gorm:"column:today" json:"today"`
	ThisMonth uint64 `gorm:"column:this_month" json:"this_month"`
	LastMonth uint64 `gorm:"column:last_month" json:"last_month"`
}

func (Merchant) TableName() string {
	return "merchants"
}
