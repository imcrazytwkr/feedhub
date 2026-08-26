package http

import (
	"net/http"

	"github.com/imcrazytwkr/feedhub/constants"
)

var headers = http.Header{
	constants.UserAgent:            {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/113.0"},
	constants.AcceptHeader:         {"text/html;q=0.9,*/*;q=0.8"},
	constants.AcceptLanguageHeader: {"en-US,en;q=0.5"},
}

const host = "https://cursor.com"
const blogPath = "/blog"

// Covers a listing plus article overlap for concurrent requests.
const maxCacheEntries = 256
