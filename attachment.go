package main

// Field in attachemt
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// MatchAttachment struct, use in Slack
type MatchAttachment struct {
	Color      string  `json:"color"`
	Title      string  `json:"title"`
	Text       string  `json:"text"`
	Fields     []Field `json:"fields"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	Timestamp  int64   `json:"ts"`
}
