package datastruct

import "time"

type Merchant struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name             string    `gorm:"column:name" json:"name"`
	IsRecommendation int       `gorm:"column:is_recomendation" json:"is_recommendation"`
	AddressID        int       `gorm:"column:addresses_id" json:"address_id"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Merchant) TableName() string {
	return "merchants"
}
