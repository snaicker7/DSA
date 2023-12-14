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
	meeting := []Meeting{
		{start: 1, end: 2},
		{start: 5, end: 7},
		{start: 8, end: 9},
		{start: 3, end: 4},
		{start: 0, end: 6},

		{start: 5, end: 9},
	}
	sort.Slice(meeting, func(i, j int) bool {
		return meeting[i].end < meeting[j].end
	})
	fmt.Println(meeting)
	currentend := meeting[0].end
	var result int = 1
	for i := 1; i < len(meeting); i++ {
		if meeting[i].start < currentend {
			result++
			currentend = meeting[i].end
		}
	}
	fmt.Println(result)
}
