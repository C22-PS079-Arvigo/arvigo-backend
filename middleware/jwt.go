package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/pkg/cache"
	"github.com/yusufwib/arvigo-backend/utils"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		// Parse JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid token signing method")
			}
			// Get secret key from config
			secretKey := viper.GetString("jwt.secret")
			return []byte(secretKey), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Check if token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := fmt.Sprintf("%v", claims["id"])
			roleID := fmt.Sprintf("%v", claims["role_id"])
			fullName := fmt.Sprintf("%v", claims["full_name"])

			// Set data to struct
			UserAuthData := datastruct.UserAuth{
				ID:       utils.StrToUint64(userID, 0),
				FullName: fullName,
				RoleID:   utils.StrToUint64(roleID, 0),
			}

			// Set to Context
			c.Set("userAuth", &UserAuthData)

			// Store data in Redis
			err := storeUserAuthInRedis(userID, &UserAuthData)
			if err != nil {
				log.Println("Failed to store userAuth in Redis:", err)
			}

			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
	}
}

func storeUserAuthInRedis(userID string, userAuth *datastruct.UserAuth) error {
	// Get the Redis client from the global variable or initialize it if not already done
	redisClient, err := cache.ConnectRedis()
	if err != nil {
		return err
	}

	// Convert userAuth data to JSON string
	userAuthJSON, err := json.Marshal(userAuth)
	if err != nil {
		return err
	}

	// Store the userAuth data in Redis with a unique key
	err = redisClient.Set(context.Background(), "userAuth:"+userID, userAuthJSON, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
