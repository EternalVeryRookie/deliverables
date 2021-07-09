package search

import "testing"

func TestLinearSearchBanhei(t *testing.T) {
	list := make([]int, 100)
	for i := range list {
		list[i] = i
	}

	type searchTarget struct {
		num      int
		isExists bool
	}

	searchTargets := []searchTarget{
		{
			num:      0,
			isExists: true,
		},
		{
			num:      99,
			isExists: true,
		},
		{
			num:      100,
			isExists: false,
		},
	}

	for i, target := range searchTargets {
		index := LinearSearchBanhei(list, target.num)
		if (index >= 0 && !target.isExists) || (index < 0 && target.isExists) {
			t.Errorf("error index: %d", i)
		}
	}
}
