package main

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	printbase(nbr, base)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' && base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func printbase(nbr int, base string) {
	baseLen := len(base)
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	digits := []rune{}
	for nbr > 0 {
		remainder := nbr % baseLen
		digits = append(digits, rune(base[remainder]))
		nbr /= baseLen
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func main() {
	PrintNbrBase(1, "0123456789")
}
