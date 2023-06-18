package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type ArknightsProvider interface {
	GetNews(ctx context.Context, language string) (*models.Feed, error)
}
