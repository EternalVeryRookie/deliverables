package dp

func solveKnapsack(articles []article, capacity int) int {
	memo := make([][]int, len(articles))
	for i := range memo {
		memo[i] = make([]int, capacity+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return solveKnapsackWithMemo(articles, 0, capacity, memo)
}

func solveKnapsackWithMemo(articles []article, i, capacity int, memo [][]int) int {
	if i >= len(articles) {
		return 0
	}

	if memo[i][capacity] < 0 {
		if articles[i].Weight <= capacity {
			tmp1 := solveKnapsackWithMemo(articles, i, capacity-articles[i].Weight, memo) + articles[i].Value
			tmp2 := solveKnapsackWithMemo(articles, i+1, capacity-articles[i].Weight, memo) + articles[i].Value
			if tmp1 > tmp2 {
				memo[i][capacity] = tmp1
			} else {
				memo[i][capacity] = tmp2
			}
		}

		c := solveKnapsackWithMemo(articles, i+1, capacity, memo)
		if c > memo[i][capacity] {
			memo[i][capacity] = c
		}
	}

	return memo[i][capacity]
}
