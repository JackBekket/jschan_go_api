package jschan

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackbekket/jschan_go_api/app/models"
)

type GetManageRecentOptions struct {
	Board string
}

type GetManageRecentResponse []models.Post

func (c *Client) GetManageRecent(ctx context.Context, options *GetManageRecentOptions) (GetManageRecentResponse, error) {

	url := "/globalmanage/recent.json"
	if options != nil && options.Board != "" {
		url = fmt.Sprintf("%s/%s/manage/recent.json", c.BaseURL, options.Board)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetManageRecentResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return res, nil

}

type GetManageReportsOptions struct {
	Page  int
	IP    string
	Board string
}

type GetManageReportsResponse struct {
	Reports     []models.Post `json:"reports"`
	Page        int           `json:"page"`
	IP          string        `json:"ip"`
	QueryString string        `json:"queryString"`
}

func (c *Client) GetManageReports(ctx context.Context, options *GetManageReportsOptions) (*GetManageReportsResponse, error) {

	query := url.Values{}
	if options != nil {
		if options.IP != "" {
			query.Set("ip", options.IP)
		}
		if options.Page != 0 {
			query.Set("page", strconv.Itoa(options.Page))
		}
	}

	url := fmt.Sprintf("%s/globalmanage/reports.json", c.BaseURL)
	if options != nil && options.Board != "" {
		url = fmt.Sprintf("%s/%s/manage/reports.json", c.BaseURL, options.Board)
	}
	if len(query.Encode()) > 0 {
		url = fmt.Sprintf("%s?%s", url, query.Encode())
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetManageReportsResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}

type GetManageBoardsOptions struct {
	Search          string `json:"search"`
	Sort            string `json:"sort"`
	SortDirection   string `json:"direction"`
	Page            int    `json:"page"`
	FilterUnlisted  bool   `json:"filter_unlisted"`
	FilterSfw       bool   `json:"filter_sfw"`
	FilterAbandoned bool   `json:"filter_abandoned"`
}

func (c *Client) GetManageBoards(ctx context.Context, options *GetManageBoardsOptions) (*GetBoardsResponse, error) {

	page := 1
	search := ""
	sort := "popularity"
	direction := "desc"
	filter_unlisted := false
	filter_sfw := false
	filter_abandoned := false
	if options != nil {
		search = options.Search
		sort = options.Sort
		direction = options.SortDirection
		filter_unlisted = options.FilterUnlisted
		filter_sfw = options.FilterSfw
		filter_abandoned = options.FilterAbandoned
		page = options.Page
	}

	query := url.Values{}
	query.Set("search", search)
	query.Set("page", fmt.Sprintf("%d", page))
	query.Set("sort", sort)
	query.Set("direction", direction)
	if filter_unlisted {
		query.Set("filter_unlisted", "true")
	}
	if filter_sfw {
		query.Set("filter_sfw", "true")
	}
	if filter_abandoned {
		query.Set("filter_abandoned", "true")
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
