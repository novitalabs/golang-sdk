package request

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/novitalabs/golang-sdk/types"
	"github.com/novitalabs/golang-sdk/util"
)

func TestClient_SyncTxt2Img(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"))
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	txt2ImgReq := types.NewTxt2ImgRequest("a dog flying in the sky", "", "AnythingV5_v5PrtRE.safetensors")
	res, err := client.SyncTxt2img(ctx, txt2ImgReq,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_txt2img_sync.png"
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)
}

func TestClient_SyncTxt2ImgWithLora(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"))
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	txt2ImgReq := types.NewTxt2ImgRequest("a dog flying in the sky, <lora:add_detail_44319:1>", "", "AnythingV5_v5PrtRE.safetensors")
	res, err := client.SyncTxt2img(ctx, txt2ImgReq,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			return "test_txt2img_sync.png"
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)
}

func TestClient_SyncTxt2ImgControlNet(t *testing.T) {
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
	txt2ImgReq := types.NewTxt2ImgRequest("a dog flying in the sky", "", "")
	controlNetReq := types.NewControlNetUnit(types.Canny, "control_v11p_sd15_canny", initImageBase64)
	txt2ImgReq.ControlNetUnits = []*types.ControlNetUnit{controlNetReq}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	res, err := client.SyncTxt2img(ctx, txt2ImgReq,
		WithSaveImage("out", 0777, func(taskId string, fileIndex int, fileName string) string {
			if fileIndex == 0 {
				return "test_txt2img_controlnet_sync.png"
			} else {
				return "test_txt2img_controlnet_processor.png"
			}
		}))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status = %d", res.Data.Status)

}
