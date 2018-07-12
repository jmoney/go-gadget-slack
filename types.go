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
