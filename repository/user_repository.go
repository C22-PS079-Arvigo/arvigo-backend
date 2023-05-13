package repository

import (
	"errors"
	"net/http"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetUserByID(id uint64) (res datastruct.UserDetailResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var userDetail datastruct.UserDetail
	if err := db.Table("users").
		Select([]string{
			"users.*",
			"roles.name as role_name",
		}).
		Where("users.id = ?", id).
		Joins("JOIN roles ON roles.id = users.role_id").
		First(&userDetail).
		Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	var addressDetail datastruct.UserAddress
	if err := db.Table("addresses").
		Select([]string{
			"street",
			"prov_name as province",
			"city_name as city",
			"dis_name as district",
			"subdis_name as sub_district",
			"postal_code",
		}).
		Where("addresses.id", userDetail.ID).
		Joins("LEFT JOIN provinces ON prov_id = province_id").
		Joins("LEFT JOIN cities ON cities.city_id = addresses.city_id").
		Joins("LEFT JOIN districts ON dis_id = district_id").
		Joins("LEFT JOIN subdistricts ON subdis_id = subdistrict_id").
		Joins("LEFT JOIN postal_codes ON postal_id = postal_code_id").
		Find(&addressDetail).
		Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	res = datastruct.UserDetailResponse{
		Address:    addressDetail,
		UserDetail: userDetail,
	}
	return
}

func GetUsers() (res []datastruct.UserDetail, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Table("users").
		Select([]string{
			"users.*",
			"roles.name as role_name",
		}).
		Where("users.role_id = ?", constant.MobileApp).
		Joins("JOIN roles ON roles.id = users.role_id").
		First(&res).
		Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	return
}

func GetPartners() (res []datastruct.User, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Table("users").
		Select([]string{
			"users.*",
			"roles.name as role_name",
		}).
		Where("users.role_id = ?", constant.PartnerApp).
		Joins("JOIN roles ON roles.id = users.role_id").
		First(&res).
		Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}
	return
}
