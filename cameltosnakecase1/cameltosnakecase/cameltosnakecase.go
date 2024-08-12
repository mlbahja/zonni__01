package test

func iscontien_dijits(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func isUpper(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func isMorethan_or_egale_two_alpha(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if isUpper(s[i]) && isUpper(s[i+1]) {
			return true
		}
	}
	return false
}

func hasUpperAtEnd(s string) bool {
	i := len(s) - 1
	return isUpper(s[i])
}

func CamelToSnakeCase(s string) string {
	res := ""
	if s == "" {
		return s
	}
	if isMorethan_or_egale_two_alpha(s) {
		return s
	}
	if iscontien_dijits(s) {
		return s
	}
	if hasUpperAtEnd(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isUpper(c) && i == 0 {
			res += string(c)
		} else if isUpper(c) {
			res += "_" + string(c)
		} else {
			res += string(c)
		}
	}
	return res
}
