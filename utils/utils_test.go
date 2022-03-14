package utils

import (
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestExtractParamString(t *testing.T) {
	t.Skip("// TODO: Implement")
}

func TestExtractParamUint(t *testing.T) {
	t.Skip("// TODO: Implement")
}

func TestExtractParamInt(t *testing.T) {
	t.Skip("// TODO: Implement")
}

// TODO: Move to github.com/stefanoschrs/go-utils
func TestString(t *testing.T) {
	str1 := "hello"
	p := String(str1)

	assert.Equal(t, &str1, p)
}

// TODO: Move to github.com/stefanoschrs/go-utils
func TestSliceUintToString(t *testing.T) {
	arr1 := []uint{0, 1, 2, 3, 4}
	arr2 := SliceUintToString(arr1)

	assert.Equal(t, len(arr1), len(arr2))

	for i, s := range arr2 {
		val, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, val, arr1[i])
	}
}
