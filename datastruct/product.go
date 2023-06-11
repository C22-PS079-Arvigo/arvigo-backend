package datastruct

import "time"

type (
	Product struct {
		ID                   uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name                 string    `gorm:"column:name" json:"name"`
		Description          string    `gorm:"column:description" json:"description"`
		Images               string    `gorm:"column:images" json:"images"`
		LinkExternal         string    `gorm:"column:link_external" json:"link_external"`
		CategoryID           uint64    `gorm:"column:category_id" json:"category_id"`
		BrandID              uint64    `gorm:"column:brand_id" json:"brand_id"`
		MerchantID           uint64    `gorm:"column:merchant_id" json:"merchant_id"`
		Status               string    `gorm:"column:status" json:"status"`
		IsSubscriptionActive bool      `gorm:"column:is_subscription_active" json:"is_subscription_active"`
		RejectedNote         string    `gorm:"column:rejected_note" json:"rejected_note"`
		Price                float64   `gorm:"column:price" json:"price"`
		CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	// DetailProductCategory struct {
	// 	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	// 	ProductID  int       `gorm:"column:product_id" json:"product_id"`
	// 	CategoryID int       `gorm:"column:category_id" json:"category_id"`
	// 	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	// 	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	// }

	// DetailProductBrand struct {
	// 	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	// 	ProductID int       `gorm:"column:product_id" json:"product_id"`
	// 	BrandID   int       `gorm:"column:brand_id" json:"brand_id"`
	// 	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	// 	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	// }

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
		IsPrimaryVariant bool      `gorm:"column:is_primary_variant" json:"is_primary_variant"`
		ProductID        uint64    `gorm:"column:product_id" json:"product_id"`
		CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductTag struct {
		ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		TagID     uint64    `gorm:"column:tag_id" json:"tag_id"`
		ProductID uint64    `gorm:"column:product_id" json:"product_id"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductMarketplace struct {
		ID              uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		MarketplaceID   uint64    `gorm:"column:marketplace_id" json:"marketplace_id"`
		ParentProductID uint64    `gorm:"column:parent_product_id" json:"parent_product_id"`
		ProductID       uint64    `gorm:"column:product_id" json:"product_id"`
		AddressID       uint64    `gorm:"column:addresses_id" json:"addresses_id"`
		Link            string    `gorm:"column:link" json:"link"`
		Clicked         uint64    `gorm:"column:clicked" json:"clicked"`
		CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailLinkedProduct struct {
		ID                uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		InitialProductID  uint64    `gorm:"column:initial_product_id" json:"initial_product_id"`
		MerchantProductID uint64    `gorm:"column:merchant_product_id" json:"merchant_product_id"`
		MerchantID        uint64    `gorm:"column:merchant_id" json:"merchant_id"`
		CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	DetailProductMarketplaceClicked struct {
		ID                         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		DetailProductMarketplaceID uint64    `gorm:"column:detail_product_marketplaces" json:"detail_product_marketplaces"`
		MerchantID                 uint64    `gorm:"column:merchant_id" json:"merchant_id"`
		UserID                     uint64    `gorm:"column:user_id" json:"user_id"`
		CreatedAt                  time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt                  time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	ProductFromML struct {
		ID string `json:"id"`
	}
)

func (Product) TableName() string {
	return "products"
}

// func (DetailProductCategory) TableName() string {
// 	return "detail_product_categories"
// }

// func (DetailProductReview) TableName() string {
// 	return "detail_product_reviews"
// }

// func (DetailProductBrand) TableName() string {
// 	return "detail_product_brands"
// }

func (DetailProductVariant) TableName() string {
	return "detail_product_variants"
}

func (DetailProductTag) TableName() string {
	return "detail_product_tags"
}

func (DetailProductMarketplace) TableName() string {
	return "detail_product_marketplaces"
}

func (DetailLinkedProduct) TableName() string {
	return "detail_linked_products"
}
