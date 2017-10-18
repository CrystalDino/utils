package utils

import (
	"crypto/sha1"
	"fmt"
)

func MakeSessionID(data string) string {
	h := sha1.New()
	return fmt.Sprintf("%x", h.Sum([]byte(data)))
}
