package main

import (
	"SecureSeed/Encryption"
	"SecureSeed/RandomOrg"
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
	var data string
	if len(args) > 0 && strings.ToLower(args[0]) == "-e" {
		data = args[1]
		if len(data) < 100 {
			log.Fatal("Please provide at least 100 dice throws")
			os.Exit(2)
		}
	} else {
		dice, err := RandomOrg.GetDiceRoll(100)
		if err != nil {
			log.Fatalf("Error calling Random.org: %s", err)
			os.Exit(3)
		}
		data = diceToStr(dice)
	}
	fmt.Printf("Dice throws:\n%s\n\n", data)

	//calculate entropy
	var entropy []byte = Encryption.GetEntropy(data)
	fmt.Printf("Private key:\n%x\n\n", entropy)

	//get mnemonic from entropy
	mnemonic, err := Encryption.GetMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mnemonic:")
	printMnemonic(mnemonic)
}
