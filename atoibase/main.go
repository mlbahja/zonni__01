package main

import "fmt"

func AtoiBase(s string, base string) int {
	if len(base) < 2 || containsChar(base, '+') || containsChar(base, '-') {
		return 0
	}
	seen := make(map[rune]bool)
	for _, char := range base {
		if seen[char] {
			return 0
		}
		seen[char] = true
	}
	baseLen := len(base)
	value := 0

	for _, char := range s {
		index := indexRune(base, char)
		if index == -1 {
			return 0
		}
		value = value*baseLen + index
	}

	return value
}

func containsChar(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

func indexRune(s string, char rune) int {
	for i, c := range s {
		if c == char {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))      // 125
	fmt.Println(AtoiBase("1111101", "01"))          // 125
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF")) // 125
	fmt.Println(AtoiBase("uoi", "choumi"))          // 125
	fmt.Println(AtoiBase("bbbbbab", "-ab"))         // 0
}
