package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GenerateId() string {
	return fmt.Sprintf("%v", time.Now().UnixNano())
}
