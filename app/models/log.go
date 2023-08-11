package models

import (
	"fmt"
)

type LogList struct {
	Date  LogDate `json:"date"`
	Count int     `json:"count"`
}

type Log struct {
	ID        string         `json:"_id"`
	ShowLinks bool           `json:"showLinks"`
	PostLinks []LogPostLinks `json:"postLinks"`
	Actions   []string       `json:"actions"`
	Date      ScuffedTime    `json:"date"`
	ShowUser  bool           `json:"showUser"`
	Message   string         `json:"message"`
	User      string         `json:"user"`
	Board     string         `json:"board"`
}

type LogPostLinks struct {
	PostID int    `json:"postId"`
	Thread int    `json:"thread,omitempty"`
	Board  string `json:"board,omitempty"`
}

type LogDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (m *LogDate) String() string {
	return fmt.Sprintf("%02d-%02d-%02d", m.Month, m.Day, m.Year)
}
