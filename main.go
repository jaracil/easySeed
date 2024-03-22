package main

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Prompt user for 12 or 24 words
	prompt12 := promptui.Select{
		Label: "Select 12 or 24 words",
		Items: []string{"12 words", "24 words"},
	}
	i, _, err := prompt12.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	// Prompt user for passPhrase with hidden input
	promptPass := promptui.Prompt{
		Label: "Enter pass phrase",
		Mask:  '*',
	}
	passPhrase1, err := promptPass.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	promptPass2 := promptui.Prompt{
		Label: "Enter pass phrase again to confirm",
		Mask:  '*',
	}

	passPhrase2, err := promptPass2.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if passPhrase1 != passPhrase2 {
		fmt.Println("Pass phrases do not match")
		return
	}

	// Get sha256 hash of the passPhrase
	hash256 := sha256.Sum256([]byte(passPhrase1))
	var entropy []byte
	if i == 0 {
		entropy = hash256[:16]
	} else {
		entropy = hash256[:32]
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	words := strings.Fields(mnemonic)
	for i, word := range words {
		fmt.Printf("%02d: %s\n", i+1, word)
	}
}
