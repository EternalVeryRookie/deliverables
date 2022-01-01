package dp

func solveKnapsack(articles []article, capacity int) int {
	memo := make([][]map[int]int, len(articles))
	for i := range memo {
		memo[i] = make([]map[int]int, capacity+1)
		for j := range memo[i] {
			memo[i][j] = map[int]int{}
		}
	}

	return solveKnapsackWithMemo(articles, 0, capacity, 1, memo)
}

func solveKnapsackWithMemo(articles []article, i, capacity, count int, memo [][]map[int]int) int {
	if i >= len(articles) {
		return 0
	}

	if _, ok := memo[i][capacity][count]; !ok {
		if articles[i].Weight <= capacity {
			tmp1 := solveKnapsackWithMemo(articles, i, capacity-articles[i].Weight, count+1, memo) + articles[i].Value
			tmp2 := solveKnapsackWithMemo(articles, i+1, capacity-articles[i].Weight, 1, memo) + articles[i].Value
			if tmp1 > tmp2 {
				memo[i][capacity][count] = tmp1
			} else {
				memo[i][capacity][count] = tmp2
			}
		}

		c := solveKnapsackWithMemo(articles, i+1, capacity, 1, memo)
		if c > memo[i][capacity][count] {
			memo[i][capacity][count] = c
		}
	}

	return memo[i][capacity][count]

}
