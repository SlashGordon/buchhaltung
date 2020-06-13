package pdfrename

import (
	"fmt"
	"os"
	"path"
	"strings"

	ty "github.com/SlashGordon/buchhaltung/types"
	"github.com/SlashGordon/buchhaltung/utils"
)

// Start ...
func Start(conf ty.RenameConfig, inputPath string, outputPath string) error {
	// pre processing
	// docker pull jbarlow83/ocrmypdf
	// docker run -v $PWD:/data --workdir /data --rm -i jbarlow83/ocrmypdf 201912151411.pdf 201912151411ocr.pdf
	ocrPath := path.Join(inputPath, "ocr")
	if !utils.DirExists(ocrPath) {
		os.MkdirAll(ocrPath, os.ModePerm)
	}
	pdfFiles, err := utils.WalkMatch(inputPath, "*.pdf")
	if utils.Cmd("docker", "pull", "jbarlow83/ocrmypdf:latest") == 0 {
		for _, pdfFile := range pdfFiles {
			relPdfFile := strings.Replace(pdfFile, inputPath, "", 1)
			if strings.HasPrefix(relPdfFile, "/") {
				relPdfFile = strings.Replace(relPdfFile, "/", "", 1)
			}
			if utils.Cmd("docker", "run", "-v",
				fmt.Sprintf("%v:/data", inputPath), "--workdir", "/data",
				"--rm", "-i", "jbarlow83/ocrmypdf",
				relPdfFile, path.Join("ocr", relPdfFile)) == 0 {
				utils.Logger.Infof("OCR for file %v was successful.", pdfFile)
			} else {
				utils.Logger.Warnf("OCR for file %v was unsuccessful.", pdfFile)
			}
		}
	} else {
		utils.Logger.Warn("Couldn't start ocrmypdf because of missing docker dependencies.")
	}
	err = conf.Rename(inputPath, outputPath)
	return err
}
