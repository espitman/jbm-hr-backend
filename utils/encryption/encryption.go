package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/espitman/jbm-hr-backend/utils/config"
)

var (
	key []byte
)

func init() {
	// Get encryption key from config
	keyStr := config.GetConfig("ENCRYPTION_KEY", "jabama1404!!!!")

	// Create a 32-byte key using SHA-256
	hash := sha256.Sum256([]byte(keyStr))
	key = hash[:]
}

// Encrypt encrypts a string using AES-256-GCM with a deterministic nonce
func Encrypt(text string) (string, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Create a new GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	// Create a deterministic nonce from the input text
	nonceHash := sha256.Sum256([]byte(text))
	nonce := nonceHash[:gcm.NonceSize()]

	// Encrypt the text
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)

	// Encode the result in base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts an encrypted string using AES-256-GCM
func Decrypt(encryptedText string) (string, error) {
	// Decode the base64 string
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %v", err)
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Create a new GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	// Extract the nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the text
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %v", err)
	}

	return string(plaintext), nil
}
