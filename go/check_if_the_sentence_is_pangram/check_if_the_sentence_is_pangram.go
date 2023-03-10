package main

import "fmt"

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}

func checkIfPangram(sentence string) bool {
	m := make(map[rune]struct{})
	for _, v := range sentence {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	return len(m) == 26
}
