package main

import "strings"

type token string

func parseTokens(s string) []token {
	parts := strings.Split(s, "")
	tokens := []token{}

	for i := 0; i < len(parts); i++ {
		t := token(parts[i])
		if !t.isNum() {
			tokens = append(tokens, token(t))
		} else {
			j := i
			for j < len(parts) {
				if !token(parts[j]).isNum() {
					break
				}
				j++
			}

			value := strings.Join(parts[i:j], "")
			tokens = append(tokens, token(value))
			i += j - i - 1
		}
	}

	return tokens
}

func (t token) isNum() bool {
	return t != token("[") && t != token("]") && t != token(",")
}
