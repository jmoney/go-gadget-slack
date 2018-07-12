package slack

import "net/http"

// Client struct for use built via constructor
type Client struct {
	webhook string
	http    http.Client
}

// New Constructor for the client
func New(http http.Client, url string) *Client {
	return &Client{webhook: url, http: http}
}
