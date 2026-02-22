package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type TomRigbyProvider interface {
	GetPosts(ctx context.Context) (*models.Feed, error)
}
