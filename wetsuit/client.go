package wetsuit

type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Origin       string
}

func NewClient(or string, cid string, cs string, at string) *Client {
	return &Client{
		ClientID:     cid,
		ClientSecret: cs,
		Origin:       or,
		AccessToken:  at,
	}
}
