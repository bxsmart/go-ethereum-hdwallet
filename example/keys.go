package main

import (
	"fmt"
	"github.com/bxsmart/go-hdwallet"
	"log"
)

func main() {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"

	wallet, err := go_hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := go_hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account address: %s\n", account.Address.Hex())

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key in hex: %s\n", privateKey)

	publicKey, _ := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Public key in hex: %s\n", publicKey)

}
