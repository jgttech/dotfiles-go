package gui

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "gui",
		Usage: "Run the Dotfiles GUI.",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("Hello world")
			return nil
		},
	}
}
