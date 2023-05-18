package datastruct

import "time"

type Wishlist struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	ProductID  int       `gorm:"column:product_id" json:"product_id"`
	IsMerchant bool      `gorm:"column:is_merchant" json:"is_merchant"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Wishlist) TableName() string {
	return "wishlists"
}
