package usersvc

import (
	"context"
	"errors"
	"fmt"

	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s *UserService) ChangePassword(
	ctx context.Context,
	userID uservo.UserID,
	currentPassword uservo.PasswordRaw,
	newPassword uservo.PasswordRaw,
	newPasswordConfirm uservo.PasswordRaw,
) (err error) {

	if !newPassword.Equals(newPasswordConfirm) {
		return errors.New("as senhas não coincidem")
	}

	user, err := s.repo.GetUserInfo(ctx, userID)
	if err != nil {
		return err
	}

	encodedCurrentPassword, err := s.passwordEncoder.EncodePassword(currentPassword)
	if err != nil {
		return err
	}

	if _, err := s.repo.GetUserByEmailAndPassword(ctx, user.Email, encodedCurrentPassword); err != nil {
		return err
	}

	encodedNewPassword, err := s.passwordEncoder.EncodePassword(newPassword)
	if err != nil {
		return fmt.Errorf("erro interno do servidor: %w", err)
	}

	if err := s.repo.UpdatePassword(ctx, userID, encodedNewPassword); err != nil {
		return fmt.Errorf("erro ao atualizar senha: %w", err)
	}

	return nil
}
