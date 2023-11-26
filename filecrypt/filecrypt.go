package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"os"
)

func Encrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		log.Fatalf("Source file not found: %s", err)
	}

	plainText, err := os.ReadFile(source)
	if err != nil {
		log.Fatalf("Error reading source file: %s", err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("Error generating nonce: %s", err)
	}

	key := password
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %s", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Error creating GCM cipher: %s", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plainText, nil)
	ciphertext = append(ciphertext, nonce...)

	err = os.WriteFile(source, ciphertext, 0644)
	if err != nil {
		log.Fatalf("Error writing to destination file: %s", err)
	}
}

func Decrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		log.Fatalf("Source file not found: %s", err)
	}

	cipherText, err := os.ReadFile(source)
	if err != nil {
		log.Fatalf("Error reading source file: %s", err)
	}

	salt := cipherText[len(cipherText)-12:]
	nonce, err := hex.DecodeString(hex.EncodeToString(salt))
	if err != nil {
		log.Fatalf("Error decoding nonce: %s", err)
	}

	key := password
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %s", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Error creating GCM cipher: %s", err)
	}

	plainText, err := aesgcm.Open(nil, nonce, cipherText[:len(cipherText)-12], nil)
	if err != nil {
		log.Fatalf("Error decrypting: %s", err)
	}

	err = os.WriteFile(source, plainText, 0644)
	if err != nil {
		log.Fatalf("Error writing to destination file: %s", err)
	}
}
