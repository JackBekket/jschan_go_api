package jschan

import (
	"context"
	"fmt"
	"jschan/app/models"
	"net/http"
	"net/url"
	"strings"
)

type GetOverboardOptions struct {
	AddBoards      []string
	RemoveBoards   []string
	IncludeDefault bool
}

type GetOverboardResponse struct {
	Threads []models.Post `json:"threads"`
}

func getOverboardQuery(options *GetOverboardOptions) (*url.Values, error) {
	include_default := false
	add := []string{}
	rem := []string{}
	if options != nil {
		add = options.AddBoards
		rem = options.RemoveBoards
	}
	query := url.Values{}
	query.Set("add", strings.Join(add, ","))
	query.Set("rem", strings.Join(rem, ","))
	if include_default {
		query.Set("include_default", "true")
	}
	return &query, nil
}

func (c *Client) GetOverboardIndex(ctx context.Context, options *GetOverboardOptions) (*GetOverboardResponse, error) {

	query, err := getOverboardQuery(options)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/overboard.json?%s", c.BaseURL, query.Encode())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetOverboardResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) GetOverboardCatalog(ctx context.Context, options *GetOverboardOptions) (*GetOverboardResponse, error) {

	query, err := getOverboardQuery(options)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/catalog.json?%s", c.BaseURL, query.Encode())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetOverboardResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}
