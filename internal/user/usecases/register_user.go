package useruc

import (
	"context"
	"errors"
	"fmt"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"net/http"

	"github.com/carlmjohnson/resperr"
)

func (s UserUseCases) RegisterUser(
	ctx context.Context,
	user userent.User,
	password uservo.PasswordRaw,
	passwordConfirm uservo.PasswordRaw,
) (userID uservo.UserID, err error) {
	if !password.Equals(passwordConfirm) {
		return userID, resperr.WithCodeAndMessage(
			errors.New("passwords doesn't match"),
			http.StatusBadRequest,
			"As senhas não coincidem",
		)
	}

	encodedPassword, err := s.passwordEncoder.EncodePassword(password)
	if err != nil {
		return userID, fmt.Errorf("error encoding user's current password: %w", err)
	}

	return s.repo.AddUser(ctx, user, encodedPassword)
}
