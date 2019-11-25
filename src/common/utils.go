package common

import (
	"crypto/rc4"
	"fmt"
	"os"
	"time"
)

func GetFileSize(file string) (int64, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	if fi.IsDir() {
		return 0, fmt.Errorf("target file %s is not file", file)
	}

	return fi.Size(), nil
}

func InStringArray(value string, arrays []string) bool {
	for _, v := range arrays {
		if v == value {
			return true
		}
	}

	return false
}

func Rc4(content []byte, key []byte) ([]byte, error) {
	rc4Cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(content))
	rc4Cipher.XORKeyStream(plainText, content)

	return plainText, nil
}

func HumanDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func IF(cond bool, f interface{}, s interface{}) interface{} {
	if cond {
		return f
	}

	return s
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}