package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.Equal(t, Must(100, nil), 100)
}
