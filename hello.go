package main

import (
	"deliverables/procon/search"
	"fmt"
)

func main() {
	list := make([]int, 10)
	for i := range list {
		list[i] = i
	}

	index := search.LinearSearchBanhei(list, 22)
	fmt.Printf("%d\n", index)
}
