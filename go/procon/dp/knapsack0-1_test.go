package dp

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
)

func unmarshalTestcase(t *testing.T, testcaseFilepath string, target interface{}) {
	t.Helper()

	b, err := ioutil.ReadFile(testcaseFilepath)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKnapsack0_1(t *testing.T) {
	type testcase struct {
		Capacity int       `json:"capacity"`
		Input    []article `json:"input"`
		Want     int       `json:"want"`
	}

	const testcaseFilepath = "testdata/knapsack0-1_test.json"
	testcases := []testcase{}
	unmarshalTestcase(t, testcaseFilepath, &testcases)

	for i := range testcases {
		solve := SolveKnapsack0_1(testcases[i].Input, testcases[i].Capacity)
		if solve != testcases[i].Want {
			t.Errorf("testcase %d, want: %d, actual: %d\n", i, testcases[i].Want, solve)
		}
	}
}

func TestKnapsack0_1Performance(t *testing.T) {
	const capacity = 10000
	const N = 100
	articles := make([]article, N)
	for i := range articles {
		weight := rand.Intn(1001) + 1
		value := rand.Intn(1001) + 1
		articles[i] = article{
			Weight: weight,
			Value:  value,
		}
	}

	threshold := int64(1)
	now := time.Now()
	SolveKnapsack0_1(articles, capacity)
	if time.Since(now).Milliseconds() > threshold*1000 {
		t.Errorf("実行時間が%d秒を超えました", threshold)
	}
}
