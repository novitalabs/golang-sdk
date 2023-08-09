package request

import (
	"context"
	"net/http"

	"github.com/novitalabs/golang-sdk/types"
)

func (c *Client) Progress(ctx context.Context, request *types.ProgressRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	responseData, err := doRequest[types.ProgressRequest, types.ProgressResponse](ctx, c.httpCli, http.MethodGet, BaseURL+"/progress", c.apiKey, map[string]interface{}{
		"task_id": request.TaskId,
	}, nil)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
