package jschan

import (
	"context"
	"fmt"
	"jschan/app/models"
	"net/http"
	"net/url"
)

type GetWebringResponse struct {
	*models.Webring
}

func (c *Client) GetWebring(ctx context.Context) (*GetWebringResponse, error) {

	url := fmt.Sprintf("%s/webring.json", c.BaseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetWebringResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}

type GetBoardsResponse struct {
	Boards  []models.Board `json:"boards"`
	Page    int            `json:"page"`
	MaxPage int            `json:"maxPage"`
}

type GetBoardsPublicOptions struct {
	Search        string
	Sort          string
	SortDirection string
	Page          int
	LocalFirst    bool
	Sites         []string
}

func (c *Client) GetBoardsPublic(ctx context.Context, options *GetBoardsPublicOptions) (*GetBoardsResponse, error) {

	page := 1
	search := ""
	sort := "popularity"
	direction := "desc"
	local_first := false
	sites := []string{}
	if options != nil {
		search = options.Search
		sort = options.Sort
		direction = options.SortDirection
		local_first = options.LocalFirst
		sites = options.Sites
		page = options.Page
	}

	query := url.Values{}
	query.Set("search", search)
	query.Set("page", fmt.Sprintf("%d", page))
	query.Set("sort", sort)
	query.Set("direction", direction)
	if local_first {
		query.Set("local_first", "true")
	}
	for _, site := range sites {
		query.Add("sites", site)
	}

	url := fmt.Sprintf("%s/boards.json?%s", c.BaseURL, query.Encode())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetBoardsResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}
