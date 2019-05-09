package sea

type SeaClient struct {
  ClientID      string
  ClientSecret  string
  AccessToken   string
}

func NewSeaClient () *SeaClient {
  return &SeaClient{}
}
