package main

import (
	"SecureSeed/encryption"
	"SecureSeed/randomOrg"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func diceToStr(data []uint8) string {
	var ret string
	for _, c := range data {
		ret += strconv.Itoa(int(c))
	}
	return ret
}

func printMnemonic(mnemonic string) {
	fmt.Println("Mnemonic:")
	for i, w := range strings.Split(mnemonic, " ") {
		fmt.Printf("%2d: %-8s\t", i+1, w)
		if (i+1)%4 == 0 {
			fmt.Print("\n")
		}
	}
}

func printAddresses(addresses []string) {
	fmt.Println("Ethereum Addresses:")
	for i, w := range addresses {
		fmt.Printf("%2d: %s\n", i+1, w)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error reading env file: %s", err)
		os.Exit(1)
	}
	//get 100 dice rolls, either provided by user, or from Random.org
	var entropy string
	var ethAddress, bitcoinAddress uint
	var legacy bool

	flag.StringVar(&entropy, "e", "", "Provide 100 dice throws as a string")
	flag.UintVar(&ethAddress, "eth", 0, "Derive n Ethereum addresses")
	flag.UintVar(&bitcoinAddress, "btc", 0, "Derive n Bitcoin addresses")
	flag.BoolVar(&legacy, "legacy", false, "Derive Legacy Bitcoin addresses (instead of SegWit)")
	flag.Parse()

	if entropy != "" {
		if len(entropy) < 100 {
			log.Fatal("Please provide at least 100 dice throws")
			os.Exit(2)
		}
	} else {
		dice, err := randomOrg.GetDiceRoll(100)
		if err != nil {
			log.Fatalf("Error calling Random.org: %s", err)
			os.Exit(3)
		}
		entropy = diceToStr(dice)
	}
	fmt.Printf("Dice throws:\n%s\n\n", entropy)

	//calculate privateKey
	privateKey := encryption.GetPrivateKeyFromEntropy(entropy)
	fmt.Printf("Private key:\n%x\n\n", privateKey)

	//get mnemonic from entropy
	mnemonic, err := encryption.GetMnemonic(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	printMnemonic(mnemonic)
	fmt.Println()

	if ethAddress > 0 {
		//get Ethereum addresses from mnemonic
		addresses, err := encryption.DeriveEthereumAddresses(mnemonic, ethAddress)
		if err != nil {
			log.Fatal(err)
		}
		printAddresses(addresses)
	}

	if bitcoinAddress > 0 {
		fmt.Printf("BTC addresses TBD. %d addresses (did you ask for legcay format? %t)\n", bitcoinAddress, legacy)
	}
}
