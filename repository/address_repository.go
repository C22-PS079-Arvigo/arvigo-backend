package repository

import (
	"fmt"
	"net/http"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetAddressByID(addressID uint64) (res string, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var (
		address     datastruct.Address
		province    datastruct.Province
		city        datastruct.City
		district    datastruct.District
		subDistrict datastruct.SubDistrict
		postalCode  datastruct.PostalCode
	)

	if err = db.Where("id", addressID).First(&address).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err = db.Where("prov_id", address.ProvinceID).First(&province).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err = db.Where("city_id", address.CityID).First(&city).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err = db.Where("dis_id", address.DistrictID).First(&district).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err = db.Where("subdis_id", address.SubdistrictID).First(&subDistrict).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err = db.Where("postal_id", address.PostalCodeID).First(&postalCode).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = fmt.Sprintf("%s, %s, %s, %s, %s, %d", address.Street, subDistrict.Name, district.Name, city.Name, province.Name, postalCode.PostalCodeNumber)
	return
}
