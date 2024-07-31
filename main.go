package main

import "github.com/01-edu/z01"

func str(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	str("tzedgj")
}
