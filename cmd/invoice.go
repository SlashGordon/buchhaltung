package cmd

import (
	"github.com/SlashGordon/buchhaltung/pdf"
	"github.com/SlashGordon/buchhaltung/types"
	"github.com/SlashGordon/buchhaltung/utils"
	"github.com/spf13/cobra"
)

func init() {

	var (
		configPath string
		inputPath  string
		outputPath string
	)

	startCmd := &cobra.Command{
		Use:   "invoice",
		Short: "Renames invoices based by given config",
		Long:  "Renames invoices based by given config",
		Run: func(cmd *cobra.Command, args []string) {
			pdf.InitLicense()
			config := types.BillRenameItemList{}
			err := config.Unmarshal(configPath)
			if err != nil {
				utils.Logger.Error(err)
			}
			err = config.Rename(inputPath, outputPath)
			if err != nil {
				utils.Logger.Error(err)
			}
		},
	}
	configFlag := startCmd.PersistentFlags()
	configFlag.StringVarP(&configPath, "config", "c", "", "path to config file")
	configFlag.StringVarP(&inputPath, "input", "i", "", "path to pdf directory")
	configFlag.StringVarP(&outputPath, "output", "o", "", "path to output directory")
	cobra.MarkFlagRequired(configFlag, "config")

	RootCmd.AddCommand(startCmd)
}
