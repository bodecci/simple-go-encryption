package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain text:", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted text:", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted text:", decrypted)
}

func hashLetterFn(key int, letter string) string {
	key = key % len(letter) // Ensure the key is within bounds
	return letter[len(letter)-key:] + letter[:len(letter)-key]
}

func encrypt(key int, plainText string) string {
	hashLetter := hashLetterFn(key, originalLetter)
	return translateText(plainText, originalLetter, hashLetter)
}

func decrypt(key int, encryptedText string) string {
	hashLetter := hashLetterFn(key, originalLetter)
	return translateText(encryptedText, hashLetter, originalLetter)
}

// translateText maps characters from srcAlphabet to dstAlphabet.
func translateText(text, srcAlphabet, dstAlphabet string) string {
	mapping := make(map[rune]rune)
	for i, char := range srcAlphabet {
		mapping[char] = rune(dstAlphabet[i])
	}
	return strings.Map(func(r rune) rune {
		if newChar, ok := mapping[r]; ok {
			return newChar
		}
		return r
	}, text)
}
