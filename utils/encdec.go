package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Encrypt encrypts a string using a 32 bytes key using AES-GCM algorithm.
func Encrypt(stringValue string, keyString string) (encryptedString string) {

	keyHex := hex.EncodeToString([]byte(keyString))
	key, _ := hex.DecodeString(keyHex)
	plaintext := []byte(stringValue)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// Decrypt decrypts a string using a 32 bytes key using AES-GCM algorithm.
func Decrypt(encryptedStringValue string, keyString string) (decryptedString string) {

	keyHex := hex.EncodeToString([]byte(keyString))
	key, _ := hex.DecodeString(keyHex)
	enc, _ := hex.DecodeString(encryptedStringValue)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
