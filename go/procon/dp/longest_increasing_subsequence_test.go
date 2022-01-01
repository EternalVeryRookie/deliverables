package dp

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSolveLongestIncreasingSubsequence(t *testing.T) {
	type testcase struct {
		Input []int `json:"input"`
		Want  int   `json:"want"`
	}

	testcases := []testcase{}
	const testcaseFilepath = "testdata/longest_increasing_subsequence_test.json"
	unmarshalTestcase(t, testcaseFilepath, &testcases)

	for i := range testcases {
		actual := solveLongestIncreasingSubsequence(testcases[i].Input)
		if actual != testcases[i].Want {
			t.Errorf("case %d, want: %d, actual: %d\n", i, testcases[i].Want, actual)
		}
	}
}

func TestSolveLongestIncreasingSubsequencePerformance(t *testing.T) {
	const N = 100000
	slice := make([]int, N)
	for i := range slice {
		slice[i] = rand.Intn(int(math.Pow10(9)))
	}

	threshold := int64(1)
	now := time.Now()
	solveLongestIncreasingSubsequence(slice)
	if time.Since(now).Milliseconds() > threshold*1000 {
		t.Errorf("実行時間が%d秒を超えました", threshold)
	}
}
