package dp

import "math"

func BinarySearch(input []int, target int) (index int) {
	return bisect(input, 0, len(input), target)
}

func bisect(input []int, leftIndex, rightIndex, target int) int {
	if leftIndex == 0 && rightIndex < 1 {
		return -1
	}

	center := (rightIndex + leftIndex) / 2
	if center == len(input)-1 {
		if input[len(input)-1] < target {
			return len(input)
		} else {
			return len(input) - 1
		}
	}

	if input[center] <= target && input[center+1] > target {
		return center
	}

	if input[center] < target {
		return bisect(input, center+1, rightIndex, target)
	} else {
		return bisect(input, leftIndex, center, target)
	}
}

func solveLongestIncreasingSubsequence(slice []int) int {
	dpTable := make([]int, len(slice))
	for i := range dpTable {
		dpTable[i] = math.MaxInt64
	}

	dpTable[0] = slice[0]

	for length := range slice {
		index := BinarySearch(dpTable, slice[length])
		if index >= len(dpTable)-1 {
			index = len(dpTable) - 2
		}

		if index < 0 {
			dpTable[index+1] = slice[length]
		} else if dpTable[index] != slice[length] {
			dpTable[index+1] = slice[length]
		}
	}

	length := 0
	for i := range dpTable {
		if dpTable[i] < math.MaxInt64 {
			length++
		} else {
			break
		}
	}

	return length
}
