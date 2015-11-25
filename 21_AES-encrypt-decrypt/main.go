package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	// AES key, either 16, 24, 0r 32 bytes to select AES-128, AES192, or AES-256
	key := "opensesame123456" // 16 bytes
	block, _:=aes.NewCipher([]byte(key))
	fmt.Printf("%d bytes NewCipher key with block size of %d bytes/n", len(key), block.BlockSize)
	str := []byte("Hello World, and everyone else in the universe!")
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] //const BlockSize = 16
	// encrypt
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)
	fmt.Printf("%s encrypted to %v\n", str, encrypted)
	// decrypt
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(str))
	decrypter.XORKeyStream(decrypted, encrypted)
	fmt.Printf("%v decrypt to %s\n", encrypted, decrypted)
}

