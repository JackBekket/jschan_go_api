package jschan

import (
	"context"
	"fmt"
	"jschan/app/models"
	"net/http"
)

type GetGlobalSettingsResponse struct {
	*models.GlobalSettings
}

func (c *Client) GetGlobalSettings(ctx context.Context) (*CatalogResponse, error) {

	url := fmt.Sprintf("%s/settings.json", c.BaseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := CatalogResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}
