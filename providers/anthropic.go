package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type AnthropicProvider interface {
	GetNews(ctx context.Context) (*models.Feed, error)
	GetEngineering(ctx context.Context) (*models.Feed, error)
	GetResearch(ctx context.Context, team string) (*models.Feed, error)
}
