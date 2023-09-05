package jschan

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackbekket/jschan_go_api/app/models"
)

type GetCatalogOptions struct {
	Board string
}

type CatalogResponse []models.Post

func (c *Client) GetCatalog(ctx context.Context, options *GetCatalogOptions) (CatalogResponse, error) {

	url := fmt.Sprintf("%s/%s/catalog.json", c.BaseURL, options.Board)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := CatalogResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return res, nil

}

type GetIndexOptions struct {
	Board string
	Page  int
}

type GetIndexResponse struct {
	*models.Post
}

func (c *Client) GetIndex(ctx context.Context, options *GetIndexOptions) (*[]GetIndexResponse, error) {

	pageString := strconv.Itoa(options.Page)
	if pageString == "1" {
		pageString = "index"
	}

	url := fmt.Sprintf("%s/%s/%s.json", c.BaseURL, options.Board, pageString)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := []GetIndexResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}

type GetLogsListOptions struct {
	Board string
}

type GetLogsListResponse []models.LogList

func (c *Client) GetLogsList(ctx context.Context, options *GetLogsListOptions) (GetLogsListResponse, error) {

	url := fmt.Sprintf("%s/%s/logs.json", c.BaseURL, options.Board)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetLogsListResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return res, nil

}

type GetLogsOptions struct {
	Board string
	Date  models.LogDate
}

type GetLogsResponse []models.Log

func (c *Client) GetLogs(ctx context.Context, options *GetLogsOptions) (GetLogsResponse, error) {

	url := fmt.Sprintf("%s/%s/logs/%s.json", c.BaseURL, options.Board, options.Date.String())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetLogsResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return res, nil

}

type GetCustomPageOptions struct {
	Board      string
	CustomPage string
}

type GetCustomPageResponse struct {
	*models.CustomPage
}

func (c *Client) GetCustomPage(ctx context.Context, options *GetCustomPageOptions) (*GetCustomPageResponse, error) {

	url := fmt.Sprintf("%s/%s/custompage/%s.json", c.BaseURL, options.Board, options.CustomPage)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetCustomPageResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}

type GetBannersOptions struct {
	Board string
}

type GetBannersResponse []string

func (c *Client) GetBanners(ctx context.Context, options *GetBannersOptions) (GetBannersResponse, error) {

	url := fmt.Sprintf("%s/%s/banners.json", c.BaseURL, options.Board)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetBannersResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return res, nil

}

type GetSettingsOptions struct {
	Board string
}

type GetSettingsResponse struct {
	*models.BoardSettings
}

func (c *Client) GetSettings(ctx context.Context, options *GetSettingsOptions) (*GetSettingsResponse, error) {

	url := fmt.Sprintf("%s/%s/settings.json", c.BaseURL, options.Board)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetSettingsResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}
