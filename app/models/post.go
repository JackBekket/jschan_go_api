package models

import (
	"encoding/json"
	"time"
)

type ScuffedTime time.Time

func (m *ScuffedTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` || string(data) == "0" {
		return nil
	}
	return json.Unmarshal(data, (*time.Time)(m))
}

type Post struct {
	ID               string        `json:"_id"`
	Date             ScuffedTime   `json:"date"`
	U                int64         `json:"u"`
	Name             string        `json:"name"`
	Country          Country       `json:"country"`
	Board            string        `json:"board"`
	Tripcode         string        `json:"tripcode"`
	Capcode          string        `json:"capcode"`
	Subject          string        `json:"subject"`
	Message          string        `json:"message"`
	Messagehash      string        `json:"messagehash"`
	Nomarkup         string        `json:"nomarkup"`
	Thread           int           `json:"thread"`
	Email            string        `json:"email"`
	Spoiler          bool          `json:"spoiler"`
	Banmessage       string        `json:"banmessage"`
	UserID           string        `json:"userId"`
	Files            []Files       `json:"files"`
	Quotes           []interface{} `json:"quotes"`
	Crossquotes      []interface{} `json:"crossquotes"`
	Backlinks        []Backlinks   `json:"backlinks"`
	Replyposts       int           `json:"replyposts"`
	Replyfiles       int           `json:"replyfiles"`
	Sticky           int           `json:"sticky"`
	Locked           int           `json:"locked"`
	Bumplocked       int           `json:"bumplocked"`
	Cyclic           int           `json:"cyclic"`
	Bumped           ScuffedTime   `json:"bumped,omitempty"`
	PostID           int           `json:"postId"`
	Replies          []Post        `json:"replies,omitempty"`
	Previewbacklinks []Backlinks   `json:"previewbacklinks,omitempty"`
	Omittedfiles     int           `json:"omittedfiles,omitempty"`
	Omittedposts     int           `json:"omittedposts,omitempty"`
	Edited           Edited        `json:"edited,omitempty"`
	Reports          []Report      `json:"reports"`
	GlobalReports    []Report      `json:"globalreports"`
	IP               IP            `json:"ip,omitempty"`
}

type Report struct {
	ID     string      `json:"id"`
	Reason string      `json:"reason"`
	Date   ScuffedTime `json:"date"`
	IP     IP          `json:"ip"`
}

type IP struct {
	Raw    string `json:"raw,omitempty"`
	Cloak  string `json:"cloak"`
	Pruned bool   `json:"pruned,omitempty"`
	Type   int    `json:"type,omitempty"`
}

type Files struct {
	Spoiler          interface{} `json:"spoiler"`
	Hash             string      `json:"hash"`
	Filename         string      `json:"filename"`
	OriginalFilename string      `json:"originalFilename"`
	Mimetype         string      `json:"mimetype"`
	Size             int         `json:"size"`
	Extension        string      `json:"extension"`
	Thumbextension   string      `json:"thumbextension"`
	Geometry         Geometry    `json:"geometry"`
	GeometryString   string      `json:"geometryString"`
	HasThumb         bool        `json:"hasThumb"`
	SizeString       string      `json:"sizeString"`
}

type Geometry struct {
	Width       int `json:"width"`
	Height      int `json:"height"`
	Thumbwidth  int `json:"thumbwidth"`
	Thumbheight int `json:"thumbheight"`
}

type Country struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Src    string `json:"src,omitempty"`
	Custom bool   `json:"custom"`
}

type Quotes struct {
	ID     string `json:"_id"`
	Thread int    `json:"thread"`
	PostID int    `json:"postId"`
}

type Backlinks struct {
	ID     string `json:"_id"`
	PostID int    `json:"postId"`
}

type Edited struct {
	Username string      `json:"username"`
	Date     ScuffedTime `json:"date"`
}
