package repository

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/pkg/database"
	"gorm.io/gorm"
)

func Database() (db *gorm.DB) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting database")
		return
	}
	return
}

func GenerateToken(user datastruct.User) (tokenString string, err error) {
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
