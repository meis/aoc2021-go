package main

type path []string

func (p path) countDoubleSmallCaves() int {
	small := make(map[string]int)
	for _, c := range p {
		if c[0] > 97 {
			small[c]++
		}
	}
	count := 0
	for _, v := range small {
		if v > 1 {
			count += v - 1
		}
	}

	return count
}
