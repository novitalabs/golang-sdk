package request

import (
	"context"
	"net/http"

	"github.com/novitalabs/golang-sdk/types"
)

func (c *Client) Img2Img(ctx context.Context, request *types.Img2ImgRequest) (*types.AsyncResponse, error) {
	responseData, err := doRequest[types.Img2ImgRequest, types.AsyncResponse](ctx, c.httpCli, http.MethodPost, BaseURL+"/img2img", c.apiKey, nil, request)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (c *Client) SyncImg2img(ctx context.Context, request *types.Img2ImgRequest, opts ...WithGenerateImageOption) (*types.ProgressResponse, error) {
	return syncImageGeneration[*types.Img2ImgRequest](ctx, request, opts, c.Img2Img, c.waitForTask)
}
