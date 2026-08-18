package mappers

import (
	"fmt"
	"strings"
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
)

func PluckIllustrationEntries(contents *m.Response[m.IllustrationDataBody]) ([]*models.Entry, error) {
	if contents == nil {
		return nil, nil
	}

	err, hasError := processErrorFields(&contents.ApiError)
	if hasError {
		return nil, err
	}

	works := contents.Body.Works
	targetLength := len(works)

	if targetLength == 0 {
		return nil, nil
	}

	illustrations := make([]*models.Entry, targetLength)
	i := 0

	for key, work := range works {
		entry := parseImageEntry(key, &work)
		if entry != nil {
			illustrations[i] = entry
			i++
		}
	}

	if i == 0 {
		return nil, nil
	}

	if i < targetLength {
		illustrations = illustrations[0:i]
	}

	return feedutil.SortEntries(illustrations), nil
}

func parseImageEntry(key string, work *m.IllustrationWork) *models.Entry {
	if len(key) == 0 {
		return nil
	}

	if len(work.Title) == 0 {
		return nil
	}

	if len(work.UserName) == 0 {
		return nil
	}

	publicationDate, _ := time.Parse(time.RFC3339, work.CreateDate)
	updatedDate, _ := time.Parse(time.RFC3339, work.UpdateDate)
	if publicationDate.IsZero() && updatedDate.IsZero() {
		return nil
	}

	link := postPrefix + key

	description := parseDescription(work, link)
	if len(description) == 0 {
		return nil
	}

	return &models.Entry{
		Title:     work.Title,
		Published: publicationDate,
		Updated:   updatedDate,
		Author:    work.UserName,
		Link:      link,
		Content:   description,
	}
}

func parseDescription(work *m.IllustrationWork, link string) string {
	if work.PageCount == 0 {
		return ""
	}

	if len(work.URL) == 0 {
		return ""
	}

	// m[1] - post prefix, m[2] - extension
	match := imageRe.FindStringSubmatch(work.URL)
	if len(match) == 0 || len(match[1]) == 0 || len(match[2]) == 0 {
		return ""
	}

	description := strings.Builder{}

	// Post images
	for page := 0; page < work.PageCount; page++ {
		fmt.Fprintf(
			&description,
			`<p><img src="%s/img-master/img%s_p%d_square1200.%s" alt="page %d cover" /></p>`,
			cdnPrefix,
			match[1],
			page,
			match[2],
			page,
		)
	}

	// Original description, if any
	if len(work.Description) > 0 {
		fmt.Fprintf(&description, `<p>%s</p>`, work.Description)
	}

	description.WriteString(`<p>`)

	// Post link
	fmt.Fprintf(&description, `[<a href="%s">link</a>]`, link)

	// Artist link
	if len(work.UserID) > 0 {
		fmt.Fprintf(&description, ` [<a href="%s%s">artist</a>]`, artistPrefix, work.UserID)
	}

	// End links
	description.WriteString(`</p>`)

	return description.String()
}
