package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"strings"
)

func Sha256(s string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return hasher.Sum(nil)
}

const NonceSize = 16
const TagSize = 16

func RandomNonce() []byte {
	nonce := make([]byte, NonceSize)
	Must(io.ReadFull(rand.Reader, nonce))
	return nonce
}

type Gcm struct {
	Data  []byte
	Tag   []byte
	Nonce []byte
}

func GcmEncrypt(key []byte, plaintext []byte, nonce []byte) Gcm {
	if nonce == nil {
		nonce = RandomNonce()
	}
	block := Must(aes.NewCipher(key))
	aesgcm := Must(cipher.NewGCMWithNonceSize(block, NonceSize))
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return Gcm{
		Data:  ciphertext[:len(ciphertext)-TagSize],
		Tag:   ciphertext[len(ciphertext)-TagSize:],
		Nonce: nonce,
	}
}

func GcmDecrypt(key []byte, gcm Gcm) []byte {
	block := Must(aes.NewCipher(key))
	aesgcm := Must(cipher.NewGCMWithNonceSize(block, NonceSize))
	sealed := append(gcm.Data, gcm.Tag...)
	return Must(aesgcm.Open(nil, gcm.Nonce, sealed, nil))
}

const Prefix = "aes:"

func Open(s string, password string) string {
	key := Sha256(password)

	if !strings.HasPrefix(s, Prefix) {
		return s
	}

	s = strings.TrimPrefix(s, Prefix)
	r := Must(base64.StdEncoding.DecodeString(s))

	nonce := r[0:NonceSize]
	tag := r[NonceSize : NonceSize+TagSize]
	data := r[NonceSize+TagSize:]
	gcm := Gcm{Data: data, Tag: tag, Nonce: nonce}

	return string(GcmDecrypt(key, gcm))
}

func Seal(s string, password string) string {
	key := Sha256(password)
	gcm := GcmEncrypt(key, []byte(s), nil)
	r := make([]byte, 0, len(gcm.Tag)+len(gcm.Nonce)+len(gcm.Data))
	r = append(r, gcm.Nonce...)
	r = append(r, gcm.Tag...)
	r = append(r, gcm.Data...)
	return Prefix + base64.StdEncoding.EncodeToString(r)
}
