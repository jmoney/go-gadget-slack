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

// Payload represent a payload passed to slack on the incoming webhook
type Payload struct {
	Attachments []Attachment `json:"attachments"`
	Channel     string       `json:"channel"`
}

// Attachment represents an attachment posted to an inbound webhook
type Attachment struct {
	Color           string            `json:"color,omitempty"`
	Pretext         string            `json:"pretext,omitempty"`
	AuthorName      string            `json:"author_name,omitempty"`
	AuthorLink      string            `json:"author_link,omitempty"`
	AuthorIcon      string            `json:"author_icon,omitempty"`
	Title           string            `json:"title,omitempty"`
	TitleLink       string            `json:"title_link,omitempty"`
	Text            string            `json:"text,omitempty"`
	AttachmentField []AttachmentField `json:"fields,omitempty"`
	ImageURL        string            `json:"image_url,omitempty"`
	ThumbURL        string            `json:"thumb_url,omitempty"`
	Footer          string            `json:"footer,omitempty"`
	FooterIcon      string            `json:"footer_icon,omitempty"`
	Ts              int64             `json:"ts,omitempty"`
}

// AttachmentField represents a field in the slack attachment fields array
type AttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
