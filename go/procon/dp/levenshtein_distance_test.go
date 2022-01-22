package dp

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	type testcase struct {
		S1   string
		S2   string
		Want int
	}

	const testcaseFilepath = "testdata/levenshtein_distance_test.json"
	testcases := []testcase{}
	unmarshalTestcase(t, testcaseFilepath, &testcases)
	for i := range testcases {
		actual := levenshteinDistance(testcases[i].S1, testcases[i].S2)
		if actual != testcases[i].Want {
			t.Errorf("%d   want: %d, actual: %d", i, testcases[i].Want, actual)
		}
	}
}
