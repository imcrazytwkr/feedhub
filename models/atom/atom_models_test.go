package atom_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/models/atom"
)

const testXml = `<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en"><generator>FeedHub/Atom</generator><id>test_link</id><title>test_title</title><subtitle>test_subtitle</subtitle><published>2023-02-18T16:28:00Z</published><updated>2023-02-18T16:28:00Z</updated><link href="test_self_link" rel="self" type="application/atom+xml"></link><link href="test_link"></link><author><name>test_author</name></author></feed>`

func Test(t *testing.T) {
	feed := atom.NewAtomFeed(&models.Feed{
		Title:       "test_title",
		Description: "test_subtitle",
		Updated:     time.Date(2023, 2, 18, 16, 28, 0, 0, time.UTC),
		SelfLink:    "test_self_link",
		Link:        "test_link",
		Author:      "test_author",
	})

	bytes, err := xml.Marshal(feed)
	if err != nil {
		t.Fatal(err)
	}

	if string(bytes) != testXml {
		t.Fatalf("Invalid XML result:\n%s\n", bytes)
	}
}
