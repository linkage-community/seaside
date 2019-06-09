package main

import (
	"os"

	"github.com/otofune/seaside/command"
	"github.com/otofune/seaside/config"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "seaside"
	app.Usage = "A brazing simple rinsuki/sea client."
	app.Commands = command.Commands
	app.Version = config.Version
	app.Run(os.Args)
}
