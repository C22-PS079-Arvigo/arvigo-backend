package repository

import (
	"errors"
	"net/http"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetUserByID(id uint64) (res datastruct.User, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Where(&datastruct.User{ID: id}).First(&res).Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	return
}

func GetUsers() (res []datastruct.User, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Where(&datastruct.User{RoleID: constant.MobileApp}).Find(&res).Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	return
}

func GetPartners() (res []datastruct.User, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err := db.Where(&datastruct.User{RoleID: constant.PartnerApp}).Find(&res).Error; err != nil {
		return res, http.StatusNotFound, errors.New("user not found")
	}

	return
}
