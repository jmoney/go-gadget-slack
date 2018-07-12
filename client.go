package slack

import "net/http"

type Client struct {
	apiURL string
	http   http.Client
}

func New(http http.Client, url string) *Client {
	return &Client{apiURL: url, http: http}
}
