package main

import "fmt"

func main() {
	fixture := []string{
		"{ [ ] ( ) }",
		"{ [ () { []} ] ( ) }",
		"{ [ ( ] ) }",
		"{ [ }",
	}

	for _, str := range fixture {
		fmt.Println(str, "=>", validate(str))
	}
}

func validate(str string) bool {
	counts := make(map[rune]int, 3)
	lastOpeners := make([]rune, len(str))

	for _, c := range str {
		if c == '[' || c == '(' || c == '{' {
			counts[c]++
			lastOpeners = append(lastOpeners, c)
		}
		if c == ']' || c == ')' || c == '}' {
			increment(counts, c)
			lastIndex := len(lastOpeners) - 1
			if match(lastOpeners[lastIndex], c) {
				lastOpeners = lastOpeners[0:lastIndex]
			} else {
				return false
			}
		}
	}

	for _, count := range counts {
		if count%2 != 0 {
			return false
		}
	}

	return true
}

func increment(m map[rune]int, c rune) {
	if c == ']' {
		m['[']++
	} else if c == ')' {
		m['(']++
	} else if c == '}' {
		m['{']++
	}
}

func match(a, b rune) bool {
	if a == '[' && b == ']' {
		return true
	} else if a == '(' && b == ')' {
		return true
	} else if a == '{' && b == '}' {
		return true
	}
	return false
}
