package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"nhentaiAgnet/cliReg"
	"os"
)

func main() {
	cliInit()
}
func cliInit() {
	app := &cli.App{
		Name:        "nAgent",
		Usage:       ":)",
		Description: "nhentai comic download agent",
		Version:     "alpha v0.0.1",
		Commands:    cliReg.Commands(),
		Flags:       cliReg.Flags(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
