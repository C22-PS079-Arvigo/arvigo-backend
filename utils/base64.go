package utils

import (
	"encoding/base64"
	"io/ioutil"
)

func ImageToBase64(imagePath string) (string, error) {
	// Read the image file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	// Encode the image data to base64
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	return base64Data, nil
}
