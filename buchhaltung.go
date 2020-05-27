package main

import (
	"fmt"

	"github.com/SlashGordon/buchhaltung/cmd"
)

func main() {
	const banner = `
___               _      _             _   _
| _ )  _  _   __  | |_   | |_    __ _  | | | |_   _  _   _ _    __ _
| _ \ | || | / _| | ' \  | ' \  / _| | | | |  _| | || | | ' \  / _| |
|___/  \_,_| \__| |_||_| |_||_| \__,_| |_|  \__|  \_,_| |_||_| \__, |
                                                               |___/
`
	fmt.Printf(banner)
	cmd.Execute()
}
