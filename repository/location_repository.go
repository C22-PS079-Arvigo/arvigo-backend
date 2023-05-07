package repository

import (
	"net/http"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetProvinces() (res []datastruct.Province, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetCities(provinceID uint64) (res []datastruct.City, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Where("prov_id", provinceID).Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetDistricts(cityID uint64) (res []datastruct.District, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Where("city_id", cityID).Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetSubDistricts(districtID uint64) (res []datastruct.SubDistrict, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Where("dis_id", districtID).Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}

func GetPostalCodes(subDistrictID uint64) (res []datastruct.PostalCode, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Where("subdis_id", subDistrictID).Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	return
}
