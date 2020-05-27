package main

import (
	"github.com/SlashGordon/buchhaltung/cmd"
	"github.com/SlashGordon/buchhaltung/utils"
)

func main() {
	utils.Logger.Info("Buchhaltung")
	cmd.Execute()
}
