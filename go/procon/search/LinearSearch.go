package search

func LinearSearchBanhei(list []int, key int) int {
	list = append(list, key)
	i := 0
	for {
		if list[i] != key {
			i++
		} else {
			break
		}
	}

	if i == len(list)-1 {
		return -1
	}

	return i
}
