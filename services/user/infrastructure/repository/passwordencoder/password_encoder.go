package md5passwordencoder

import (
	"crypto/md5"
	"encoding/hex"
	uservo "ifoodish-store/services/user/domain/valueobject"
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
	return uservo.NewPasswordEncoded(encoded), nil
}
