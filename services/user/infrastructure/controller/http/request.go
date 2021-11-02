package userhttpcontroller

import (
	"context"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

type Request interface {
	GetURLParam(name string) (value string)
	GetUserID() (userID uservo.UserID, err error)
	GetHeader(header string) (value string)
	ParseBody(dest interface{}) (err error)
	Context() context.Context
}
