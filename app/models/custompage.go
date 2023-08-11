package models

type CustomPage struct {
	Board   string                `json:"board"`
	Page    string                `json:"page"`
	Title   string                `json:"title"`
	Message RawAndMarkdownMessage `json:"message"`
	Date    ScuffedTime           `json:"date"`
	Edited  interface{}           `json:"edited"`
	ID      string                `json:"_id"`
}
