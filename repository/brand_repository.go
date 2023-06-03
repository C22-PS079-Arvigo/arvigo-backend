package repository

import (
	"net/http"
	"strings"
	"time"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetBrands() (res []datastruct.BrandResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Table("brands").Select("id, name, image").Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetBrandByCategory(categoryID uint64) (res []datastruct.BrandResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Table("brands").Select("id, name, image").Where("category_id", categoryID).Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func CreateBrand(data datastruct.BrandInput) (statusCode int, err error) {
	statusCode = http.StatusCreated

	var (
		db          = Database()
		currentTime = time.Now()
	)

	url, err := UploadImageToGCS(data.Image)
	if err != nil {
		return
	}

	brandPayload := datastruct.Brand{
		Name:       data.Name,
		CategoryID: data.CategoryID,
		Image:      url,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	if err = db.Create(&brandPayload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return
}

func UpdateBrand(brandID uint64, data datastruct.BrandInput) (statusCode int, err error) {
	statusCode = http.StatusCreated

	var (
		db          = Database()
		currentTime = time.Now()
	)

	url, err := UploadImageToGCS(data.Image)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	brandPayload := datastruct.Brand{
		Name:       data.Name,
		CategoryID: data.CategoryID,
		Image:      url,
		UpdatedAt:  currentTime,
	}

	if err = db.Where("id", brandID).Updates(&brandPayload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return
}

func GetListProductByBrand(brandID uint64) (res []datastruct.HomeProduct, statusCode int, err error) {
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
		Where("p.merchant_id = 0 AND b.id = ?", brandID).
		Find(&res).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for i, v := range res {
		res[i].Image = strings.Split(v.Image, ",")[0]
	}

	return
}
