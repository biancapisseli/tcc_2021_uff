package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"strings"
)

const encodedAlg = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

func getSignature(secret, payload string) (signature string) {
	h := hmac.New(sha256.New, []byte(secret))
	// Error is ignore here because is a placeholder.
	// Analyzing sha256 writer, you can notice that error is never returned.
	h.Write(
		[]byte(strings.Join([]string{encodedAlg, payload}, ".")),
	)
	return b64.StdEncoding.EncodeToString(
		[]byte(hex.EncodeToString(h.Sum(nil))),
	)
}
