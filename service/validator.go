package service

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strings"
	"time"
)

var phoneRe = regexp.MustCompile(`^[0-9]{9,10}$`)

func parseBirthDay(s string) (time.Time, error) {
	if t, err := time.Parse("02/01/2006", s); err == nil {
		return t, nil
	}
	return time.Parse("2006-01-02", s)
}

func validateBase64(b64 string) error {
	if idx := strings.IndexByte(b64, ','); idx >= 0 {
		b64 = b64[idx+1:]
	}
	_, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return errors.New("invalid base64")
	}
	return nil
}
