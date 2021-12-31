package dp

type article struct {
	Value  int `json:"value"`
	Weight int `json:"weight"`
}

func SolveKnapsack0_1(articles []article, capacity int) int {
	memo := make([][]int, len(articles))
	for i := range memo {
		memo[i] = make([]int, capacity+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return solveWithMemo(articles, 0, capacity, memo)
}

func solveWithMemo(articles []article, i int, capacity int, memo [][]int) int {
	if i >= len(articles) {
		return 0
	}

	a := memo[i][capacity]
	if a < 0 && articles[i].Weight <= capacity {
		a = solveWithMemo(articles, i+1, capacity-articles[i].Weight, memo) + articles[i].Value
		memo[i][capacity] = a
	}

	b := solveWithMemo(articles, i+1, capacity, memo)

	if a > b {
		return a
	} else {
		return b
	}
}
