package main

import "fmt"

func main() {
	str := "-1223456"
	result := 0
	sign := 1
	for i, char := range str {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				sign = -sign
			}
			continue
		}
		if char < '0' || char > '9' {
			fmt.Println('0')
			return
		}
		result = int(char-'0') + result*10
	}
	fmt.Println(result * sign)
}
