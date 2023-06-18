package feedutil

import (
	"sort"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

func SortEntries(entries []*models.Entry) []*models.Entry {
	if len(entries) < 2 {
		return entries
	}

	sort.Slice(entries, func(i, j int) bool {
		illustA := entries[i]
		timeA := timeutil.MaxOfTwo(illustA.Updated, illustA.Published)

		illustB := entries[j]
		timeB := timeutil.MaxOfTwo(illustB.Updated, illustB.Published)

		// Reverse sort
		return timeB.Before(timeA)
	})

	return entries
}
