package handler

// MaxBodySize is the maximum allowed body size for any request
const MaxBodySize int64 = 4 * 1024 * 1024

// Env is a struct used for injected the repository into the various route handler
type Env struct {
	UserService  IUserService
	TokenService ITokenService
}
