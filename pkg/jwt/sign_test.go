package jwt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSignature(t *testing.T) {
	require := require.New(t)

	// payload is base64 for `{"lala":"lala","lele":"lele"}`
	payload := "eyJsYWxhIjoibGFsYSIsImxlbGUiOiJsZWxlIn0"
	secret := "lalalala"

	// This value is retrieve from online HMAC SHA256 and base64 encoders
	signature := getSignature(secret, payload)
	require.Equal("MTE0YmY1Y2YzZmMyMGU1YWJhMTgxODQ3ZTMwOGEyMDQ2N2YwYTM5YmYzZjhlMzNjMTMxYjVhMDhiZTgxNWEyMA==", signature)
}
