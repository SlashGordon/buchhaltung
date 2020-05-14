package main

import (
	"flag"
	"os"

	"github.com/SlashGordon/buchhaltung/pdf"
)

func main() {
	confPathPtr := flag.String("config", "", "Path to config file. (Required)")
	flag.Parse()

	if *confPathPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	pdf.GetText("")
}
