package main

import "fmt"

func main() {
	// fmt.Println(cellsInRange("K1:L2"))
	fmt.Println(cellsInRange("A1:F1"))
}

func cellsInRange(s string) []string {
	i := s[0]
	j := s[3]
	var str []string
	for i <= j {
		x := s[1]
		y := s[4]
		for x <= y {
			str = append(str, fmt.Sprintf("%c%c", rune(i), x))
			x++
		}
		i++
	}

	return str
}
