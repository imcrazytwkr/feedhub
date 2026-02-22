package feedutil

import (
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

func GetLastPublishedTime(entries []*models.Entry) time.Time {
	lastPublished := time.Time{}
	for _, entry := range entries {
		lastPublished = timeutil.MaxOfTwo(lastPublished, entry.Published)
	}

	return lastPublished
}

func GetLastUpdatedTime(entries []*models.Entry) time.Time {
	lastUpdated := time.Time{}
	for _, entry := range entries {
		lastUpdated = timeutil.MaxOfTwo(lastUpdated, entry.Updated)
	}

	return lastUpdated
}
