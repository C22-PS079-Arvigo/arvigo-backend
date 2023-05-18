package repository

import (
	"net/http"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetCategories() (res []datastruct.CategoryResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Table("categories").Select("name").Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}
