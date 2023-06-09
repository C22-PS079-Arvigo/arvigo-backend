package datastruct

import "time"

type Marketplace struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Image     string    `gorm:"column:image" json:"image"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type MerchantMarketplace struct {
	Name    string `gorm:"column:name" json:"name"`
	Clicked uint64 `gorm:"column:clicked" json:"clicked"`
}

func (Marketplace) TableName() string {
	return "marketplaces"
}
