package models

type Notification struct {
	title   string `json:"title"`
	content string `json:"content"`
	link    string `json:"link"`
}
