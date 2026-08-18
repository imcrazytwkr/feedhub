package mappers_test

import (
	"testing"

	"github.com/imcrazytwkr/feedhub/providers/pixiv/mappers"
	m "github.com/imcrazytwkr/feedhub/providers/pixiv/models"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
)

func TestIllustrationIdsExtraction(t *testing.T) {
	var contents m.Response[m.IllustrationIDsBody]
	err := testutil.ReadJson("testdata/user_illustration_ids_body.json", &contents)
	if err != nil {
		t.Fatal(err)
	}

	ids, err := mappers.PluckIllustrationIds(&contents)
	if err != nil {
		t.Fatal(err)
	}

	var expectedIllustrationIds []int
	err = testutil.ReadJson("testdata/expected_illustration_ids.json", &expectedIllustrationIds)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Parsed ids: %v\n", ids)
	if !testutil.SlicesEqual(ids, expectedIllustrationIds) {
		t.FailNow()
	}
}
