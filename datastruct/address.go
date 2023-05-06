package datastruct

import "time"

type Address struct {
	ID            uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID        uint64    `gorm:"column:user_id" json:"user_id"`
	Street        string    `gorm:"column:street" json:"street"`
	ProvinceID    uint64    `gorm:"column:province_id" json:"province_id"`
	DistrictID    uint64    `gorm:"column:district_id" json:"district_id"`
	SubdistrictID uint64    `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	PostalCodeID  uint64    `gorm:"column:postal_code_id" json:"postal_code_id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Address) TableName() string {
	return "addresses"
}
