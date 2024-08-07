package main

import (
	"fmt"
)

func FirstWord(s string) string {
	result := ""
	is := false
	var n []string
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == ' ' && s[i+1] != ' ' && is {
			n = append(n, result)
			result = ""
		} else if s[i] != ' ' {
			is = true
			result += string(s[i])
		}
	}
	return n[len(n)-1]
}

func main() {
	input := "     Hel                lo,          wo           rld!                    test"
	output := FirstWord(input)
	fmt.Print(output)
}
