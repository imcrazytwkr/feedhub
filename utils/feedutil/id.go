package feedutil

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

// @TODO: replace with `time.DateOnly` when stable Go hits 1.20.x branch
const dateFormat = "2006-01-02"

func GenerateId(entry *models.Entry) string {
	if len(entry.Id) > 0 {
		return entry.Id
	}

	if len(entry.Link) > 0 {
		return generateGuid(entry)
	}

	return generateUuid()
}

func generateGuid(entry *models.Entry) string {
	dateString := timeutil.FormatMaxOfTwo(dateFormat, entry.Updated, entry.Published)
	if len(dateString) == 0 {
		return entry.Link
	}

	url, err := url.Parse(entry.Link)
	if err != nil {
		return generateUuid()
	}

	return fmt.Sprintf("tag:%s,%s:%s", url.Host, dateString, url.Path)
}

func generateUuid() string {
	return "urn:uuid:" + uuid.NewString()
}
