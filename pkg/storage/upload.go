package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
	"github.com/yusufwib/arvigo-backend/utils"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func UploadImageToGCS(filename, imagePath, folder string) (publicURL string, err error) {
	ctx := context.Background()

	bucketName := os.Getenv("STORAGE_BUCKET_NAME")
	if bucketName == "" {
		bucketName = "arvigo-bucket"
	}

	keyJSON, err := ioutil.ReadFile("./gcp-cred.json")
	if err != nil {
		return publicURL, fmt.Errorf("failed to read service account key JSON file: %v", err)
	}

	creds, err := google.CredentialsFromJSON(ctx, keyJSON, storage.ScopeReadWrite)
	if err != nil {
		return publicURL, fmt.Errorf("failed to create GCS client: %v", err)
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return publicURL, fmt.Errorf("failed to create GCS client: %v", err)
	}
	defer client.Close()

	file, err := os.Open(imagePath)
	if err != nil {
		return publicURL, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Create the object name with the desired folder and new filename
	objectName := fmt.Sprintf("%s/%s_%s", folder, utils.GenerateRandomStringWithTimestamp(10), filename)

	writer := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return publicURL, fmt.Errorf("failed to upload image to GCS: %v", err)
	}

	publicURL = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	return publicURL, nil
}
