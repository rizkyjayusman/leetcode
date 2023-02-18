package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	res := 1
	y := 1
	for i := len(digits) - 1; i >= 0; i-- {
		res += digits[i] * y
		y *= 10
	}

	y /= 10

	arr := make([]int, len(digits)+1)
	length := 0
	for res%10 != 0 {
		arr[length] = res / y
		res %= y
		y /= 10
		length++
	}

	return arr
}
