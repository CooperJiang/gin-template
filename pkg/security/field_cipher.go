package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

type FieldCipher struct {
	aead cipher.AEAD
}

func NewFieldCipher(secret string) (*FieldCipher, error) {
	trimmed := strings.TrimSpace(secret)
	if trimmed == "" {
		return nil, fmt.Errorf("encryption key is required")
	}

	key := sha256.Sum256([]byte(trimmed))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &FieldCipher{aead: aead}, nil
}

func (c *FieldCipher) Encrypt(plainText string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("field cipher is nil")
	}
	if plainText == "" {
		return "", nil
	}

	nonce := make([]byte, c.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	sealed := c.aead.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(sealed), nil
}

func (c *FieldCipher) Decrypt(cipherText string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("field cipher is nil")
	}
	if strings.TrimSpace(cipherText) == "" {
		return "", nil
	}

	raw, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	nonceSize := c.aead.NonceSize()
	if len(raw) < nonceSize {
		return "", fmt.Errorf("invalid ciphertext length")
	}

	nonce := raw[:nonceSize]
	payload := raw[nonceSize:]
	plain, err := c.aead.Open(nil, nonce, payload, nil)
	if err != nil {
		return "", err
	}

	return string(plain), nil
}
