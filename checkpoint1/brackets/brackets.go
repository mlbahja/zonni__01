package Brackets

import "os"

func Brackets(s string) string {
	var str string
	for _, v := range s {
		if v == '[' || v == '{' || v == '(' || v == ']' || v == '}' || v == ')' {
			str += string(v)
		}
	}
	for i := 0; i < len(str); i++ {
		if i+2 <= len(str) && (str[i:i+2] == "()" || str[i:i+2] == "{}" || str[i:i+2] == "[]") {
			str = str[:i] + str[i+2:]
			i = -1
		}
	}
	if len(str) != 0 {
		return "Error"
	}
	return "Ok"
}

func main() {
	s := os.Args[1]
	Brackets(s)
}
