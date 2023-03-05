package main

import "fmt"

func main() {
	fmt.Println(separateDigits([]int{113, 25, 83, 77}))
}

func separateDigits(nums []int) []int {
	var arr []int
	for _, v := range nums {
		fmt.Println(v / 10)
		arr = append(arr, v)
	}
	return arr
}

func separate(num int) []int {
	var arr []int
	return append(arr)
}
