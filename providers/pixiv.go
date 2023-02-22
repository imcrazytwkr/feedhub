package providers

import (
	"context"

	"github.com/imcrazytwkr/feedhub/models"
)

type PixivProvider interface {
	GetUserIllustrations(ctx context.Context, userId int) (*models.Feed, error)
}
