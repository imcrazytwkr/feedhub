package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type CursorProvider interface {
	GetBlog(ctx context.Context, topic string, lang models.Language) (*models.Feed, error)
}
