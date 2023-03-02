package main

import "fmt"

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
	fmt.Println(getConcatenation([]int{1, 3, 2, 1}))
}

func getConcatenation(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		fmt.Println(i)
		nums = append(nums, nums[i])
	}
	return nums
}
