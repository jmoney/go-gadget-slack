package slack

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Send function to post a slack payload to incoming webhook
func (slackClient *Client) Send(payload Payload) (string, error) {
	jsonBytes, err := json.Marshal(payload)
	rawRequestJSON := string(jsonBytes[:])
	req, err := http.NewRequest("POST", slackClient.webhook, strings.NewReader(rawRequestJSON))
	resp, err := slackClient.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
