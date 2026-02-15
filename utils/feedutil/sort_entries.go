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
		entryA := entries[i]
		timeA := timeutil.MaxOfTwo(entryA.Updated, entryA.Published)

		entryB := entries[j]
		timeB := timeutil.MaxOfTwo(entryB.Updated, entryB.Published)

		// Reverse sort
		return timeB.Before(timeA)
	})

	return entries
}
