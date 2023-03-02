package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return nil
	}

	g := candies[0] + extraCandies
	var statuses []bool
	for k := range candies {
		statuses = append(statuses, candies[k]+extraCandies >= g)
	}
	return statuses
}
