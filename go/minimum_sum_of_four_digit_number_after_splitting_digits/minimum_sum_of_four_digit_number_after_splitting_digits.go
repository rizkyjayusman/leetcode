package main

import "fmt"

func main() {
	fmt.Println(minimumSum(2932))
	// fmt.Println(minimumSum(4009))
}

func minimumSum(num int) int {
	if num < 1000 && num > 9999 {
		return -1
	}

	i := 1
	min := num
	for i < 1000 {
		i *= 10
		res := (num / i) + (num % i)
		fmt.Printf("%v + %v = %v\n", num/i, (num % i), res)
		if res < min {
			min = res
		}
	}
	return min
}
