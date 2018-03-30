package spiderutils

import (
	"encoding/base64"
	"github.com/phongntt/crypto_helper"
	"os"
)

const EVN_MASTER_KEY = "GO_SPIDER_MASTERKEY"
const DEFAULTKEY = "gospidergospidergospidergospider"

func GetMasterKey() string {
	if os.Getenv(EVN_MASTER_KEY) == "" {
		return DEFAULTKEY
	}
	return os.Getenv(EVN_MASTER_KEY)
}

func SpiderEncrypt(text string) (string, error) {
	byteText := []byte(text)
	byteKey := []byte(GetMasterKey())

	ciphertext, err := crypto_helper.Encrypt(byteText, byteKey)
	if err != nil {
		return "", err
	}

	encoded64 := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded64, nil
}

func SpiderDecrypt(text64 string) (string, error) {
	byteKey := []byte(GetMasterKey())
	byteText, err64 := base64.StdEncoding.DecodeString(text64)
	if err64 != nil {
		return "", err64
	}

	ciphertext, err := crypto_helper.Decrypt(byteText, byteKey)
	if err != nil {
		return "", err
	}

	return string(ciphertext), nil
}
