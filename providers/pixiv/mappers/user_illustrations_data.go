package mappers

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/timeutil"
	"github.com/valyala/fastjson"
)

func PluckIllustrationEntries(contents *fastjson.Value) ([]*models.Entry, error) {
	if contents == nil {
		return nil, nil
	}

	err, hasError := processErrorFields(contents)
	if hasError {
		return nil, err
	}

	payload := contents.GetObject(bodyKey, worksKey)
	if payload == nil || payload.Len() == 0 {
		return nil, nil
	}

	targetLength := payload.Len()
	illustrations := make([]*models.Entry, targetLength)
	i := 0

	payload.Visit(func(key []byte, v *fastjson.Value) {
		entry := parseImageEntry(key, v)
		if entry != nil {
			illustrations[i] = entry
			i++
		}
	})

	if i == 0 {
		return nil, nil
	}

	if i < targetLength {
		illustrations = illustrations[0:i]
	}

	sort.Slice(illustrations, func(i, j int) bool {
		illustA := illustrations[i]
		timeA := timeutil.MaxOfTwo(illustA.Updated, illustA.Published)

		illustB := illustrations[j]
		timeB := timeutil.MaxOfTwo(illustB.Updated, illustB.Published)

		return timeA.Before(timeB)
	})

	if i == targetLength {
		return illustrations, nil
	}

	return illustrations[0:i], nil
}

func parseImageEntry(key []byte, v *fastjson.Value) *models.Entry {
	if len(key) == 0 {
		return nil
	}

	link := postPrefix + string(key)

	title := v.GetStringBytes(titleKey)
	if len(title) == 0 {
		return nil
	}

	author := v.GetStringBytes(userNameKey)
	if len(author) == 0 {
		return nil
	}

	publicationDate, _ := time.Parse(time.RFC3339, string(v.GetStringBytes(createDateKey)))
	updatedDate, _ := time.Parse(time.RFC3339, string(v.GetStringBytes(updateDateKey)))
	if publicationDate.IsZero() && updatedDate.IsZero() {
		return nil
	}

	description := parseDescription(v, link)
	if len(description) == 0 {
		return nil
	}

	return &models.Entry{
		Title:     string(title),
		Published: publicationDate,
		Updated:   updatedDate,
		Author:    string(author),
		Link:      link,
		Content:   description,
	}
}

func parseDescription(v *fastjson.Value, link string) string {
	pageCount := v.GetInt(pageCountKey)
	if pageCount == 0 {
		return ""
	}

	previewUrl := v.GetStringBytes(urlKey)
	if len(previewUrl) == 0 {
		return ""
	}

	// m[1] - post prefix, m[2] - extension
	match := imageRe.FindSubmatch(previewUrl)
	if len(match) == 0 || len(match[1]) == 0 || len(match[2]) == 0 {
		return ""
	}

	description := strings.Builder{}

	// Post images
	for page := 0; page < pageCount; page++ {
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
	sourceDescription := v.GetStringBytes(descriptionKey)
	if len(sourceDescription) > 0 {
		fmt.Fprintf(&description, `<p>%s</p>`, sourceDescription)
	}

	description.WriteString(`<p>`)

	// Post link
	fmt.Fprintf(&description, `[<a href="%s">link</a>]`, link)

	// Artist link
	artistId := v.GetStringBytes(userIdKey)
	if len(artistId) > 0 {
		fmt.Fprintf(&description, ` [<a href="%s%s">artist</a>]`, artistPrefix, artistId)
	}

	// End links
	description.WriteString(`</p>`)

	return description.String()
}
