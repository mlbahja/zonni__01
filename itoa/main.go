package main

import "fmt"

func itoa(nn int) string {
	res := ""
	res1 := ""
	if nn < 0 {
		res += "-"
		nn = -nn
	}
	for nn > 0 {
		res += string(rune(nn%10) + '0')
		nn /= 10
	}
	res1 += res
	res11 := []rune(res1)
	i := 0
	j := len(res11) - 1
	if res1[i] == '-' {
		i++
	}
	for j > i {
		res11[i], res11[j] = res11[j], res11[i]
		i++
		j--
	}
	return string(res11)
}

func main() {
	num := itoa(-1234556)
	fmt.Println(num)
}
