package pdfrename

import (
	ty "github.com/SlashGordon/buchhaltung/types"
)

// Start ...
func Start(conf ty.RenameConfig, inputPath string, outputPath string) error {
	err := conf.Rename(inputPath, outputPath)
	return err
}
