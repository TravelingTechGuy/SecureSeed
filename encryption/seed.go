package encryption

import (
	"crypto/sha256"
	"fmt"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func GetPrivateKeyFromEntropy(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

func GetMnemonic(entropy []byte) (string, error) {
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

func DeriveEthereumAddresses(mnemonic string, num int) ([]string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	result := make([]string, num)

	for i := 0; i < num; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			return nil, err
		}
		result[i] = account.Address.Hex()
	}
	return result, nil
}
