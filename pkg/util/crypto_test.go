package util

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
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
	encrypted := "aes:Mzwp2NLkHPferlXX6uMI4H7jP+xiTTAJvoZ1zR/2pTGjbVp7xAN6m4k="
	password := "ingress-t.h.i.s.i.s.v.e.r.y.s.e.c.u.r.e.-api"
	assert.Equal(t, "All good?", Open(encrypted, password))
}

func TestSealOpen(t *testing.T) {
	password := "-password-"
	s := "All good?"
	sealed := Seal(s, password)
	assert.True(t, strings.HasPrefix(sealed, Prefix))
	assert.Equal(t, 60, len(sealed))
	assert.Equal(t, s, Open(sealed, password))
}
