package main

import "fmt"

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain text", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encypted text: ", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted text: ", decrypted)
}

func encrypt(key int, plainText string) (result string) {

}

func decrypt(key int, encryptedText string) (result string) {

}
