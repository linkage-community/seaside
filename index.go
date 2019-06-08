package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/otofune/seaside/config"
	"github.com/otofune/seaside/wetsuit"
	"github.com/pkg/errors"
)

const (
	SeaOrigin      = "https://c.linkage.community"
	AuthorizeState = "omg one state"
)

func authorize(c *config.Config) (string, error) {
	authorizeURL, _ := wetsuit.GetAuthorizeURL(SeaOrigin, c.ClientID, AuthorizeState)
	fmt.Printf("Start authorization. Open URL in your browser:\n\n%s\n\n", authorizeURL)
	fmt.Printf("Input your authorization code: ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	code := sc.Text()
	fmt.Println()
	token, err := wetsuit.GetToken(SeaOrigin, c.ClientID, c.ClientSecret, AuthorizeState, code)

	if err != nil {
		return "", errors.Wrap(err, "can't get token from sea")
	}

	return token, nil
}

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Can't load LoadConfig: %v\n", err)
		os.Exit(1)
	}

	token, err := authorize(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(token)
}
