package datastruct

import "time"

type Brand struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	CategoryID uint64    `gorm:"column:category_id" json:"category_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Brand) TableName() string {
	return "brands"
}
