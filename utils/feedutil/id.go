package feedutil

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/timeutil"
)

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
	dateString := timeutil.FormatMaxOfTwo(time.DateOnly, entry.Updated, entry.Published)
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
