package main

import (
	"fmt"
	"sort"
)

type Meeting struct {
	start int
	end   int
}

func main() {
	cost := []int{6, 5, 7, 9}
	fmt.Println(mincostcandy(cost))
}
func mincostcandy(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})
	fmt.Println(cost)

	var result int = 0
	n := len(cost)
	for i := 0; i < n; i++ {
		if i%3 != 2 {
			result = result + cost[i]
		}

	}
	fmt.Println(result)
	return result
}
