/*
Copyright 2018 Jonathan Monette

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
