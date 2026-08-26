package mappers

import (
	"net/url"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/imcrazytwkr/feedhub/constants"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/feedutil"
	"golang.org/x/net/html"
)

func PluckBlogEntries(root *html.Node) []*models.Entry {
	main := lastMain(root)
	if main == nil {
		return nil
	}

	rows := htmlquery.Find(main, ".//a[contains(@class,'blog-directory__row')]")
	if len(rows) == 0 {
		return nil
	}

	entries := make([]*models.Entry, len(rows))
	i := 0
	for _, row := range rows {
		entry := mapDirectoryRow(row)
		if entry == nil {
			continue
		}

		entries[i] = entry
		i++
	}

	if i == 0 {
		return nil
	}

	if i < len(entries) {
		entries = entries[:i]
	}

	entries = feedutil.SortEntries(entries)
	if len(entries) > entryLimit {
		return entries[:entryLimit]
	}

	return entries
}

func PluckBlogDescription(root *html.Node) string {
	node := htmlquery.FindOne(root, "//head/meta[@property='og:description']")
	if node == nil {
		return ""
	}

	return strings.TrimSpace(htmlquery.SelectAttr(node, "content"))
}

func lastMain(root *html.Node) *html.Node {
	mains := htmlquery.Find(root, "//*[@id='main']")

	size := len(mains)
	switch size {
	case 0:
		return nil
	case 1:
		return mains[0]
	default:
		return mains[size-1]
	}
}

func mapDirectoryRow(row *html.Node) *models.Entry {
	link := absolutize(htmlquery.SelectAttr(row, "href"))
	if len(link) == 0 {
		return nil
	}

	titleNode := htmlquery.FindOne(row, ".//p")
	if titleNode == nil {
		return nil
	}

	title := strings.TrimSpace(htmlquery.InnerText(titleNode))
	if len(title) == 0 {
		return nil
	}

	entry := models.Entry{
		Title: title,
		Link:  link,
	}

	timeNode := htmlquery.FindOne(row, ".//time")
	if timeNode != nil {
		entry.Published = parseTimestamp(htmlquery.SelectAttr(timeNode, "datetime"))
	}

	return &entry
}

func absolutize(href string) string {
	href = strings.TrimSpace(href)
	if len(href) == 0 {
		return ""
	}

	// Shorthand covering most, if not all, URLs
	if strings.HasPrefix(href, blogPath) {
		return host + href
	}

	// Fallback option in case we get a non-standard link
	parsed, err := url.Parse(href)
	if err != nil {
		return ""
	}

	if parsed.IsAbs() {
		return parsed.String()
	}

	return host + href
}

func parseTimestamp(value string) time.Time {
	text := strings.TrimSpace(value)
	if len(text) == 0 {
		return constants.TimeZero
	}

	timestamp, err := time.Parse(time.RFC3339, text)
	if err != nil {
		return constants.TimeZero
	}

	return timestamp
}
