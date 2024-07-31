package main

import "fmt"

func ItoaBase(value, base int) string {
	result := ""
	isneg := false
	if base < 2 || base > 16 {
		return ""
	}
	if value < 0 {
		value *= -1
		isneg = true
	}
	for value != 0 {
		conv := value % base
		if conv >= 10 {
			result = string(rune(conv-10+'A')) + result
		} else {
			result = string(rune(conv+'0')) + result
		}
		value /= base
	}
	if isneg {
		result = "-" + result
	}
	return result
}

func main() {
	str := ItoaBase(1, 1)
	fmt.Println(str)
}
