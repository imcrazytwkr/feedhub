package mappers_test

import (
	"encoding/json"
	"os"
	"testing"

	m "github.com/imcrazytwkr/feedhub/providers/j-novel/models"
	"google.golang.org/protobuf/proto"
)

func TestNewsItemsExtraction(t *testing.T) {
	feedFile, err := os.ReadFile("test_data/response.pb")
	if err != nil {
		t.Fatal(err)
	}

	var response m.Payload
	err = proto.Unmarshal(feedFile, &response)
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := json.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("test_data/actual_response.json", encoded, 0644)
	if err != nil {
		t.Fatal(err)
	}

	/* var expectedEntries []*models.Entry
	err = testutil.ReadJson("test_data/expected_entries.json", &expectedEntries)
	if err != nil {
		t.Fatal(err)
	}

	if !testutil.SlicesEqual(entries, expectedEntries) {
		t.Log("slices aren't equal")
		t.FailNow()
	} */
}
