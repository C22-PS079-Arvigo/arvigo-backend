package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/pkg/storage"
	"github.com/yusufwib/arvigo-backend/utils"
)

func FaceShapeRecognition(form *multipart.Form, userID uint64) (res datastruct.FaceShapeResponse, statusCode int, err error) {
	db := Database()
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

	var (
		isHumanRes  datastruct.IsHumanRes
		faceTestRes datastruct.FaceTestRes
	)
	response, err := utils.FetchMachineLearningAPI("POST", "/is_human", datastruct.FaceShapeMachineLearningPayload{
		Image: encodedImg,
	})

	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	err = json.Unmarshal(response, &isHumanRes)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	if !isHumanRes.Result {
		statusCode = http.StatusBadRequest
		err = errors.New("is not human")
		return
	}
	responseFaceShape, err := utils.FetchMachineLearningAPI("POST", "/face_shape", datastruct.FaceShapeMachineLearningPayload{
		Image: encodedImg,
	})

	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	err = json.Unmarshal(responseFaceShape, &faceTestRes)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	faceShapeID := constant.GetIDByShape[faceTestRes.Shape]
	if err = db.Model(&datastruct.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"face_shape_id":         faceShapeID,
			"is_complete_face_test": 1,
			"updated_at":            time.Now(),
		}).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	res.Result = faceTestRes.Shape
	return
}
