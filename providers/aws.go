package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type AWSDirectoryProvider interface {
	GetBlogs(ctx context.Context, category string, language models.Language) (*models.Feed, error)
}
