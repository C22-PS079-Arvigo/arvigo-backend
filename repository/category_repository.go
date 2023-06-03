package repository

import (
	"net/http"
	"strings"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetCategories() (res []datastruct.CategoryResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Table("categories").Select("id, name").Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetListProductByCategory(categoryID uint64) (res []datastruct.HomeProduct, statusCode int, err error) {
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
		Where("p.merchant_id = 0 AND p.category_id = ?", categoryID).
		Find(&res).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range res {
		res[i].Image = strings.Split(v.Image, ",")[0]
	}

	return
}
