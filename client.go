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
