package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/SlashGordon/buchhaltung/pdf"
	"github.com/SlashGordon/buchhaltung/utils"
)

// RenameConfig ...
type RenameConfig interface {
	Unmarshal(path string) error
	Rename(input string, output string) error
}

// BillRenameItem ...
type BillRenameItem struct {
	Identifyers map[string]string `json:"identifyers"`
	OutputName  string            `json:"outputname"`
}

// BillRenameItemList ...
type BillRenameItemList []BillRenameItem

//Unmarshal config for bill items
func (b *BillRenameItemList) Unmarshal(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	return nil
}

// Rename ...
func (b *BillRenameItemList) Rename(input string, output string) error {
	if !utils.DirExists(input) {
		return fmt.Errorf("Input dir %v doesn't exist", input)
	}

	if !utils.DirExists(output) {
		os.MkdirAll(output, os.ModePerm)
	}
	notfound := path.Join(output, "notfound")
	if !utils.DirExists(notfound) {
		os.MkdirAll(notfound, os.ModePerm)
	}

	pdfFiles, err := utils.WalkMatch(input, "*.pdf")

	if err != nil {
		return err
	}

	for _, pdfFile := range pdfFiles {
		found := false
		pdfText, errParse := pdf.GetText(pdfFile)
		if errParse != nil {
			return errParse
		}
		for _, item := range *b {
			outputName := item.OutputName
			match := true
			matches := make(map[string]string, len(item.Identifyers))
			for k, v := range item.Identifyers {
				r, regexErr := regexp.Compile(v)
				if regexErr != nil {
					return regexErr
				}
				stringMatches := r.FindStringSubmatch(pdfText)
				if len(stringMatches) >= 2 {
					matches[k] = stringMatches[1]
				} else {
					match = false
				}
			}

			if match {
				for k, v := range matches {
					outputName = strings.ReplaceAll(outputName, fmt.Sprintf("{%v}", k), v)
				}
				outputName = path.Join(output, outputName)
				utils.Logger.Info(fmt.Sprintf("Move %v to %v", pdfFile, outputName))
				err := os.Rename(pdfFile, outputName)
				if err != nil {
					return err
				}
				found = true
			}
		}
		if !found {

			err := os.Rename(pdfFile, path.Join(notfound, filepath.Base(pdfFile)))
			if err != nil {
				return err
			}
			utils.Logger.Warn(fmt.Sprintf("Pdf %v has no match", pdfFile))
		}
	}

	return nil
}
