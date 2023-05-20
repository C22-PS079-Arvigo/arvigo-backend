package repository

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/pkg/database"
	"github.com/yusufwib/arvigo-backend/pkg/storage"
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
		RoleID:   user.RoleID,
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

func UploadImageToGCS(fileHeader *multipart.FileHeader) (publicURL string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return publicURL, errors.New("failed to open uploaded image")
	}
	defer file.Close()

	// Create a temporary file to save the uploaded image
	tempDir := "./public/temp"
	tempFile, err := ioutil.TempFile(tempDir, "uploaded_image_*.jpg")
	if err != nil {
		return publicURL, errors.New("failed to create temporary file")
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}()

	// Save the uploaded image to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return publicURL, errors.New("failed to save uploaded image")
	}

	// Get the path of the uploaded image
	imagePath := tempFile.Name()
	objectName := fileHeader.Filename
	publicURL, err = storage.UploadImageToGCS(objectName, imagePath, os.Getenv("STORAGE_BUCKET_IMAGE_FOLDER"))
	if err != nil {
		return publicURL, fmt.Errorf("failed to upload image: %v", err)
	}

	return
}
