package datastruct

import "time"

type Wishlist struct {
	ID                         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID                     uint64    `gorm:"column:user_id" json:"user_id"`
	ProductID                  *uint64   `gorm:"column:product_id" json:"product_id"`
	DetailProductMarketplaceID *uint64   `gorm:"column:detail_product_marketplace_id" json:"detail_product_marketplace_id"`
	CreatedAt                  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Wishlist) TableName() string {
	return "wishlists"
}
