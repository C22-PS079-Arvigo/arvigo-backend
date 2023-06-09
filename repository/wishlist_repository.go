package repository

import (
	"net/http"
	"strings"
	"time"

	"github.com/yusufwib/arvigo-backend/constant"
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

func DeleteWhislistProduct(userID uint64, data datastruct.AddWhislistProductInput) (statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var wishlist datastruct.Wishlist
	if data.ProductID != nil {
		if err = db.Where("product_id = ? AND user_id = ?", data.ProductID, userID).Delete(&wishlist).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	} else if data.DetailProductMarketplaceID != nil {
		if err = db.Where("detail_product_marketplace_id AND user_id = ?", data.DetailProductMarketplaceID, userID).Delete(&wishlist).Error; err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return
}

func GetUserWishlist(userID uint64) (res datastruct.UserWishlistResponse, statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db                           = Database()
		wishlists                    []datastruct.Wishlist
		initProductID                []uint64
		merchantProductMarketplaceID []uint64

		initProduct     []datastruct.ProductWishlist
		merchantProduct []datastruct.ProductMarketplaceWishlist
	)

	if err = db.Table("wishlists").
		Select("product_id, detail_product_marketplace_id").
		Where("user_id", userID).
		Find(&wishlists).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, v := range wishlists {
		if v.ProductID != nil {
			initProductID = append(initProductID, *v.ProductID)
		} else if v.DetailProductMarketplaceID != nil {
			merchantProductMarketplaceID = append(merchantProductMarketplaceID, *v.DetailProductMarketplaceID)
		}
	}

	if err = db.Table("products p").
		Select("p.id, p.name, p.images, b.name as brand").
		Joins("join brands b on b.id = p.brand_id").
		Where("p.id in (?) AND merchant_id = 0", initProductID).
		Find(&initProduct).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range initProduct {
		initProduct[i].Image = strings.Split(v.Image, ",")[0]
	}

	if err = db.Table("detail_product_marketplaces").
		Select([]string{
			"detail_product_marketplaces.id AS id",
			"products.name",
			"brands.name AS brand",
			"products.images",
			"products.price",
			"merchants.name AS merchant",
			"detail_product_marketplaces.link AS marketplace_link",
			"detail_product_marketplaces.marketplace_id",
			"detail_product_marketplaces.addresses_id",
		}).
		Joins("LEFT JOIN products ON products.id = detail_product_marketplaces.product_id").
		Joins("LEFT JOIN brands ON brands.id = products.brand_id").
		Joins("LEFT JOIN merchants ON products.merchant_id = merchants.id").
		Where("detail_product_marketplaces.id IN (?)", merchantProductMarketplaceID).
		Find(&merchantProduct).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range merchantProduct {
		merchantProduct[i].Image = strings.Split(v.Image, ",")[0]

		if v.AddressID != 0 {
			merchantProduct[i].Type = "offline"
			addr, _, _, err := GetAddressByID(v.AddressID)
			if err == nil {
				merchantProduct[i].Address = &addr
			}
			continue
		}

		if v.MarketplaceID != 0 {
			merchantProduct[i].Type = "online"
			marketplaceName := constant.Marketplace[v.MarketplaceID]
			merchantProduct[i].Marketplace = &marketplaceName
		}
	}

	res = datastruct.UserWishlistResponse{
		Product: initProduct,
		Store:   merchantProduct,
	}
	return
}
