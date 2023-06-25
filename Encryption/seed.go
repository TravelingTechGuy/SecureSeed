package Encryption

import (
	"crypto/sha256"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

func GetEntropy(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

func GetMnemonic(entropy []byte) ([]string, error) {
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	return strings.Split(mnemonic, " "), nil
}
