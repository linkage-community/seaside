package command

import (
	"bufio"
	"fmt"
	"os"

	"github.com/otofune/seaside/config"
	"github.com/otofune/wetsuit"
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
	if err := sc.Err(); err != nil {
		fmt.Println(errors.Wrap(err, "Can't read input"))
		os.Exit(1)
	}

	code := sc.Text()
	token, err := wetsuit.GetToken(config.SeaOrigin, config.ClientID, config.ClientSecret, AuthorizeState, code)

	if err != nil {
		return "", errors.Wrap(err, "can't get token from sea")
	}

	return token, nil
}

func doAuthorize(ctx *cli.Context) error {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't load config"))
		os.Exit(1)
	}

	token, err := ensureTokenByPrompt(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// padding :innocent:
	fmt.Println()

	if err := c.SaveCredential(&config.Credential{AccessToken: token}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("Welcome to the seaside.")

	return nil
}

var AuthozizeCommand = cli.Command{
	Name:    "authorize",
	Aliases: []string{"a"},
	Usage:   "Make seaside authorized.",
	Action:  doAuthorize,
}
