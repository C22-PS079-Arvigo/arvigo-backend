package datastruct

import "mime/multipart"

type (
	LoginUserInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	UserRegisterInput struct {
		Email                string `json:"email" validate:"required,email"`
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required"`
		FullName             string `json:"full_name" validate:"required"`
	}

	PartnerRegisterInput struct {
		StoreName            string `json:"store_name" validate:"required"`
		Email                string `json:"email" validate:"required,email"`
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required"`
		Street               string `json:"street" validate:"required"`
		ProvinceID           uint64 `json:"province_id" validate:"required,numeric"`
		CityID               uint64 `json:"city_id" validate:"required,numeric"`
		DistrictID           uint64 `json:"district_id" validate:"required,numeric"`
		SubdistrictID        uint64 `json:"subdistrict_id" validate:"required,numeric"`
		PostalCodeID         uint64 `json:"postal_code_id" validate:"required,numeric"`
	}

	CreateInitialProductInput struct {
		Name                  string                  `form:"name" validate:"required"`
		Description           string                  `form:"description"`
		Images                []*multipart.FileHeader `form:"images"`
		LinkExternal          string                  `form:"link_external"`
		CategoryID            uint64                  `form:"category_id" validate:"required"`
		BrandID               uint64                  `form:"brand_id" validate:"required"`
		MerchantID            uint64                  `form:"merchant_id"`
		DetailProductTags     string                  `form:"detail_product_tags" validate:"required"`
		DetailProductVariants string                  `form:"detail_product_variants" validate:"required"`
	}

	CreateMerchantProductInput struct {
		ProductID                uint64                  `form:"product_id" validate:"required"`
		Name                     string                  `form:"name" validate:"required"`
		Description              string                  `form:"description"`
		Images                   []*multipart.FileHeader `form:"images"`
		MerchantID               uint64                  `form:"merchant_id"`
		DetailProductMarketplace string                  `form:"detail_product_marketplaces" validate:"required"`
		Price                    float64                 `form:"price" validate:"required"`
	}

	VerifyProductInput struct {
		ProductID    uint64 `json:"product_id" validate:"required"`
		Status       string `json:"status" validate:"required"`
		RejectedNote string `json:"rejected_note"`
	}

	UpdateProductInput struct {
		ProductID   uint64  `json:"product_id" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
		Description string  `json:"description" validate:"required"`
	}

	BrandInput struct {
		Name       string                `form:"column:name" json:"name"`
		Image      *multipart.FileHeader `form:"column:image" json:"image"`
		CategoryID uint64                `form:"column:category_id" json:"category_id"`
	}

	AddWhislistProductInput struct {
		ProductID                  *uint64 `json:"product_id"`
		DetailProductMarketplaceID *uint64 `json:"detail_product_marketplace_id"`
	}
)
