package repository

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetMerchantAppHome(userID uint64) (res datastruct.MerchantHome, statusCode int, err error) {
	statusCode = http.StatusOK

	var (
		db               = Database()
		merchantProducts []datastruct.MerchantProduct
		visitors         datastruct.MerchantVisitor
		currentTime      = time.Now()
	)

	// list stores
	if err = db.Table("products p").
		Select("p.id, images, name, price, status, sum(dpm.clicked) as clicked").
		Joins("left join detail_product_marketplaces dpm on p.id = dpm.product_id").
		Where("merchant_id = ?", userID).
		Group("p.id").
		Scan(&merchantProducts).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range merchantProducts {
		merchantProducts[i].Images = strings.Split(v.Images, ",")[0]
	}

	currentMonth := currentTime.Day()
	lastMonth := currentTime.AddDate(0, -1, 0).Day()

	if err = db.Raw(`SELECT today, this_month, last_month
						FROM (
							SELECT COUNT(id) AS today
							FROM detail_product_marketplace_clicked
							WHERE merchant_id = ? AND created_at BETWEEN ? AND ?
						) AS subquery_today,
						(
							SELECT COUNT(id) AS this_month
							FROM detail_product_marketplace_clicked
							WHERE merchant_id = ? AND created_at BETWEEN ? AND ?
						) AS subquery_this_month,
						(
							SELECT COUNT(id) AS last_month
							FROM detail_product_marketplace_clicked
							WHERE merchant_id = ? AND created_at BETWEEN ? AND ?
						) AS subquery_last_month`,
		userID,
		currentTime.Format(constant.DateOnly)+" 00:00:00",
		currentTime.Format(constant.DateOnly)+" 23:59:59",
		userID,
		currentTime.Format(constant.YearMonth)+"-"+fmt.Sprintf("%d", currentMonth)+" 00:00:00",
		currentTime.Format(constant.YearMonth)+"-"+fmt.Sprintf("%d", currentMonth)+" 23:59:59",
		userID,
		currentTime.AddDate(0, -1, 0).Format(constant.YearMonth)+"-"+fmt.Sprintf("%d", lastMonth)+" 00:00:00",
		currentTime.AddDate(0, -1, 0).Format(constant.YearMonth)+"-"+fmt.Sprintf("%d", lastMonth)+" 23:59:59",
	).
		Scan(&visitors).Error; err != nil {
		return
	}

	res.MerchantProduct = merchantProducts
	res.MerchantVisitor = visitors

	return
}

func GetMerchantHomeProductByID(productID uint64) (res datastruct.MerchantHomeDetail, statusCode int, err error) {
	statusCode = http.StatusOK
	var (
		db               = Database()
		merchantProducts datastruct.MerchantProductByID
		marketplaces     []datastruct.MerchantMarketplace
	)

	if err = db.Table("products p").
		Select("p.id, images, name, price, status, description").
		Where("p.id = ?", productID).
		Scan(&merchantProducts).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}
	merchantProducts.Images = strings.Split(merchantProducts.Image, ",")

	if err = db.Table("detail_product_marketplaces dpm").
		Select("IFNULL(m.name, 'Offline') as name, clicked").
		Joins("left join marketplaces m on dpm.marketplace_id = m.id").
		Where("product_id = ?", productID).
		Scan(&marketplaces).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}
	res.MerchantProductByID = merchantProducts
	res.Marketplace = marketplaces
	return
}
