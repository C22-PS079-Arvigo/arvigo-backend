package repository

import (
	"fmt"
	"net/http"
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

// func GetMerchantAppHome(userID uint64) (res datastruct.HomeResponse, statusCode int, err error) {
// 	db := Database()
// 	statusCode = http.StatusOK

// 	var (
// 		user                  datastruct.UserWithPersonalityTag
// 		faceShapeProduct      []datastruct.HomeProduct
// 		personalityProduct    []datastruct.HomeProduct
// 		recommendationProduct []datastruct.HomeProduct
// 	)

// 	if err := db.Table("users").
// 		Select([]string{
// 			"users.*",
// 			"up.tag_ids",
// 		}).
// 		Where("users.id = ?", userID).
// 		Joins("LEFT JOIN user_personalities up on users.personality_id = up.id").
// 		First(&user).
// 		Error; err != nil {
// 		return res, http.StatusNotFound, errors.New("user not found")
// 	}

// 	if user.TagID != "" {
// 		tags := strings.Split(user.TagID, ",")
// 		for _, v := range tags {
// 			user.TagIDs = append(user.TagIDs, utils.StrToUint64(v, 0))
// 		}
// 	}

// 	// faceshape
// 	if user.IsCompleteFaceTest {
// 		if err := db.Table("products p").
// 			Select([]string{
// 				"p.id",
// 				"p.name",
// 				"p.images",
// 				"b.name as brand",
// 			}).
// 			Joins("LEFT JOIN brands b on b.id = p.brand_id").
// 			Joins("LEFT JOIN detail_product_tags dpt on p.id = dpt.product_id").
// 			Where("p.merchant_id = 0 AND p.category_id = ? AND dpt.tag_id IN (?)", constant.GlassesCategoryID, constant.GetFaceShapeTags[user.FaceShapeID]).
// 			Find(&faceShapeProduct).
// 			Error; err != nil {
// 			return res, http.StatusInternalServerError, err
// 		}

// 		for i, v := range faceShapeProduct {
// 			faceShapeProduct[i].Image = strings.Split(v.Image, ",")[0]
// 			var tagIDs []uint64
// 			var tags []string
// 			if err := db.Table("detail_product_tags").
// 				Select([]string{
// 					"tag_id",
// 				}).
// 				Where("product_id = ?", v.ID).
// 				Find(&tagIDs).
// 				Error; err != nil {
// 				return res, http.StatusInternalServerError, err
// 			}

// 			for _, v := range tagIDs {
// 				tags = append(tags, constant.GetTagNameByDetailTag[v]...)
// 			}
// 			faceShapeProduct[i].Tags = utils.RemoveDuplicates(tags)
// 		}
// 	}

// 	if user.IsCompletePersonalityTest {
// 		if err := db.Table("products p").
// 			Select([]string{
// 				"p.id",
// 				"p.name",
// 				"p.images",
// 				"b.name as brand",
// 			}).
// 			Joins("LEFT JOIN brands b on b.id = p.brand_id").
// 			Joins("LEFT JOIN detail_product_tags dpt on p.id = dpt.product_id").
// 			Where("p.merchant_id = 0 AND p.category_id = ? AND dpt.tag_id IN (?)", constant.MakeupCategoryID, user.TagIDs).
// 			Find(&personalityProduct).
// 			Error; err != nil {
// 			return res, http.StatusInternalServerError, err
// 		}

// 		for i, v := range personalityProduct {
// 			personalityProduct[i].Image = strings.Split(v.Image, ",")[0]
// 			var tagIDs []uint64
// 			var tags []string
// 			if err := db.Table("detail_product_tags").
// 				Select([]string{
// 					"tag_id",
// 				}).
// 				Where("product_id = ?", v.ID).
// 				Find(&tagIDs).
// 				Error; err != nil {
// 				return res, http.StatusInternalServerError, err
// 			}

// 			for _, v := range tagIDs {
// 				tags = append(tags, constant.GetTagNameByDetailTag[v]...)
// 			}
// 			personalityProduct[i].Tags = utils.RemoveDuplicates(tags)
// 		}
// 	}

// 	// recommendation TODO: integrate with ML/sort subs
// 	if err := db.Table("products p").
// 		Select([]string{
// 			"p.id",
// 			"p.name",
// 			"p.images",
// 			"b.name as brand",
// 		}).
// 		Joins("LEFT JOIN brands b on b.id = p.brand_id").
// 		Where("p.merchant_id = 0").
// 		Find(&recommendationProduct).
// 		Error; err != nil {
// 		return res, http.StatusInternalServerError, err
// 	}

// 	for i, v := range recommendationProduct {
// 		recommendationProduct[i].Image = strings.Split(v.Image, ",")[0]
// 		var tagIDs []uint64
// 		var tags []string
// 		if err := db.Table("detail_product_tags").
// 			Select([]string{
// 				"tag_id",
// 			}).
// 			Where("product_id = ?", v.ID).
// 			Find(&tagIDs).
// 			Error; err != nil {
// 			return res, http.StatusInternalServerError, err
// 		}

// 		for _, v := range tagIDs {
// 			tags = append(tags, constant.GetTagNameByDetailTag[v]...)
// 		}
// 		recommendationProduct[i].Tags = utils.RemoveDuplicates(tags)
// 	}

// 	res = datastruct.HomeResponse{
// 		Personality:    personalityProduct,
// 		FaceShape:      faceShapeProduct,
// 		Recommendation: recommendationProduct,
// 	}
// 	return
// }
