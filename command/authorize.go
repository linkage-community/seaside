package command

import (
	"bufio"
	"fmt"
	"os"

	"github.com/otofune/seaside/config"
	"github.com/otofune/seaside/wetsuit"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
)

const (
	AuthorizeState = "omg one state"
)

func ensureTokenByPrompt(config *config.Config) (string, error) {
	authorizeURL, _ := wetsuit.GetAuthorizeURL(config.SeaOrigin, config.ClientID, AuthorizeState)
	fmt.Printf("Start authorization. Open URL in your browser:\n\n%s\n\n", authorizeURL)
	fmt.Printf("Input your authorization code: ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	code := sc.Text()
	fmt.Println()
	token, err := wetsuit.GetToken(config.SeaOrigin, config.ClientID, config.ClientSecret, AuthorizeState, code)

	if err != nil {
		return "", errors.Wrap(err, "can't get token from sea")
	}

	return token, nil
}

func doAuthorize(ctx *cli.Context) error {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't load LoadConfig"))
		os.Exit(1)
	}

	token, err := ensureTokenByPrompt(config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := wetsuit.NewClient(config.SeaOrigin, config.ClientID, config.ClientSecret, token)
	// test...
	r, e := client.Get("/v1/timelines/public")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%s", r)
	}

	return nil
}

var AuthozizeCommand = cli.Command{
	Name:    "authorize",
	Aliases: []string{"a"},
	Usage:   "Make seaside authorized.",
	Action:  doAuthorize,
}
