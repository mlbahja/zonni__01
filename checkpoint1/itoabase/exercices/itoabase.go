package exercice

func ItoaBase(value int, base int) string {
	sign := 1
	var num uint
	if base < 2 || base > 16 {
		return ""
	}
	if value < 0 {
		sign = -1
		num = uint(value * -1)
	} else {
		num = uint(value)
	}

	var num1 uint
	num1 = uint(base)

	digits := "0123456789ABCDEF"
	isNegative := num < 0
	res := ""
	if num == 0 {
		return "0"
	}

	if isNegative {
		num = -num
	}

	for num > 0 {
		res = string(digits[num%num1]) + res
		num /= num1
	}

	if sign == -1 {
		res = "-" + res
	}

	return res
}
