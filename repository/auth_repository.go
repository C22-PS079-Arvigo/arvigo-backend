package repository

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
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
	if err = db.Where(&datastruct.User{Email: loginData.Email}).First(&user).Error; err != nil {
		return tokenResponse, http.StatusNotFound, err
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

	tokenString, err := generateToken(user)
	if err != nil {
		return
	}

	tokenResponse.Token = tokenString
	return
}

func Register(userData datastruct.User) (tokenResponse datastruct.LoginRegisterResponse, statusCode int, err error) {
	db := Database()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return tokenResponse, http.StatusBadRequest, err
	}

	userPayload := datastruct.User{
		FullName: userData.FullName,
		Email:    userData.Email,
		Password: string(hashedPassword),
		// Role:      userData.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = db.Create(&userPayload).Error; err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenString, err := generateToken(userPayload)
	if err != nil {
		return tokenResponse, http.StatusInternalServerError, err
	}

	tokenResponse.Token = tokenString

	return tokenResponse, http.StatusCreated, nil
}

func generateToken(user datastruct.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(24 * 365 * time.Hour)
	claims := &Claims{
		ID:       user.ID,
		FullName: user.FullName,
		// Role:      user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(viper.GetString("jwt_secret")))
	if err != nil {
		return
	}

	return
}
