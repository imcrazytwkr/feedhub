package mappers_test

import (
	"os"
	"testing"

	"github.com/imcrazytwkr/feedhub/providers/pixiv/mappers"
	"github.com/imcrazytwkr/feedhub/utils/testutil"
	"github.com/valyala/fastjson"
)

func TestIllustrationIdsExtraction(t *testing.T) {
	sourceData, err := os.ReadFile("test_data/user_illustration_ids_body.json")
	if err != nil {
		t.Fatal(err)
	}

	parser := &fastjson.Parser{}

	contents, err := parser.ParseBytes(sourceData)
	if err != nil {
		t.Fatal(err)
	}

	ids, err := mappers.PluckIllustrationIds(contents)
	if err != nil {
		t.Fatal(err)
	}

	var expectedIllustrationIds []int
	err = testutil.ReadJson("test_data/expected_illustration_ids.json", &expectedIllustrationIds)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Parsed ids: %v\n", ids)
	if !testutil.SlicesEqual(ids, expectedIllustrationIds) {
		t.FailNow()
	}
}
