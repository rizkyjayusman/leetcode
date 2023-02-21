package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(convertToTitle(1))
	// fmt.Println(convertToTitle(28))
	// fmt.Println(convertToTitle(701))
	// fmt.Println(convertToTitle(2759))
}

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		if columnNumber%26 > 0 {
			res = string(rune((columnNumber%26)+64)) + res
			columnNumber -= columnNumber % 26
		} else if columnNumber%26 == 0 {
			div := columnNumber / 26
			if div > 26 {
				diff := div % 26
				res = string(rune(diff+64)) + res
				columnNumber = div - diff
			} else {
				res = string(rune((columnNumber/26)+64)) + res
				columnNumber -= columnNumber
			}
		}
		time.Sleep(1 * time.Second)
	}
	return res
}
