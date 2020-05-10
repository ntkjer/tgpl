package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...))
	fmt.Println(min(values...))
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return -1
	}
	curr := vals[0]
	for _, val := range vals {
		if val >= curr {
			curr = val
		}
	}
	return curr
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return -1
	}
	curr := vals[0]
	for _, val := range vals {
		if val < curr {
			curr = val
		}
	}
	return curr
}
