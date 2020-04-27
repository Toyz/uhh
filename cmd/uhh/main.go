package main

import (
	"strings"
	"log"
	"fmt"
	"os"

	"github.com/theprimeagen/uhh"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cfg, created := getConfig()

    if created && !gitClone(cfg) {
        log.Fatal("unable to get config")
    }

	uhh := uhh.New(cfg)
	ucli := newUhhCli(uhh)

	app := &cli.App{
		Name:   "uhh",
		Usage:  "find commands from your repo",
		Action: ucli.findHandler,
		Commands: []cli.Command{
			{Name: "add", Action: ucli.addHandler},
			{Name: "delete", Action: ucli.deleteHandler},
		},
	}

    err := app.Run(os.Args)

    if err != nil {
        fmt.Printf("%+v\n", err)
        os.Exit(1)
    }
}

func readTermLine() string {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatal("could not change state of terminal")
	}
	defer terminal.Restore(0, oldState)
	t := terminal.NewTerminal(os.Stdin, ">")

	line, err := t.ReadLine()
	if err != nil {
		log.Fatal("unable to read line")
	}

	return strings.TrimSpace(line)
}
