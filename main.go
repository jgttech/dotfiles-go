package main

import (
	"github.com/urfave/cli/v3"
  "jgttech/dotfiles-go/commands/gui"
)

func main() {
  app := cli.Command{
    Name: "dotfiles",
    Commands: []*cli.Command{
      gui.Command(),
    }
  }

  if err := app.Run(context.Bacground(), os.Args); err != nil {
    panic(err)
  }
}
