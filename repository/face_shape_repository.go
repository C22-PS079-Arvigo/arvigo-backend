package repository

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/pkg/storage"
	"github.com/yusufwib/arvigo-backend/utils"
)

func FaceShapeRecognition(form *multipart.Form) (res datastruct.FaceShapeResponse, statusCode int, err error) {
	fileHeaders := form.File["image"]
	if len(fileHeaders) == 0 {
		return res, http.StatusBadRequest, errors.New("failed to get image from form data")
	}

	fileHeader := fileHeaders[0]
	file, err := fileHeader.Open()
	if err != nil {
		return res, http.StatusInternalServerError, errors.New("failed to open uploaded image")
	}
	defer file.Close()

	// Create a temporary file to save the uploaded image
	tempDir := "./public/temp"
	tempFile, err := ioutil.TempFile(tempDir, "uploaded_image_*.jpg")
	if err != nil {
		return res, http.StatusInternalServerError, errors.New("failed to create temporary file")
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}()

	// Save the uploaded image to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return res, http.StatusInternalServerError, errors.New("failed to save uploaded image")
	}

	// Get the path of the uploaded image
	imagePath := tempFile.Name()
	encodedImg, err := utils.ImageToBase64(imagePath)
	if err != nil {
		return res, http.StatusInternalServerError, fmt.Errorf("failed to converting img to base64")
	}

	objectName := fileHeader.Filename
	publicURL, err := storage.UploadImageToGCS(objectName, imagePath, os.Getenv("STORAGE_BUCKET_IMAGE_FOLDER"))
	if err != nil {
		return res, http.StatusInternalServerError, fmt.Errorf("failed to upload image: %v", err)
	}

	res.ImageUrl = publicURL
	res.Result = encodedImg
	return
}
