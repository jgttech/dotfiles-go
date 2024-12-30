package main

import (
	"context"
	"jgttech/dotfiles-go/commands/gui"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.Command{
		Name: "dotfiles",
		Commands: []*cli.Command{
			gui.Command(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
