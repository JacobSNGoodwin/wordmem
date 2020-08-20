package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

// ImageRepository is data/repository implementation
// of service layer IUserRepository
type ImageRepository struct {
	Storage *storage.Client
}

const bucketName = "wordmem_profile_images"

// UploadUserImage uploads a user iamge file to provided storage client
func (r *ImageRepository) UploadUserImage(uid string, imageFile multipart.File) error {
	log.Printf("Uploading user image for uid: %v\n", uid)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	wc := r.Storage.Bucket(bucketName).Object(uid).NewWriter(ctx)

	// multipart.File has a writer!
	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("Unable to write file to Google Cloud Storage: %v\n", err)
		return errors.NewUnknown(http.StatusInternalServerError)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
