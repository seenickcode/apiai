package apiai

const (
	APIAIBaseURL = "https://api.api.ai/v1"
)

type Client struct {
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
	}
}
