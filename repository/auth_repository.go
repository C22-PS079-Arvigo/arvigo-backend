package repository

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func Login(loginData datastruct.LoginUserInput) (tokenResponse datastruct.LoginRegisterResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var user datastruct.User
	if err = db.Where("email = ?", loginData.Email).
		Where("role_id = ?", constant.ConvertRoleID[loginData.Role]).
		First(&user).Error; err != nil {
		return tokenResponse, http.StatusNotFound, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Passwords do not match
			return tokenResponse, http.StatusUnauthorized, err
		}
		// Other bcrypt comparison error occurred
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenString, err := GenerateToken(user)
	if err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenResponse = datastruct.LoginRegisterResponse{
		UserID: user.ID,
		Token:  tokenString,
	}
	return
}

func RegisterUser(userData datastruct.UserRegisterInput) (tokenResponse datastruct.LoginRegisterResponse, statusCode int, err error) {
	db := Database()

	if strings.TrimSpace(userData.Password) != strings.TrimSpace(userData.PasswordConfirmation) {
		return tokenResponse, http.StatusBadRequest, errors.New("password is doesn't match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return tokenResponse, http.StatusBadRequest, err
	}

	userPayload := datastruct.User{
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  string(hashedPassword),
		RoleID:    constant.MobileApp,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = db.Create(&userPayload).Error; err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenString, err := GenerateToken(userPayload)
	if err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenResponse = datastruct.LoginRegisterResponse{
		UserID: userPayload.ID,
		Token:  tokenString,
	}
	return tokenResponse, http.StatusCreated, nil
}

func RegisterPartner(partnerData datastruct.PartnerRegisterInput) (tokenResponse datastruct.LoginRegisterResponse, statusCode int, err error) {
	db := Database()

	if strings.TrimSpace(partnerData.Password) != strings.TrimSpace(partnerData.PasswordConfirmation) {
		return tokenResponse, http.StatusBadRequest, errors.New("password is doesn't match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(partnerData.Password), bcrypt.MinCost)
	if err != nil {
		return tokenResponse, http.StatusBadRequest, err
	}

	merchantPayload := datastruct.Merchant{
		Name:             partnerData.StoreName,
		IsRecommendation: 0, // default
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err = db.Create(&merchantPayload).Error; err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	addressPayload := datastruct.Address{
		Street:        partnerData.Street,
		ProvinceID:    partnerData.ProvinceID,
		DistrictID:    partnerData.DistrictID,
		SubdistrictID: partnerData.SubdistrictID,
		PostalCodeID:  partnerData.PostalCodeID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err = db.Create(&addressPayload).Error; err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	userPayload := datastruct.User{
		FullName:   partnerData.StoreName,
		Email:      partnerData.Email,
		Password:   string(hashedPassword),
		RoleID:     constant.PartnerApp,
		MerchantID: merchantPayload.ID,
		AddressID:  addressPayload.ID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err = db.Create(&userPayload).Error; err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenString, err := GenerateToken(userPayload)
	if err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenResponse = datastruct.LoginRegisterResponse{
		UserID: userPayload.ID,
		Token:  tokenString,
	}
	return tokenResponse, http.StatusCreated, nil
}
