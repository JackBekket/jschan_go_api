package models

type BoardSettings struct {
	CustomPages            []string              `json:"customPages"`
	Announcement           RawAndMarkdownMessage `json:"announcement"`
	AllowedFileTypes       AllowedFileTypes      `json:"allowedFileTypes"`
	MaxFiles               int                   `json:"maxFiles"`
	CaptchaMode            int                   `json:"captchaMode"`
	ForceAnon              bool                  `json:"forceAnon"`
	SageOnlyEmail          bool                  `json:"sageOnlyEmail"`
	CustomFlags            bool                  `json:"customFlags"`
	ForceThreadMessage     bool                  `json:"forceThreadMessage"`
	ForceThreadFile        bool                  `json:"forceThreadFile"`
	ForceThreadSubject     bool                  `json:"forceThreadSubject"`
	DisableReplySubject    bool                  `json:"disableReplySubject"`
	MinThreadMessageLength int                   `json:"minThreadMessageLength"`
	MinReplyMessageLength  int                   `json:"minReplyMessageLength"`
	MaxThreadMessageLength int                   `json:"maxThreadMessageLength"`
	MaxReplyMessageLength  int                   `json:"maxReplyMessageLength"`
	DefaultName            string                `json:"defaultName"`
}

type AllowedFileTypes struct {
	AnimatedImage bool `json:"animatedImage"`
	Image         bool `json:"image"`
	Video         bool `json:"video"`
	Audio         bool `json:"audio"`
	Other         bool `json:"other"`
}

type GlobalSettings struct {
	CaptchaOptions CaptchaOptions `json:"captchaOptions"`
}

type CaptchaOptions struct {
	Type string `json:"type"`
	Grid Grid   `json:"grid,omitempty"`
}

type Grid struct {
	Size     int    `json:"size"`
	Question string `json:"question"`
}
