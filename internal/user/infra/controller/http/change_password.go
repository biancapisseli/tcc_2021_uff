package userhttpctl

import (
	"context"
	"errors"
	"fmt"
)

func (s *UserService) ChangePassword(
	ctx context.Context,
	userID userdom.UserID,
	newPassword userdom.PasswordRaw,
	newPasswordConfirm userdom.PasswordRaw,
) (err error) {
	if !newPassword.Equals(newPasswordConfirm) {
		return errors.New("as senhas não coincidem")
	}

	encodedPassword, err := s.passwordEncoder.EncodePassword(newPassword)
	if err != nil {
		return fmt.Errorf("erro interno do servidor: %w", err)
	}

	if err := s.repo.UpdatePassword(ctx, userID, encodedPassword); err != nil {
		return fmt.Errorf("erro ao atualizar senha: %w", err)
	}

	return nil
}
