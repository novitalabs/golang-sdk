package request

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/novitalabs/golang-sdk/types"
)

func TestClient_Models(t *testing.T) {
	client, err := NewClient(os.Getenv("API_KEY"))
	if err != nil {
		t.Error(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	models, err := client.Models(ctx, WithRefresh())
	if err != nil {
		t.Error(err)
		return
	}
	// test filtering and sorting
	t.Log(models)
	top := models.FilterType(types.Checkpoint).TopN(10, func(m *types.Model) float32 {
		return float32(m.CivitaiDownloadCount)
	})
	t.Log(top)
}
