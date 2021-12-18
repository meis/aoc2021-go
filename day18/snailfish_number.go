package main

import "strconv"

type snailfishNumber []token

func parseSnailfishNumbers(in []string) []snailfishNumber {
	tokens := []snailfishNumber{}

	for _, s := range in {
		tokens = append(tokens, parseTokens(s))
	}

	return tokens
}

func (n1 snailfishNumber) add(n2 snailfishNumber) snailfishNumber {
	tokens := []token{}
	tokens = append(tokens, token("["))
	tokens = append(tokens, n1...)
	tokens = append(tokens, token(","))
	tokens = append(tokens, n2...)
	tokens = append(tokens, token("]"))

	reduced := true
	for reduced {
		tokens, reduced = reduce(tokens)
	}

	return tokens
}

func (num snailfishNumber) magnitude() int {
	if len(num) == 1 {
		val, _ := strconv.Atoi(string(num[0]))
		return val
	}

	startLeft := 1
	endLeft := 1
	opened := 0
	for true {
		if num[endLeft] == "," && opened == 0 {
			break
		} else if num[endLeft] == "[" {
			opened++
		} else if num[endLeft] == "]" {
			opened--
		}
		endLeft++
	}
	startRight := endLeft + 1
	endRight := len(num) - 1

	return (3 * num[startLeft:endLeft].magnitude()) + (2 * num[startRight:endRight].magnitude())
}

func reduce(n snailfishNumber) (snailfishNumber, bool) {
	opened := 0
	for i := 0; i < len(n); i++ {
		if n[i] == "[" {
			opened++

		} else if n[i] == "]" {
			opened--
		}

		if opened == 5 {
			leftPos := i + 1
			rightPos := i + 3
			for j := leftPos - 1; j >= 0; j-- {
				if n[j].isNum() {
					v1, _ := strconv.Atoi(string(n[leftPos]))
					v2, _ := strconv.Atoi(string(n[j]))

					n[j] = token(strconv.Itoa(v1 + v2))
					break
				}
			}
			for j := rightPos + 1; j < len(n); j++ {
				if n[j].isNum() {
					v1, _ := strconv.Atoi(string(n[rightPos]))
					v2, _ := strconv.Atoi(string(n[j]))

					n[j] = token(strconv.Itoa(v1 + v2))
					break
				}
			}

			prefix := n[0 : leftPos-1]
			suffix := n[rightPos+2:]
			n = append(prefix, "0")
			n = append(n, suffix...)

			return n, true
		}
	}

	for i := 0; i < len(n); i++ {
		if n[i].isNum() {
			val, _ := strconv.Atoi(string(n[i]))
			if val > 9 {
				left := val / 2
				right := val / 2
				if val%2 != 0 {
					right++
				}

				preffix := n[0:i]
				split := []token{"[", token(strconv.Itoa(left)), ",", token(strconv.Itoa(right)), "]"}
				suffix := n[i+1:]

				new := []token{}
				new = append(new, preffix...)
				new = append(new, split...)
				new = append(new, suffix...)

				return new, true
			}

		}
	}

	return n, false
}
