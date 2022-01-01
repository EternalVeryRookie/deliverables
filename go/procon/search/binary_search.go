package search

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
