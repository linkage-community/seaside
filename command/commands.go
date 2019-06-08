package command

import (
	"github.com/otofune/seaside/config"
	"gopkg.in/urfave/cli.v1"
)

var Commands = []cli.Command{
	{
		Name:    "show-environment-variables",
		Aliases: []string{"sev"},
		Usage:   "Check environment variables whehter used by seaside.",
		Action: func(c *cli.Context) error {
			config.Usage()
			return nil
		},
	},
	AuthozizeCommand,
}
