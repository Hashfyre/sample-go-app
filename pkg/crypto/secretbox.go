package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"

	"golang.org/x/crypto/nacl/secretbox"
)

// Encrypt defines the encryption procedure for any given secret string
func Encrypt(secretKey string, rawValue string) (string, error) {
	var secretKeyBytes [32]byte
	copy(secretKeyBytes[:], secretKey)
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		log.Println(err)
		return "", err
	}

	encrypted := secretbox.Seal(nonce[:], []byte(rawValue), &nonce, &secretKeyBytes)
	encryptedEncoded := base64.StdEncoding.EncodeToString(encrypted)
	return encryptedEncoded, nil
}

// Decrypt defines the decryption procedure for any given encrypted string
func Decrypt(secretKey string, encrypted string) (string, error) {
	var decryptNonce [24]byte
	var secretKeyBytes [32]byte
	copy(secretKeyBytes[:], secretKey)
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open(nil, []byte(encrypted[24:]), &decryptNonce, &secretKeyBytes)
	if !ok {
		err := errors.New("Unable to decrypt string")
		log.Println(err)
		return "", err
	}
	return string(decrypted), nil
}
