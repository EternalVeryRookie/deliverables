package dp

func initLevenshteinDistanceMemo(s1Length, s2Length int) [][]int {
	memo := make([][]int, s2Length+1)
	for i := range memo {
		memo[i] = make([]int, s1Length+1)
		if i == 0 {
			continue
		}
	}

	for i := range memo {
		memo[i][0] = i
	}

	for i := range memo[0] {
		memo[0][i] = i
	}

	return memo
}

func minInt(a, b, c int) int {
	tmp := a
	if a > b {
		tmp = b
	}

	if tmp > c {
		return c
	}

	return tmp
}

//s1とs2のレーベンシュタイン距離を計算する
func levenshteinDistance(s1, s2 string) int {
	memo := initLevenshteinDistanceMemo(len(s1), len(s2))
	for i := 1; i < len(s2)+1; i++ {
		for j := 1; j < len(s1)+1; j++ {
			replaceCost := 0
			if s1[j-1] != s2[i-1] {
				replaceCost = 1
			}

			memo[i][j] = minInt(
				memo[i-1][j]+1,             //追加
				memo[i][j-1]+1,             //削除
				memo[i-1][j-1]+replaceCost, //置き換え
			)
		}
	}

	return memo[len(s2)][len(s1)]
}
