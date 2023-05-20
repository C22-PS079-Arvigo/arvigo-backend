package repository

import (
	"net/http"
	"time"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func AddWhislistProduct(userID uint64, data datastruct.AddWhislistProductInput) (statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK
	currentTime := time.Now()

	wishlistPayload := datastruct.Wishlist{
		UserID:                     userID,
		ProductID:                  data.ProductID,
		DetailProductMarketplaceID: data.DetailProductMarketplaceID,
		CreatedAt:                  currentTime,
		UpdatedAt:                  currentTime,
	}

	if err = db.Create(&wishlistPayload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return
}
