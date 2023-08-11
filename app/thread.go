package jschan

import (
	"context"
	"fmt"
	"jschan/app/models"
	"net/http"
)

type GetThreadOptions struct {
	Board    string
	ThreadId int
}

type GetThreadResponse struct {
	*models.Post
}

func (c *Client) GetThread(ctx context.Context, options *GetThreadOptions) (*GetThreadResponse, error) {

	url := fmt.Sprintf("%s/%s/thread/%d.json", c.BaseURL, options.Board, options.ThreadId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := GetThreadResponse{}
	if err := c.sendRequest(req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil

}
