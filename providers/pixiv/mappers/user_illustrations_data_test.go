package mappers_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/providers/pixiv/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
)

func TestIllustrationEntriesExtraction(t *testing.T) {
	var contents m.Response[m.IllustrationDataBody]
	err := testutil.ReadJson("testdata/user_illustration_data_body.json", &contents)
	if err != nil {
		t.Fatal(err)
	}

	entries, err := mappers.PluckIllustrationEntries(&contents)
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
