package md5passwordencoder

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

type MD5PasswordEncoder struct {
}

func New() *MD5PasswordEncoder {
	return &MD5PasswordEncoder{}
}

func (e MD5PasswordEncoder) EncodePassword(
	rawPassword uservo.PasswordRaw,
) (encodedPassword uservo.PasswordEncoded, err error) {
	hash := md5.Sum([]byte(rawPassword.String()))
	encoded := hex.EncodeToString(hash[:])
	encodedPassword, err = uservo.NewPasswordEncoded(encoded)
	if err != nil {
		return encodedPassword, fmt.Errorf("error encoding password: %w", err)
	}
	return encodedPassword, nil
}
