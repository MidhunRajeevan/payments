package util

import (
	"MidhunRajeevan/payments/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
)

func GenerateTransactionToken() (uid string, token string) {
	uuid := uuid.New()
	token, err := Encrypt(uuid.String())
	if err != nil {
		log.Println("Encrypt Error for transaction token", err.Error())
	}

	return uuid.String(), token
}

// Encrypt
func Encrypt(plaintext string) (string, error) {
	key := []byte(config.App.SecretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error during encryption :", err.Error())
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error during encryption :", err.Error())
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println("Error during encryption :", err.Error())
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	cipherhex := fmt.Sprintf("%x", ciphertext)
	return cipherhex, nil
}

// Decrypt
func Decrypt(cipherhex string) (string, error) {
	key := []byte(config.App.SecretKey)
	ciphertext, err := hex.DecodeString(cipherhex)
	if err != nil {
		log.Println("Error during decryption, at hex.DecodeString. Error not displayed.")
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error during decryption :", err.Error())
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error during decryption :", err.Error())
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Println("Error during decryption :", "Ciphertext Length Error")
		return "", errors.New("ciphertext_error")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		log.Println("Error during decryption :", err.Error())
		return "", err
	}
	return string(plaintext), nil
}
