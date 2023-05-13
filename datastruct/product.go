package datastruct

import "time"

type (
	Product struct {
		ID           uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name         string    `gorm:"column:name" json:"name"`
		Descripion   string    `gorm:"column:description" json:"description"`
		Images       string    `gorm:"column:images" json:"images"`
		LinkExternal string    `gorm:"column:link_external" json:"link_external"`
		MerchantID   int       `gorm:"column:merchant_id" json:"merchant_id"`
		CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductCategory struct {
		ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		ProductID  int       `gorm:"column:product_id" json:"product_id"`
		CategoryID int       `gorm:"column:category_id" json:"category_id"`
		CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductBrand struct {
		ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		ProductID int       `gorm:"column:product_id" json:"product_id"`
		BrandID   int       `gorm:"column:brand_id" json:"brand_id"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	// DetailProductReview struct {
	// 	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	// 	ProductID int       `gorm:"column:product_id" json:"product_id"`
	// 	UserID    int       `gorm:"column:user_id" json:"user_id"`
	// 	Comment   string    `gorm:"column:comment" json:"comment"`
	// 	Rating    float64   `gorm:"column:rating;default:0" json:"rating"`
	// 	Images    string    `gorm:"column:images" json:"images"`
	// 	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	// 	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// }

	DetailProductVariant struct {
		ID               uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name             string    `gorm:"column:name" json:"name"`
		LinkAR           string    `gorm:"column:link_ar" json:"link_ar"`
		IsPrimaryVariant int       `gorm:"column:is_primary_variant" json:"is_primary_variant"`
		ProductID        int       `gorm:"column:product_id" json:"product_id"`
		CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductTag struct {
		ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		TagID     string    `gorm:"column:tag_id" json:"tag_id"`
		ProductID int       `gorm:"column:product_id" json:"product_id"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductMarketplace struct {
		ID            uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		MarketplaceID int       `gorm:"column:marketplace_id" json:"marketplace_id"`
		ProductID     int       `gorm:"column:product_id" json:"product_id"`
		Link          string    `gorm:"column:link" json:"link"`
		Clicked       int       `gorm:"column:clicked" json:"clicked"`
		CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	}
)

func (Product) TableName() string {
	return "products"
}

func (DetailProductCategory) TableName() string {
	return "detail_product_categories"
}

// func (DetailProductReview) TableName() string {
// 	return "detail_product_reviews"
// }

func (DetailProductBrand) TableName() string {
	return "detail_product_brands"
}

func (DetailProductVariant) TableName() string {
	return "detail_product_variants"
}

func (DetailProductTag) TableName() string {
	return "detail_product_tags"
}

func (DetailProductMarketplace) TableName() string {
	return "detail_product_marketplaces"
}
