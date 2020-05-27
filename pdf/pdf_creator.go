package pdf

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/SlashGordon/buchhaltung/utils"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type Paragraph struct {
	Text     string    `json:"text"`
	Font     string    `json:"font"`
	FontSize float64   `json:"fontsize"`
	Margins  []float64 `json:"margins"`
	Color    []byte    `json:"color"`
}

type Paragraphs []Paragraph

// Invoice ...
type Invoice struct {
	OutputPath  string      `json:"outputname"`
	LogoPath    string      `json:"logopath"`
	FontDir     string      `json:"fontdir"`
	FontPackage string      `json:"fontpackage"`
	FontBold    string      `json:"fontbold"`
	FontRegular string      `json:"fontregular"`
	InvoiceName string      `json:"invoicename"`
	Paragraphs  []Paragraph `json:"paragraphs"`
}

// InvoiceConfig ...
type InvoiceConfig interface {
	Start() error
}

// Start ...
func (p *Invoice) Start() error {
	logoPath := p.LogoPath
	if utils.IsValidURL(p.LogoPath) {
		logoPath = path.Base(p.LogoPath)
		if !utils.Exists(logoPath) {
			utils.DownloadFile(logoPath, p.LogoPath)
		}
	}

	fontDir := p.FontDir
	if utils.IsValidURL(p.FontDir) {
		if strings.HasSuffix(strings.ToLower(p.FontPackage), ".zip") {
			fontDir = strings.Replace(p.FontPackage, ".zip", "", 1)
		} else {
			return errors.New("Given font archive is not supported yes")
		}

		if !utils.Exists(fontDir) {
			utils.DownloadFile(p.FontPackage, p.FontDir)
			utils.Unzip(p.FontPackage, fontDir)
		}
	}

	font, err := model.NewPdfFontFromTTFFile(path.Join(fontDir, p.FontRegular))
	if err != nil {
		return err
	}

	c := creator.New()
	c.SetPageMargins(50, 50, 100, 70)

	// Setup a front page (always placed first).
	c.CreateFrontPage(func(args creator.FrontpageFunctionArgs) {
		for _, para := range p.Paragraphs {
			myP := c.NewParagraph(para.Text)
			myP.SetFont(font)
			myP.SetFontSize(para.FontSize)
			myP.SetMargins(para.Margins[0], para.Margins[1], para.Margins[2], para.Margins[3])
			myP.SetColor(creator.ColorRGBFrom8bit(para.Color[0], para.Color[1], para.Color[2]))
			c.Draw(myP)
		}
	})

	// Draw a header on each page.
	c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
		if logoPath == "" {
			logoImg, err := c.NewImageFromFile(logoPath)
			if err == nil {
				logoImg.ScaleToHeight(25)
				logoImg.SetPos(58, 20)
				block.Draw(logoImg)
			}
		}
	})

	// Draw footer on each page.
	c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		// Draw the on a block for each page.
		p := c.NewParagraph(p.InvoiceName)
		p.SetFont(font)
		p.SetFontSize(8)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)

		strPage := fmt.Sprintf("Page %d of %d", args.PageNum, args.TotalPages)
		p = c.NewParagraph(strPage)
		p.SetFont(font)
		p.SetFontSize(8)
		p.SetPos(300, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)
	})

	dirPat := filepath.Dir(p.OutputPath)
	if !utils.DirExists(dirPat) {
		os.MkdirAll(dirPat, os.ModePerm)
	}

	err = c.WriteToFile(p.OutputPath)
	if err != nil {
		return err
	}

	return nil
}
