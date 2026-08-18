package feedutil

import (
	"slices"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

func SortEntries(entries []*models.Entry) []*models.Entry {
	if len(entries) < 2 {
		return entries
	}

	slices.SortFunc(entries, func(a, b *models.Entry) int {
		timeA := timeutil.MaxOfTwo(a.Updated, a.Published)
		timeB := timeutil.MaxOfTwo(b.Updated, b.Published)
		// Reverse sort
		return timeB.Compare(timeA)
	})

	return entries
}
