package repository

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/utils"
)

func GetHome(userID uint64) (res datastruct.HomeResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var (
		user                  datastruct.UserWithPersonalityTag
		faceShapeProduct      []datastruct.HomeProduct
		personalityProduct    []datastruct.HomeProduct
		recommendationProduct []datastruct.HomeProduct
	)

	if err := db.Table("users").
		Select([]string{
			"users.*",
			"up.tag_ids",
		}).
		Where("users.id = ?", userID).
		Joins("LEFT JOIN user_personalities up on users.personality_id = up.id").
		First(&user).
		Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	if user.TagID != "" {
		tags := strings.Split(user.TagID, ",")
		for _, v := range tags {
			user.TagIDs = append(user.TagIDs, utils.StrToUint64(v, 0))
		}
	}

	// faceshape
	if user.IsCompleteFaceTest {
		if err := db.Table("products p").
			Select([]string{
				"p.id",
				"p.name",
				"p.images",
				"b.name as brand",
			}).
			Joins("LEFT JOIN brands b on b.id = p.brand_id").
			Joins("LEFT JOIN detail_product_tags dpt on p.id = dpt.product_id").
			Where("p.merchant_id = 0 AND p.category_id = ? AND dpt.tag_id IN (?)", constant.GlassesCategoryID, constant.GetFaceShapeTags[user.FaceShapeID]).
			Find(&faceShapeProduct).
			Error; err != nil {
			return res, http.StatusInternalServerError, err
		}

		for i, v := range faceShapeProduct {
			faceShapeProduct[i].Image = strings.Split(v.Image, ",")[0]
		}
	}

	// personality TODO: integrate with ML
	if user.IsCompletePersonalityTest {
		if err := db.Table("products p").
			Select([]string{
				"p.id",
				"p.name",
				"p.images",
				"b.name as brand",
			}).
			Joins("LEFT JOIN brands b on b.id = p.brand_id").
			Joins("LEFT JOIN detail_product_tags dpt on p.id = dpt.product_id").
			Where("p.merchant_id = 0 AND p.category_id = ? AND dpt.tag_id IN (?)", constant.MakeupCategoryID, user.TagIDs).
			Find(&personalityProduct).
			Error; err != nil {
			return res, http.StatusInternalServerError, err
		}

		for i, v := range personalityProduct {
			personalityProduct[i].Image = strings.Split(v.Image, ",")[0]
		}
	}

	// recommendation TODO: integrate with ML
	if err := db.Table("products p").
		Select([]string{
			"p.id",
			"p.name",
			"p.images",
			"b.name as brand",
		}).
		Joins("LEFT JOIN brands b on b.id = p.brand_id").
		Where("p.merchant_id = 0").
		Find(&recommendationProduct).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range recommendationProduct {
		recommendationProduct[i].Image = strings.Split(v.Image, ",")[0]
	}

	res = datastruct.HomeResponse{
		Personality:    personalityProduct,
		FaceShape:      faceShapeProduct,
		Recommendation: recommendationProduct,
	}
	return
}

func GetHomeSeach(search string) (res []datastruct.HomeProduct, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Table("products p").
		Select([]string{
			"p.id",
			"p.name",
			"p.images",
			"b.name as brand",
		}).
		Joins("LEFT JOIN brands b on b.id = p.brand_id").
		Where("p.merchant_id = 0 AND p.name LIKE ?", fmt.Sprintf("%%%s%%", search)).
		Find(&res).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range res {
		res[i].Image = strings.Split(v.Image, ",")[0]
	}

	return
}
