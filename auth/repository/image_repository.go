package repository

import (
	"mime/multipart"

	"cloud.google.com/go/storage"
)

// ImageRepository is data/repository implementation
// of service layer IUserRepository
type ImageRepository struct {
	Storage *storage.Client
}

// UploadUserImage uploads a user iamge file to provided storage client
func (r *ImageRepository) UploadUserImage(uid string, imageFile *multipart.File) (string, error) {
	panic("Not implemented")
}
