package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type BleepingComputerProvider interface {
	GetNews(ctx context.Context) (*models.Feed, error)
}
