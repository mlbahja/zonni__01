package main

import "fmt"

func ItoaBase(value, base int) string {
	rst := ""
	isneg := false
	if base < 2 || base > 16 {
		return rst
	}
	if value < 0 {
		value *= -1
		isneg = true
	}
	for value != 0 {
		result := value % base
		if result >= 10 {
			rst = string(rune(result)+55) + rst
		} else {
			rst = string(rune(result)+'0') + rst
		}
		value /= base
	}
	if isneg {
		rst = "-" + rst
	}
	return rst
}

func main() {
	str := ItoaBase(11, 16)
	fmt.Println(str)
}

/*
Write a function that:
converts an int value to a string using the specified base in the argument
and that returns this string
The base is expressed as an int, from 2 to 16.
The characters comprising the base are the digits from 0 to 9, followed by uppercase letters from A to F.
For example,the base 4 would be the equivalent
of "0123" and the base 16 would be the equivalent of "0123456789ABCDEF".
If the value is negative, the resulting string has to be preceded by a minus sign -.
Only valid inputs will be tested.*/
