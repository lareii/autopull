package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

/*
 * Signature, checks if the request signature matches the
 * signature of the specified values.
 */

func Signature(secret, header, payload string) error {
	// Split the header to get the signature.
	// The header is in the format: sha256=xxxxx
	header = strings.Split(header, "=")[1]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))

	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expected), []byte(header)) {
		return errors.New("request signatures didn't match")
	}

	return nil
}
