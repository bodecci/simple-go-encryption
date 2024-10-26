package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain text", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encypted text: ", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOversLetter := string(runes[0 : len(letter)-key])

	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOversLetter)
}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""
	findOne := func(r rune) rune {

	}

	strings.Map(findOne, plainText)
}

func decrypt(key int, encryptedText string) (result string) {

}
