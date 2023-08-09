package request

import (
	"context"
	"net/http"

	"github.com/novitalabs/golang-sdk/types"
)

func (c *Client) Upscale(ctx context.Context, request *types.UpscaleRequest) (*types.AsyncResponse, error) {
	responseData, err := doRequest[types.UpscaleRequest, types.AsyncResponse](ctx, c.httpCli, http.MethodPost, BaseURL+"/upscale", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *Client) SyncUpscale(ctx context.Context, request *types.UpscaleRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	return syncImageGeneration[*types.UpscaleRequest](ctx, request, opts, c.Upscale, c.waitForTask)
}
