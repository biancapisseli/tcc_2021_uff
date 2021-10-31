package useruc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	uservo "ifoodish-store/internal/user/domain/valueobject"

	"github.com/carlmjohnson/resperr"
)

func (s UserUseCases) ChangePassword(
	ctx context.Context,
	userID uservo.UserID,
	currentPassword uservo.PasswordRaw,
	newPassword uservo.PasswordRaw,
	newPasswordConfirm uservo.PasswordRaw,
) (err error) {

	if !newPassword.Equals(newPasswordConfirm) {
		return resperr.WithCodeAndMessage(
			errors.New("passwords doesn't match"),
			http.StatusBadRequest,
			"As senhas não coincidem",
		)
	}

	user, err := s.repo.GetUserInfo(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user's info from repo: %w", err)
	}

	encodedCurrentPassword, err := s.passwordEncoder.EncodePassword(currentPassword)
	if err != nil {
		return fmt.Errorf("error encoding user's current password: %w", err)
	}

	if _, err := s.repo.GetUserByEmailAndPassword(ctx, user.Email, encodedCurrentPassword); err != nil {
		return fmt.Errorf("error getting user's info by email and password from repo: %w", err)
	}

	encodedNewPassword, err := s.passwordEncoder.EncodePassword(newPassword)
	if err != nil {
		return fmt.Errorf("error encoding user's new password: %w", err)
	}

	if err := s.repo.UpdatePassword(ctx, userID, encodedNewPassword); err != nil {
		return fmt.Errorf("error updating user's password: %w", err)
	}

	return nil
}
