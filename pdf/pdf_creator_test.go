package pdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPDFCreator(t *testing.T) {
	InitLicense()
	invocie := Invoice{
		OutputPath:  "test.pdf",
		LogoPath:    "https://upload.wikimedia.org/wikipedia/commons/thumb/4/42/Opensource.svg/339px-Opensource.svg.png",
		FontDir:     "https://fonts.google.com/download?family=Roboto",
		FontPackage: "Roboto.zip",
		FontBold:    "Roboto-Bold.ttf",
		FontRegular: "Roboto-Regular.ttf",
		InvoiceName: "Test GmbH",
		Paragraphs: Paragraphs{
			Paragraph{
				Text:     "Test",
				FontSize: 22,
				Margins:  []float64{0, 0, 0, 0},
				Color:    []byte{100, 100, 199},
			},
			Paragraph{
				Text:     "Test2",
				FontSize: 22,
				Margins:  []float64{50, 50, 50, 50},
				Color:    []byte{50, 20, 30},
			},
		},
	}
	err := invocie.Start()
	assert.NoError(t, err)
}
