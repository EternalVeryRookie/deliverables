package procon

import (
	"math"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func isEqual(t *testing.T, want, actual []uint64) {
	for i := range want {
		if want[i] != actual[i] {
			t.Errorf("array equal error: want: %d, actual: %d", want[i], actual[i])
		}
	}
}

func TestConvertLengthMin2PowersMoreThan(t *testing.T) {
	var padding uint64 = 123456789
	type testdata struct {
		input []uint64
		want  []uint64
	}

	testcase := []testdata{
		{
			input: []uint64{1, 2, 3, 4, 5, 6, 7},
			want:  []uint64{1, 2, 3, 4, 5, 6, 7, padding},
		},
		{
			input: []uint64{1, 2, 3, 4},
			want:  []uint64{1, 2, 3, 4},
		},
		{
			input: []uint64{1},
			want:  []uint64{1, padding},
		},
	}

	for i, tCase := range testcase {
		actual := convertLengthMin2PowersMoreThan(tCase.input, padding)
		if len(actual) != len(tCase.want) {
			t.Errorf("length err %d, want: %d, actual: %d", i, len(tCase.want), len(actual))
		} else {
			isEqual(t, tCase.want, actual)
		}
	}
}

func linearMinSearch(arr []uint64, start, end int) uint64 {
	var min uint64 = math.MaxUint64
	for i := start; i < end; i++ {
		v := arr[i]
		if min > v {
			min = v
		}
	}

	return min
}

//確実に正しく実装できているであろう線形探索の結果とSegment Treeの探索結果を比較する
func TestRmqBySegmentTree(t *testing.T) {
	monoid := newMinOperateMonoid()
	maxGoroutine := runtime.NumCPU()
	c := make(chan struct{}, maxGoroutine)
	var wait sync.WaitGroup
	for testIndex := range make([]struct{}, 300) {
		c <- struct{}{}
		wait.Add(1)
		go func(testIndex int) {
			defer wait.Done()
			arr := make([]uint64, 100000)
			tree := NewSegmentTreeUint64(monoid, arr)
			for i := range arr {
				arr[i] = rand.Uint64()
				tree.Set(arr[i], i)
			}

			start := rand.Intn(len(arr))
			queryRange := Range{start: start, length: rand.Intn(len(arr) - start + 1)}

			want := linearMinSearch(arr, queryRange.start, queryRange.start+queryRange.length)
			actual := tree.Query(queryRange)
			if want != actual {
				t.Errorf("test patter %d, want: %d, actual: %d", testIndex, want, actual)
			}
			<-c
		}(testIndex)
	}

	wait.Wait()
}

func linearSum(arr []uint64, start, end int) uint64 {
	var sum uint64 = 0
	for i := start; i < end; i++ {
		sum += arr[i]
	}

	return sum
}

func TestRsqBySegmentTree(t *testing.T) {
	maxGoroutine := runtime.NumCPU()
	c := make(chan struct{}, maxGoroutine)
	var wait sync.WaitGroup
	for testIndex := range make([]struct{}, 300) {
		c <- struct{}{}
		wait.Add(1)
		go func(testIndex int) {
			defer func() {
				<-c
				wait.Done()
			}()

			arr := make([]uint64, 100000)
			tree := NewSegmentTreeUint64(newSumOperateMonoid(), arr)
			for i := range arr {
				newValue := rand.Uint64()
				tree.Set(0, i)
				arr[i] = newValue
				tree.Add(newValue, i)
			}

			start := rand.Intn(len(arr))
			queryRange := Range{start: start, length: rand.Intn(len(arr) - start + 1)}

			want := linearSum(arr, queryRange.start, queryRange.start+queryRange.length)
			actual := tree.Query(queryRange)
			if want != actual {
				t.Errorf("test patter %d, want: %d, actual: %d", testIndex, want, actual)
			}
		}(testIndex)
	}

	wait.Wait()
}
