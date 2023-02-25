package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int

	fmt.Scanf("%d", &n)
	tanks := make([]int, n)

	for {
		if n == 0 {
			break
		}
		fmt.Scanf("%d", &tanks[cap(tanks)-n])
		n--
	}

	var max_tank int
	min_tank := tanks[0]

	for _, tank := range tanks {
		max_tank = max(tank, max_tank)
		if tank < max_tank {
			fmt.Printf("%d\n", -1)
			return
		}
		min_tank = min(tank, min_tank)
	}

	fmt.Printf("%d\n", max_tank-min_tank)

	return
}
