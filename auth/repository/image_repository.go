package repository

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
)

// ImageRepository is data/repository implementation
// of service layer IUserRepository
type ImageRepository struct {
	Storage *storage.Client
}

const bucketName = "wordmem_profile_images"

// UploadUserImage uploads a user iamge file to provided storage client
func (r *ImageRepository) UploadUserImage(objName string, imageFile multipart.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // maybe don't hardcode this?
	defer cancel()

	bckt := r.Storage.Bucket(bucketName)

	object := bckt.Object(objName)
	wc := object.NewWriter(ctx)

	// set cache control so profile image will be served fresh by browsers
	// To do this with object handle, you'd first have to upload, then update
	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"

	// multipart.File has a writer!
	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("Unable to write file to Google Cloud Storage: %v\n", err)
		return rerrors.NewInternal()
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
