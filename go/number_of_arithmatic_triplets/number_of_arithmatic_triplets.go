package main

import "fmt"

func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	fmt.Println(arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
}

func arithmeticTriplets(nums []int, diff int) int {
	res := 0
	x := 0
	y := 1
	triplet := 0
	for x < len(nums) && y < len(nums) {
		if nums[y]-nums[x] == diff {
			triplet++
			if triplet == 2 {
				x = 0
				y = 1
				nums = nums[x+1:]
				res++
				continue
			}
			x = y
			y++
		} else if nums[y]-nums[x] < diff {
			y++
		} else {
			x++
		}
	}

	return res
}
