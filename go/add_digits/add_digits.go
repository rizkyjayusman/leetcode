package main

import "fmt"

func main() {
	fmt.Println(addDigits(1212))
	// fmt.Println(addDigits(0))
}

func addDigits(num int) int {
	for num >= 10 {
		num = (num % 10) + (num / 10)
	}
	return num
}
