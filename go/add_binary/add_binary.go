package main

import (
	"fmt"
	"strconv"
)

func main() {
	addBinary("11", "1")
}

func addBinary(a string, b string) string {
	i := 0
	if len(a) > len(b) {
		i = len(a)
	} else {
		i = len(b)
	}

	r := ""
	d := 0
	for j := i - 1; j >= 0; j-- {
		x := 0

		if j < len(a) {
			fmt.Printf("A :: %v\n", string(a[j]))
		}

		if j < len(a) && string(a[j]) == "1" {
			x = 1
		}

		y := 0

		if j < len(b) {
			fmt.Printf("B :: %v\n", string(b[j]))
		}

		if j < len(b) && string(b[j]) == "1" {
			y = 1
		}

		rr := (x + y + d)
		fmt.Printf("X :: %v\n", rr)
		if rr == 2 {
			r = r + "0"
			d = 1
		} else {
			r = r + strconv.Itoa(x+y)
			d = 0
		}
	}

	fmt.Println(r)

	return ""
}
