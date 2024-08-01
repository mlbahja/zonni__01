package main

import (
	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	s := "test1 tes    t2 te   st3 te   st4                    test5 te    st6 test7 "
	f := SplitWhiteSpaces(s)
	for i, char := range rotstring(f) {
		printstr(char)
		if i < len(rotstring(f))-1 {
			printstr(" ")
		}
	}
	printstr("\n")
}

func rotstring(s []string) []string {
	var res []string
	for _, char := range s[1:] {
		if char != "" {
			res = append(res, char)
		}
	}
	res = append(res, s[0])
	return res
}

func SplitWhiteSpaces(s string) []string {
	str := ""
	slicstring := []string{}
	count := 0
	for _, valeu := range s {
		if valeu != ' ' || valeu <= 9 && valeu >= 13 {
			str += string(valeu)
			count++
		} else if count != 0 {
			slicstring = append(slicstring, str)
			str = ""
			count = 0
		}
	}
	if str != "" {
		slicstring = append(slicstring, str)
	}
	return slicstring
}
