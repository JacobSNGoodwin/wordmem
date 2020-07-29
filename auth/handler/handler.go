package handler

import "github.com/jacobsngoodwin/wordmem/auth/repository"

// Env is a struct used for injected the repository into the various route handler
type Env struct {
	repository *repository.Repository
}
