package slack

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	// Info Logger
	Info *log.Logger
	// Warning Logger
	Warning *log.Logger
	// Error Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout,
		"[INFO]: ",
		log.Ldate|log.Ltime|log.Llongfile)

	Warning = log.New(os.Stdout,
		"[WARNING]: ",
		log.Ldate|log.Ltime|log.Llongfile)

	Error = log.New(os.Stderr,
		"[ERROR]: ",
		log.Ldate|log.Ltime|log.Llongfile)
}

func (slackClient *Client) send(incomingWebhook string, payload Payload) (string, error) {
	jsonBytes, err := json.Marshal(payload)
	rawRequestJSON := string(jsonBytes[:])
	req, err := http.NewRequest("POST", incomingWebhook, strings.NewReader(rawRequestJSON))
	resp, err := slackClient.http.Do(req)
	if err != nil {
		Warning.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	rawResponseJSON, err := ioutil.ReadAll(resp.Body)
	Info.Printf("status_code=%s slack_response_raw=%s", resp.Status, rawResponseJSON)
	return resp.Status, nil
}
