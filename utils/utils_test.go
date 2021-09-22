package utils

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestString(t *testing.T) {
	str1 := "hello"
	p := String(str1)

	assert.Equal(t, &str1, p)
}
