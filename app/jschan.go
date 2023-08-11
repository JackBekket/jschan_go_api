package jschan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	BaseURL       string
	SessionCookie string
	CsrfToken     string
	HTTPClient    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:       baseURL,
		SessionCookie: "",
		CsrfToken:     "",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

type DynamicResponse struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Redirect string `json:"redirect,omitempty"`
}

type ReturnHeaders struct {
	*http.Header
}

func cookieHeader(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}

func (c *Client) sendRequest(req *http.Request, v interface{}, h *ReturnHeaders) error {

	if req.Header.Get("Content-Type") == "" {
		if req.Method == http.MethodGet {
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		} else {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Using-XHR", "true")
	req.Header.Set("Referer", c.BaseURL)
	if c.SessionCookie != "" {
		req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.SessionCookie))
	}
	if c.CsrfToken != "" {
		req.Header.Set("Csrf-Token", c.CsrfToken)
	}

	fmt.Printf("%s %s\n", req.Method, req.URL)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes DynamicResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return err
		}
		return errors.New(errRes.Message)
	}

	h = &ReturnHeaders{
		&res.Header,
	}
	setCookieValue := h.Get("Set-Cookie")
	if setCookieValue != "" {
		parsedSetCookie := cookieHeader(setCookieValue)
		for _, parsedCookie := range parsedSetCookie {
			if parsedCookie.Name == "connect.sid" {
				c.SessionCookie = parsedCookie.Value
			}
		}
	}

	if v != nil {
		fullResponse := v
		err := json.NewDecoder(res.Body).Decode(&fullResponse)
		if err != nil {
			body, err2 := ioutil.ReadAll(res.Body)
			if err2 != nil {
				return err
			}
			err3 := json.Unmarshal(body, &fullResponse)
			if err3 != nil {
				return err
			}
		}
	}

	return nil
}
