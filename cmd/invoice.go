package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

	var (
		configPath string
	)

	startCmd := &cobra.Command{
		Use:   "invoice",
		Short: "Renames invoices based by given config",
		Long:  "Renames invoices based by given config",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	configFlag := startCmd.PersistentFlags()
	configFlag.StringVarP(&configPath, "config", "c", "", "path to config file")
	cobra.MarkFlagRequired(configFlag, "config")

	RootCmd.AddCommand(startCmd)
}
