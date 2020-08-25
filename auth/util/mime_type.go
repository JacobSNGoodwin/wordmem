package util

import "mime/multipart"

var validImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

// IsAllowedImageType determines if image is among types defined
// in map of allowed images
func IsAllowedImageType(imageFileHeader *multipart.FileHeader) bool {
	mimeType := imageFileHeader.Header.Get("Content-Type")

	_, exists := validImageTypes[mimeType]

	return exists
}
