package command

import (
	"fmt"
	"os"

	"github.com/otofune/seaside/config"
	"github.com/otofune/wetsuit"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
)

func doGetPublicTimeline(ctx *cli.Context) error {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't load LoadConfig"))
		os.Exit(1)
	}
	if err := c.LoadCurrentCredential(); err != nil {
		fmt.Println(errors.Wrap(err, "You must authenticate before call"))
	}

	client := wetsuit.NewClient(c.SeaOrigin, c.ClientID, c.ClientSecret, c.CurrentCredential.AccessToken)
	pp, err := client.GetTimeline("public", wetsuit.Limit(ctx.Int("limit")), wetsuit.SinceID(ctx.Int("since-id")), wetsuit.MaxID(ctx.Int("max-id")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	posts := *pp

	if len(posts) == 0 {
		fmt.Println("There are no posts.")
	}
	for _, p := range posts {
		fmt.Printf("=== %d ===\n\tText: %s\n\tAppl: %v\n\tUser: %v\n\t\tAvatar: %v\n", p.ID, p.Text, p.Application, p.User, p.User.AvatarFile)
		if len(p.Files) != 0 {
			fmt.Printf("\tFiles:\n")
		}
		for _, f := range p.Files {
			fmt.Printf("\t\t%v\n", f)
		}
		fmt.Println()
	}

	return nil
}

var GetPublicTimelineCommand = cli.Command{
	Name:    "public-timeline",
	Aliases: []string{"ptl"},
	Usage:   "Get posts in public timeline",
	Action:  doGetPublicTimeline,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "limit, l",
			Usage: "LIMIT N",
			Value: 20,
		},
		cli.IntFlag{
			Name:  "max-id, m",
			Usage: "WHERE id < m",
		},
		cli.IntFlag{
			Name:  "since-id, s",
			Usage: "WHERE id > s",
		},
	},
}
