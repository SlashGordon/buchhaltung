package pdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetText(t *testing.T) {
	InitLicense()
	_, err := GetText("")
	assert.Error(t, err)
	text, err2 := GetText("../r1.pdf")
	assert.NoError(t, err2)
	assert.Equal(t, "This is a test ", text)
}
