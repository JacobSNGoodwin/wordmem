package service

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository  IUserRepository
	ImageRepositroy IImageRepositroy
}

// SignUp creates a new user based on data in model.User
func (s *UserService) SignUp(email string, password string) (*model.User, error) {
	u := &model.User{
		Email:    email,
		Password: password,
	}
	if err := s.UserRepository.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

// SignIn returns a user after comparing supplied email/password with
// stored email/password
func (s *UserService) SignIn(email string, password string) (*model.User, error) {
	u, err := s.UserRepository.FindByEmail(email)

	// Will return NotAuthorized to user doesn't know if no user was found
	if err != nil {
		return nil, rerrors.NewAuthorization("Invalid email/password combination")
	}

	// verify password
	match, err := util.ComparePasswords(u.Password, password)

	if err != nil {
		return nil, rerrors.NewInternal()
	}

	if !match {
		return nil, rerrors.NewAuthorization("Invalid email and password combination")
	}

	return u, nil
}

// Get retreives a user based on their uuid
func (s *UserService) Get(uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(uid)

	return u, err
}

// UserDetails contains payload of updateable user account details
type UserDetails struct {
	Name    string
	Email   string
	Website string
}

// UpdateDetails updates a user's details
// Or everything save photo and password
func (s *UserService) UpdateDetails(u *model.User) error {
	// Update user in UserRepository
	err := s.UserRepository.Update(u)

	return err
}

// SetProfileImage reaches out to the image repository to upload an image to
// cloud storage. It also reaches out to the user repository to set the user's url
// If the provided image file is nil, the user repository will be called to update the image
// to nil
func (s *UserService) SetProfileImage(uid uuid.UUID, imageFileHeader *multipart.FileHeader) (string, error) {
	// get current users image url to get object name that will be replaced/deleted
	u, err := s.UserRepository.FindByID(uid)
	if err != nil {
		return "", err
	}

	// if the user has empty string for image,
	// create a new identifier for the image
	var objName string

	if u.StorageID == "" {
		objID, _ := uuid.NewRandom()
		objName = objID.String()
		fmt.Println("User creating fresh image")
	} else {
		objName = u.StorageID
	}

	if imageFileHeader == nil {
		return "", rerrors.NewBadRequest("Must include an imageFile")
	}

	// Validate image mime-type is allowable
	if valid := util.IsAllowedImageType(imageFileHeader); !valid {
		log.Println("Image is not an allowable mimtype")
		return "", rerrors.NewBadRequest("imageFile must be 'image/jpeg' or 'image/png'")
	}

	imageFile, err := imageFileHeader.Open()
	if err != nil {
		log.Printf("Failed to open image file: %v\n", err)
		return "", rerrors.NewInternal()
	}

	// Upload user's image to ImageRepository
	if err := s.ImageRepositroy.UploadUserImage(objName, imageFile); err != nil {
		log.Printf("Unable to upload image to cloud provider: %v\n", err)
		return "", err
	}

	// build url to store in DB (should store this path globally somewhere - maybe env)
	imageURL := fmt.Sprintf("https://storage.googleapis.com/wordmem_profile_images/%s", objName)

	if err := s.UserRepository.UpdateImage(uid, imageURL, objName); err != nil {
		log.Printf("Unable to update imageURL: %v\n", err)
		return "", err
	}

	return imageURL, nil

}

// ClearProfileImage reaches out to the user repository to
// clear the user's image url
func (s *UserService) ClearProfileImage(uid uuid.UUID) error {
	// get current users image url to get object name that will be replaced/deleted
	u, err := s.UserRepository.FindByID(uid)
	if err != nil {
		return err
	}

	// if the user has empty string for image,
	// create a new identifier for the image
	var objName string

	if u.StorageID == "" {
		objID, _ := uuid.NewRandom()
		objName = objID.String()
		fmt.Println("User creating fresh image")
	} else {
		objName = u.StorageID
	}
	// clear imageUrl in user repo
	// we'll keep the storageID and actual object
	// for future overwrites
	return s.UserRepository.UpdateImage(uid, "", objName)
}

// Delete is used to remove a user (for rollback purposes)
func (s *UserService) Delete(uuid uuid.UUID) error {
	// Update user in UserRepository
	err := s.UserRepository.Delete(uuid)

	return err
}
