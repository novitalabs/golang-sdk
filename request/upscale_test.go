package request

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/novitalabs/golang-sdk/types"
	"github.com/novitalabs/golang-sdk/util"
)

func TestClient_SyncUpscale(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"))
	if err != nil {
		t.Error(err)
		return
	}
	initImage := "out/test_txt2img_sync.png"
	initImageBase64, err := util.ReadImageToBase64(initImage)
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	upscale := types.NewUpscaleRequest(initImageBase64, 2)
	res, err := client.SyncUpscale(ctx, upscale,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_upscale_sync.png"
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)
}
