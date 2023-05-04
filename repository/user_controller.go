package repository

import (
	"errors"
	"net/http"

	"github.com/yusufwib/arvigo-backend/datastruct"
)

func GetUserByID(id uint64) (user datastruct.User, statusCode int, err error) {
	db := Database()

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{
	// 		"message": "Invalid user ID",
	// 	})
	// }

	if err := db.Where(&datastruct.User{ID: id}).First(&user).Error; err != nil {
		return user, http.StatusNotFound, errors.New("user not found")
	}

	return user, http.StatusOK, nil
}

// func GetAllUsers(c echo.Context) error {
// 	db := database.DBConn

// 	var users []datastruct.User
// 	if err := db.Find(&users).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"message": "Failed to retrieve users",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, users)
// }

// func UpdateUser(c echo.Context) error {
// 	db := database.DBConn

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"message": "Invalid user ID",
// 		})
// 	}

// 	var user datastruct.User
// 	if err := db.Where(&datastruct.User{ID: id}).First(&user).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]string{
// 			"message": "User not found",
// 		})
// 	}

// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"message": "Invalid user data",
// 		})
// 	}

// 	if err := db.Save(&user).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"message": "Failed to update user",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func DeleteUser(c echo.Context) error {
// 	db := database.DBConn

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"message": "Invalid user ID",
// 		})
// 	}

// 	var user datastruct.User
// 	if err := db.Where(&datastruct.User{ID: id}).First(&user).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]string{
// 			"message": "User not found",
// 		})
// 	}

// 	if err := db.Delete(&user).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"message": "Failed to delete user",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "User deleted successfully",
// 	})
// }
