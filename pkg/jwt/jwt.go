package jwt

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"
	"net/http"
	"strings"
	"time"
)

type jwt struct {
	UserID   uservo.UserID `json:"sub"`
	IssuedAt int64         `json:"iat"`
}

type jwtContextKey struct{}

func ParseJWT(parentCtx context.Context, secret, rawToken string) (jwtCtx context.Context, err error) {

	components := strings.Split(rawToken, ".")
	if len(components) != 3 {
		return parentCtx, resperr.WithStatusCode(
			errors.New("jwt bad format"),
			http.StatusUnauthorized,
		)
	}

	if components[0] != encodedAlg {
		return parentCtx, resperr.WithStatusCode(
			errors.New("jwt algorithm doesn't match"),
			http.StatusUnauthorized,
		)
	}

	expectedSignature, err := getSignature(secret, components[1])
	if err != nil {
		return parentCtx, resperr.WithStatusCode(
			fmt.Errorf("fail creating JWT: %w", err),
			http.StatusInternalServerError,
		)
	}

	if components[2] != expectedSignature {
		return parentCtx, resperr.WithStatusCode(
			errors.New("signatures doesn't match"),
			http.StatusUnauthorized,
		)
	}

	decodedPayload, err := b64.StdEncoding.DecodeString(components[1])
	if err != nil {
		return parentCtx, resperr.WithStatusCode(
			errors.New("impossible to decode payload"),
			http.StatusUnauthorized,
		)
	}

	type userIDClone uservo.UserID
	jwtCloneObj := struct {
		UserID   userIDClone `json:"sub"`
		IssuedAt int64       `json:"iat"`
	}{}

	if err := json.Unmarshal(decodedPayload, &jwtCloneObj); err != nil {
		return parentCtx, resperr.WithStatusCode(
			errors.New("impossible to unmarshal JWT object"),
			http.StatusUnauthorized,
		)
	}

	if jwtCloneObj.IssuedAt > time.Now().Unix() {
		return parentCtx, resperr.WithStatusCode(
			errors.New("jwt issued in the future"),
			http.StatusUnauthorized,
		)
	}

	userID, err := uservo.NewUserID(uservo.UserID(jwtCloneObj.UserID).String())
	if err != nil {
		return parentCtx, fmt.Errorf("jwt invalid user id: %w", err)
	}

	return context.WithValue(parentCtx, jwtContextKey{}, &jwt{
		IssuedAt: jwtCloneObj.IssuedAt,
		UserID:   userID,
	}), nil
}

func GetUserID(ctx context.Context) (userID uservo.UserID, err error) {
	intfcValue := ctx.Value(jwtContextKey{})
	if err != nil {
		return userID, resperr.WithStatusCode(
			errors.New("jwt not found on context"),
			http.StatusUnauthorized,
		)
	}

	jwt, ok := intfcValue.(*jwt)
	if !ok {
		return userID, resperr.WithStatusCode(
			errors.New("wrong jwt type in context"),
			http.StatusUnauthorized,
		)
	}

	return jwt.UserID, nil
}

func CreateJWT(secret string, userID uservo.UserID) (token string, err error) {
	jwtObj := jwt{
		UserID:   userID,
		IssuedAt: time.Now().Unix(),
	}

	rawPayload, err := json.Marshal(jwtObj)
	if err != nil {
		return "", resperr.WithStatusCode(
			fmt.Errorf("fail creating JWT: %w", err),
			http.StatusInternalServerError,
		)
	}

	payload := b64.StdEncoding.EncodeToString(rawPayload)

	signature, err := getSignature(secret, payload)
	if err != nil {
		return "", resperr.WithStatusCode(
			fmt.Errorf("fail creating JWT: %w", err),
			http.StatusInternalServerError,
		)
	}

	return fmt.Sprintf(
		"%s.%s.%s",
		encodedAlg,
		payload,
		signature,
	), nil
}
