package userhttpcontroller

import (
	"context"
	uservo "ifoodish-store/services/user/domain/valueobject"
)

type Request interface {
	GetUserID() (userID uservo.UserID, err error)
	GetHeader(header string) (value string)
	ParseURLParams(dest interface{}) (err error)
	ParseBody(dest interface{}) (err error)
	Context() context.Context
}
