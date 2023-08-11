package jschan

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strconv"
)

type MakePostOptions struct {
	Board         string
	Thread        int
	Name          string
	Message       string
	Subject       string
	Email         string
	PostPassword  string
	Files         []string
	Spoiler       []string
	SpoilerAll    bool
	StripFilename []string
	CustomFlag    string
	Captcha       []string
	Mod           bool
}

func (c *Client) MakePost(ctx context.Context, options *MakePostOptions) error {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if options.Files != nil && len(options.Files) > 0 {
		for _, filepath := range options.Files {
			dir, fileName := path.Split(filepath)
			filePath := path.Join(dir, fileName)
			file, _ := os.Open(filePath)
			defer file.Close()
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
			h.Set("Content-Type", "image/png")
			part, _ := writer.CreatePart(h)
			//part, _ := writer.CreateFormFile("file", h)
			io.Copy(part, file)
		}
	}
	_ = writer.WriteField("thread", strconv.Itoa(options.Thread))
	_ = writer.WriteField("name", options.Name)
	_ = writer.WriteField("message", options.Message)
	_ = writer.WriteField("subject", options.Subject)
	_ = writer.WriteField("email", options.Email)
	_ = writer.WriteField("postpassword", options.PostPassword)
	_ = writer.WriteField("customflag", options.CustomFlag)
	if options.SpoilerAll == true {
		_ = writer.WriteField("spoiler_all", "true")
	}
	for _, filename := range options.Spoiler {
		_ = writer.WriteField("spoiler", filename)
	}
	for _, filename := range options.StripFilename {
		_ = writer.WriteField("strip_filename", filename)
	}
	for _, answer := range options.Captcha {
		_ = writer.WriteField("captcha", answer)
	}
	writer.Close()

	suffix := "post"
	if options.Mod {
		suffix = "modpost"
	}
	url := fmt.Sprintf("%s/forms/board/%s/%s", c.BaseURL, options.Board, suffix)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	//	req.Header.Set("content-type", "multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req = req.WithContext(ctx)

	if err := c.sendRequest(req, nil, nil); err != nil {
		return err
	}

	return nil

}

type SubmitPostActionsOptions struct {
	CheckedPosts      []string
	CheckedReports    []string
	Board             string
	Captcha           []string
	PostPassword      string
	Ban               bool
	BanGlobal         bool
	BanHalfRange      bool
	BanQuarterRance   bool
	BanReporter       bool
	BanReporterGlobal bool
	BanReason         string
	BanDuration       string
	Move              bool
	MoveToThread      int
	Sticky            int
	ToggleCyclic      bool
	ToggleBumplock    bool
	ToggleLock        bool
	Spoiler           bool
	Delete            bool
	DeleteFiles       bool
	UnlinkFiles       bool
	DeleteIPThread    bool
	DeleteIPBoard     bool
	DeleteIPGlobal    bool
	Dismiss           bool
	DismissGlobal     bool
	Report            bool
	ReportGlobal      bool
	ReportReason      string
	HideName          bool
	NoAppeal          bool
	PreservePost      bool
	LogMessage        string
}

func (c *Client) SubmitPostActions(ctx context.Context, options *MakePostOptions) error {

	return nil

}
