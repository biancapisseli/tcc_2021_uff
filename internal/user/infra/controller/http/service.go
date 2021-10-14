package userhttpctl

type UserHTTPController struct {
	repo            userdom.UserRepository
	passwordEncoder userdom.PasswordEncoder
}

func New(repo userdom.UserRepository, passwordEncoder userdom.PasswordEncoder) *UserHTTPController {
	return &UserHTTPController{
		repo:            repo,
		passwordEncoder: passwordEncoder,
	}
}
