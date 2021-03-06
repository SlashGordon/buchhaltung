package pdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetText(t *testing.T) {
	InitLicense()
	_, err := GetText("")
	assert.Error(t, err)
	invocie := Invoice{
		OutputPath:  "testRead1.pdf",
		LogoPath:    "",
		FontDir:     "https://fonts.google.com/download?family=Roboto",
		FontPackage: "Roboto.zip",
		FontBold:    "Roboto-Bold.ttf",
		FontRegular: "Roboto-Regular.ttf",
		InvoiceName: "Test GmbH",
		Paragraphs: Paragraphs{
			Paragraph{
				Text:     "This is a test",
				FontSize: 22,
				Margins:  []float64{0, 0, 0, 0},
				Color:    []byte{100, 100, 199},
			},
		},
	}
	errPdf := invocie.Start()
	assert.NoError(t, errPdf)
	text, err2 := GetText("testRead1.pdf")
	assert.NoError(t, err2)
	assert.Equal(t, "This is a test\nTest GmbH Page 1 of 1", text)
}
