package search

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

func TestBisect(t *testing.T) {
	const testPattern = 100
	const N = 1000000
	for i := range make([]struct{}, testPattern) {
		slice := make([]int, N)
		for j := range slice {
			slice[j] = rand.Int()
		}

		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		targetIndex := rand.Intn(len(slice))
		target := slice[targetIndex]

		actual := BinarySearch(slice, target)
		if actual != targetIndex {
			t.Errorf("case %d, want: %d, actual: %d", i, targetIndex, actual)
		}
	}

	slice := make([]int, N)
	for j := range slice {
		slice[j] = rand.Intn(N) + 100
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	target := 99
	actual := BinarySearch(slice, target)
	if actual != -1 {
		t.Errorf("want: %d, actual: %d", -1, actual)
	}

	actual = BinarySearch(slice, math.MaxInt64)
	if actual != len(slice)-1 {
		t.Errorf("want: %d, actual: %d", len(slice)-1, actual)
	}

	slice[len(slice)-1] = math.MaxInt64
	actual = BinarySearch(slice, slice[len(slice)-2]+1)
	if actual != len(slice)-2 {
		t.Errorf("want: %d, actual: %d", len(slice)-1, actual)
	}

	slice = []int{1, 3, 5, math.MaxInt64, math.MaxInt64, math.MaxInt64}
	if actual := BinarySearch(slice, 2); actual != 0 {
		t.Errorf("want: %d, actual: %d", 0, actual)
	}
}
