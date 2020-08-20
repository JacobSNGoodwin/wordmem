package service

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/model"
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
	// In this case, we have a one-to-one correspondence between service method "SignUp" and repository method "Create"
	// This is not always the case, though I can understand why this looks redundant
	return s.UserRepository.Create(&model.User{
		Email:    email,
		Password: password,
	})
}

// SignIn returns a user after comparing supplied email/password with
// stored email/password
func (s *UserService) SignIn(email string, password string) (*model.User, error) {
	u, err := s.UserRepository.FindByEmail(email)

	// Will return NotAuthorized to user doesn't know if no user was found
	if err != nil {
		return nil, errors.NewUnauthorized("Invalid email/password combination")
	}

	// verify password
	match, err := util.ComparePasswords(u.Password, password)

	if err != nil {
		return nil, errors.NewUnknown(http.StatusInternalServerError)
	}

	if !match {
		return nil, errors.NewUnauthorized("Invalid email and password combination")
	}

	return u, nil
}

// Get retreives a user based on their uuid
func (s *UserService) Get(uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(uid)

	return u, err
}

// UpdateOptions contains payload of updateable user account values
type UpdateOptions struct {
	Name      string
	Email     string
	Password  string
	ImageFile *multipart.FileHeader
	Website   string
}

// Update updates a user
func (s *UserService) Update(uid uuid.UUID, options *UpdateOptions) (*model.User, error) {
	// Open user file
	imageFile, err := options.ImageFile.Open()

	if err != nil {
		log.Printf("Failed to open image file: %v\n", err)
		return nil, errors.NewUnknown(500)
	}

	// Upload user's image to ImageRepository
	s.ImageRepositroy.UploadUserImage(uid.String(), &imageFile)

	return &model.User{}, nil
	// Update user in UserRepository
}
