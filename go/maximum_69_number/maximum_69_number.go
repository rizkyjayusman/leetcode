package main

import "fmt"

func main() {
	num := 9669
	x := 1
	for {
		num /= 10
		if num == 0 {
			break
		}
		x *= 10
	}

	x *= num / x
	fmt.Println(x)
	// fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	res := 0
	for i := 1; i <= 4; i++ {
		x := 1
		for {
			res += (num % 10) * x
			num /= 10
			x *= 10
			fmt.Println(res)
			if num == 0 {
				break
			}
		}
	}
	return res
}
