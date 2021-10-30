package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

const encodedAlg = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

func getSignature(secret, payload string) (signature string, err error) {
	h := hmac.New(sha256.New, []byte(secret))
	if _, err = h.Write(
		[]byte(strings.Join([]string{encodedAlg, payload}, ".")),
	); err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}
	return b64.StdEncoding.EncodeToString(
		[]byte(hex.EncodeToString(h.Sum(nil))),
	), nil
}
