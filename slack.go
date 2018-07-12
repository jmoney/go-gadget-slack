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
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout,
		"[INFO]: ",
		log.Ldate|log.Ltime|log.Llongfile)

	warningLogger = log.New(os.Stdout,
		"[WARNING]: ",
		log.Ldate|log.Ltime|log.Llongfile)

	errorLogger = log.New(os.Stderr,
		"[ERROR]: ",
		log.Ldate|log.Ltime|log.Llongfile)
}

// Send function to post a slack payload to incoming webhook
func (slackClient *Client) Send(payload Payload) (string, error) {
	jsonBytes, err := json.Marshal(payload)
	rawRequestJSON := string(jsonBytes[:])
	req, err := http.NewRequest("POST", slackClient.webhook, strings.NewReader(rawRequestJSON))
	resp, err := slackClient.http.Do(req)
	if err != nil {
		warningLogger.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	rawResponseJSON, err := ioutil.ReadAll(resp.Body)
	infoLogger.Printf("status_code=%s slack_response_raw=%s", resp.Status, rawResponseJSON)
	return resp.Status, nil
}
