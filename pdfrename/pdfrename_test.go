package pdfrename

import (
	"os"
	"path"
	"testing"

	"github.com/SlashGordon/buchhaltung/pdf"
	ty "github.com/SlashGordon/buchhaltung/types"
	"github.com/SlashGordon/buchhaltung/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	retCode := m.Run()
	os.RemoveAll("test")
	os.RemoveAll("output")
	os.Exit(retCode)
}

func TestStart(t *testing.T) {
	pdf.InitLicense()
	conf := ty.BillRenameItemList{
		ty.BillRenameItem{
			OutputName: "{number}_{company}.pdf",
			Identifyers: map[string]string{
				"number":  "ReNr: (\\d{10})",
				"company": "(REWE)",
			},
		},
	}

	invocie := pdf.Invoice{
		OutputPath:  "test/invoice1.pdf",
		LogoPath:    "",
		FontDir:     "https://fonts.google.com/download?family=Roboto",
		FontPackage: "Roboto.zip",
		FontBold:    "Roboto-Bold.ttf",
		FontRegular: "Roboto-Regular.ttf",
		InvoiceName: "REWE",
		Paragraphs: pdf.Paragraphs{
			pdf.Paragraph{
				Text:     "ReNr: 4587694524",
				FontSize: 22,
				Margins:  []float64{0, 0, 0, 0},
				Color:    []byte{100, 100, 199},
			},
		},
	}
	errPdf := invocie.Start()

	assert.NoError(t, errPdf)
	err := Start(&conf, "testttttt", "output")
	assert.Error(t, err)

	err = Start(&conf, "test", "output")
	assert.NoError(t, err)
	assert.True(t, utils.FileExists(path.Join("output", "4587694524_REWE.pdf")))
}
