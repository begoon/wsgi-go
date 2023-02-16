package util

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	r := Sha256("abc")
	assert.Equal(t, 32, len(r))
	h := hex.EncodeToString(r)
	e := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	assert.Equal(t, e, h)
}

func TestGCM(t *testing.T) {
	key := Sha256("-key-")
	secret := "-secret-"
	gcm := GcmEncrypt(key, []byte(secret), nil)
	r := GcmDecrypt(key, Gcm{Data: gcm.Data, Tag: gcm.Tag, Nonce: gcm.Nonce})
	assert.Equal(t, secret, string(r))
}

func TestOpen(t *testing.T) {
	encrypted := "aes:sYaV5b8+w7OMkZBInQCiPqaYc56ZjGIDD8OOSxrYQyXUe0Uc3Lb9I0U="
	password := "-secret-"
	assert.Equal(t, "All good?", Open(encrypted, password))
}

func TestOpenWithoutPrefix(t *testing.T) {
	plain := "-plain-"
	password := "-secret-"
	assert.Equal(t, "-plain-", Open(plain, password))
}

func TestSealOpen(t *testing.T) {
	password := "-password-"
	s := "All good?"
	sealed := Seal(s, password)
	assert.True(t, strings.HasPrefix(sealed, Prefix))
	assert.Equal(t, 60, len(sealed))
	assert.Equal(t, s, Open(sealed, password))
}
