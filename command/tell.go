package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/otofune/seaside/config"
	"github.com/otofune/wetsuit"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
)

func doTell(ctx *cli.Context) error {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't load LoadConfig"))
		os.Exit(1)
	}
	if err := c.LoadCurrentCredential(); err != nil {
		fmt.Println(errors.Wrap(err, "You must authenticate before call"))
	}

	tb, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't read input"))
		os.Exit(1)
	}
	text := strings.TrimSpace(string(tb))

	client := wetsuit.NewClient(c.SeaOrigin, c.ClientID, c.ClientSecret, c.CurrentCredential.AccessToken)
	id, err := client.CreatePost(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)

	return nil
}

var TellCommand = cli.Command{
	Name:    "tell",
	Aliases: []string{"t"},
	Usage:   "Tell your words",
	Action:  doTell,
}
