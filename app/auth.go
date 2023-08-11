package jschan

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type PostLoginOptions struct {
	Username  string
	Password  string
	Twofactor string
}

func (c *Client) Login(ctx context.Context, options *PostLoginOptions) error {

	formData := url.Values{}
	formData.Set("username", options.Username)
	formData.Set("password", options.Password)
	formData.Set("twofactor", options.Twofactor)
	endodedBody := strings.NewReader(formData.Encode())

	url := fmt.Sprintf("%s/forms/login", c.BaseURL)

	req, err := http.NewRequest(http.MethodPost, url, endodedBody)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, nil, nil); err != nil {
		return err
	}

	return nil

}

type GetCSRFTokenResponse struct {
	Token string `json:"token"`
}

func (c *Client) GetCSRFToken(ctx context.Context) (*GetCSRFTokenResponse, error) {

	url := fmt.Sprintf("%s/csrf.json", c.BaseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := &GetCSRFTokenResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	c.CsrfToken = res.Token

	return res, nil

}
