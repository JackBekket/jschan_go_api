package models

type Board struct {
	ID                string            `json:"_id"`
	URI               string            `json:"uri,omitempty"`
	Path              string            `json:"path,omitempty"`
	LastPostTimestamp LastPostTimestamp `json:"lastPostTimestamp"`
	Tags              []interface{}     `json:"tags,omitempty"`
	SiteName          string            `json:"siteName,omitempty"`
	SequenceValue     int               `json:"sequence_value"`
	Pph               int               `json:"pph"`
	Ips               int               `json:"ips"`
	Settings          Settings          `json:"settings,omitempty"`
	Webring           bool              `json:"webring"`
	Ppd               int               `json:"ppd,omitempty"`
}

type Settings struct {
	Sfw           bool   `json:"sfw"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	UnlistedLocal bool   `json:"unlistedLocal,omitempty"`
}

type LastPostTimestamp struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
