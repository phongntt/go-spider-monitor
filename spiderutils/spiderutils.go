package spiderutils

import (
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
	ciphertext, err := crypto_helper.Encrypt_Base64(text, GetMasterKey())
	if err != nil {
		return "", err
	}

	return ciphertext, nil
}

func SpiderDecrypt(text64 string) (string, error) {
	ciphertext, err := crypto_helper.Decrypt_Base64(text64, GetMasterKey())
	if err != nil {
		return "", err
	}

	return ciphertext, nil
}
