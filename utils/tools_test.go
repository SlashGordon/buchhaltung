package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	assert.False(t, IsValidURL("())"))
	assert.False(t, IsValidURL("/test/to/file.png"))
	assert.False(t, IsValidURL("www.google.com"))
	assert.True(t, IsValidURL("http://google.com"))
}
