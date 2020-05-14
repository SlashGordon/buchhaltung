package pdf

import (
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

func CreatePdf(outputPath string) {

}

//GetText
func GetText(path string) (string, error) {
	var sb strings.Builder
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return "", err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return "", err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return "", err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	for pageNum := 1; pageNum <= numPages; pageNum++ {

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return "", err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return "", err
		}
		sb.WriteString(text)

	}

	return sb.String(), nil
}
