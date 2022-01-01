package dp

import (
	"math/rand"
	"testing"
	"time"
)

func TestKnapsack(t *testing.T) {
	type testcase struct {
		Capacity int       `json:"capacity"`
		Input    []article `json:"input"`
		Want     int       `json:"want"`
	}

	const testcaseFilepath = "testdata/knapsack_test.json"
	testcases := []testcase{}
	unmarshalTestcase(t, testcaseFilepath, &testcases)

	for i := range testcases {
		solve := solveKnapsack(testcases[i].Input, testcases[i].Capacity)
		if solve != testcases[i].Want {
			t.Errorf("testcase %d, want: %d, actual: %d\n", i, testcases[i].Want, solve)
		}
	}
}

func TestKnapsackPerformance(t *testing.T) {
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
	solveKnapsack(articles, capacity)
	if time.Since(now).Milliseconds() > threshold*1000 {
		t.Errorf("実行時間が%d秒を超えました", threshold)
	}
}
