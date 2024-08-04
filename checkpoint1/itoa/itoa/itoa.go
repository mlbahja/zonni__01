package itoa

func Itoa(n int) string {
	resullt := ""
	if n == 0 {
		resullt = "0"
	}
	if n < 0 {
		resullt += "-"
		n = -n
	}
	for n > 0 {
		resullt += string(rune(n%10) + '0')
		n /= 10
	}
	resullt1 := []rune(resullt)
	i := 0
	j := len(resullt) - 1
	if resullt1[i] == '-' {
		i++
	}
	for j > i {
		resullt1[i], resullt1[j] = resullt1[j], resullt1[i]
		i++
		j--

	}

	return string(resullt1)
}
