package main

import (
	"SecureSeed/encryption"
	"SecureSeed/randomOrg"
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

func printMnemonic(mnemonic []string) {
	for i, w := range mnemonic {
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
	args := os.Args[1:]
	var entropy string
	if len(args) > 0 && strings.ToLower(args[0]) == "-e" {
		entropy = args[1]
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
	fmt.Println("Mnemonic:")
	printMnemonic(mnemonic)

	// encryption.GetEthereumAddress(privateKey)
}
