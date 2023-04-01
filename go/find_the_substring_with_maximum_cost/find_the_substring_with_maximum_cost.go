package main

type FindTheSubstringWithMaximumCost struct{}

func (ftswmc FindTheSubstringWithMaximumCost) MaximumCostSubstring(s string, chars string, vals []int) int {
	res := 0
	m := make(map[rune]int)
	for k, v := range chars {
		m[v] = vals[k]
	}

	for _, v := range s {
		if tmp, ok := m[v]; ok {
			res += tmp
			if res < 0 {
				res = 0
			}
			continue
		}
		res += int(v - 96)

	}
	return res
}
