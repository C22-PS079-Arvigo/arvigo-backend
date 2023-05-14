package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func UploadImageToGCS(bucketName, objectName, imagePath string) error {
	// Create a context using your project's default authentication credentials
	ctx := context.Background()

	keyJSON, err := ioutil.ReadFile("gcp-cred.json")
	if err != nil {
		return fmt.Errorf("failed to read service account key JSON file: %v", err)
	}

	// Set up authentication using the service account key JSON file
	creds, err := google.CredentialsFromJSON(ctx, keyJSON, storage.ScopeReadWrite)
	if err != nil {
		return fmt.Errorf("failed to create GCS client: %v", err)
	}

	// Create a new GCS client using the authenticated credentials
	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return fmt.Errorf("failed to create GCS client: %v", err)
	}
	defer client.Close()

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Create a new GCS writer for the object
	writer := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	defer writer.Close()

	// Copy the image file contents to the GCS writer
	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("failed to upload image to GCS: %v", err)
	}

	log.Printf("Image uploaded successfully to gs://%s/%s", bucketName, objectName)
	return nil
}
