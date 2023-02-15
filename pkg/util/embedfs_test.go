package util

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed fs/*
var fs embed.FS

func TestStringFS(t *testing.T) {
	var b []byte
	StringFS(&b, "fs", fs)
	assert.Equal(t, "fs/folder/two (11)\nfs/one (4)\n", string(b))
}
