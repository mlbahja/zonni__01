package main

import "fmt"

func itoaBase(value, val int) string {
	res := ""
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	is_negative := false
	if val < 2 || val > 16 {
		return ""
	}
	if value < 0 {
		value = -value
		is_negative = true
	}
	for value != 0 {
		num := value % val
		res = string(base[num]) + res
		value /= val
	}
	if is_negative {
		res = "-" + res
	}
	return res
}

func main() {
	fmt.Println(itoaBase(-16, 16))
}

// 16 base 16
// 6*10*0+ 10*10*1
