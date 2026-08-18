package mappers_test

import (
	"os"
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/pixiv/mappers"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"github.com/valyala/fastjson"
)

func TestIllustrationEntriesExtraction(t *testing.T) {
	sourceData, err := os.ReadFile("testdata/user_illustration_data_body.json")
	if err != nil {
		t.Fatal(err)
	}

	parser := &fastjson.Parser{}

	contents, err := parser.ParseBytes(sourceData)
	if err != nil {
		t.Fatal(err)
	}

	entries, err := mappers.PluckIllustrationEntries(contents)
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) == 0 {
		t.Fatal("No entries were parsed")
	}

	var expectedEntries []*models.Entry
	err = testutil.ReadJson("testdata/expected_illustration_entries.json", &expectedEntries)
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.SlicesEqual(entries, expectedEntries) {
		t.Log("slices aren't equal")
		t.FailNow()
	}
}
