package models

type Webring struct {
	Name      string         `json:"name"`
	URL       string         `json:"url"`
	Endpoint  string         `json:"endpoint"`
	Logo      []string       `json:"logo"`
	Following []string       `json:"following"`
	Blacklist []string       `json:"blacklist"`
	Known     []string       `json:"known"`
	Boards    []WebringBoard `json:"boards"`
}

type WebringBoard struct {
	URI               string      `json:"uri"`
	Title             string      `json:"title"`
	Subtitle          string      `json:"subtitle"`
	Path              string      `json:"path"`
	PostsPerHour      int         `json:"postsPerHour"`
	PostsPerDay       int         `json:"postsPerDay"`
	TotalPosts        int         `json:"totalPosts"`
	UniqueUsers       int         `json:"uniqueUsers"`
	Nsfw              bool        `json:"nsfw"`
	Tags              []string    `json:"tags"`
	LastPostTimestamp ScuffedTime `json:"lastPostTimestamp"`
}
