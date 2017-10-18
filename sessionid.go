package utils

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func MakeSessionID(data string) string {
	h := sha1.New()
	rd, tm := rand.Int63(), time.Now().UnixNano()
	salt := strconv.FormatInt(rd+tm, 10)
	h.Write([]byte(data))
	h.Write([]byte(salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}
