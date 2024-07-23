package main

import "github.com/01-edu/z01"

func printcomb() {
	for r := '0'; r <= '7'; r++ {
		for v := r + 1; v <= '8'; v++ {
			for j := v + 1; j <= '9'; j++ {
				z01.PrintRune(r)
				z01.PrintRune(v)
				z01.PrintRune(j)
				if r != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	 z01.PrintRune('\n')
}

func main() {
	printcomb()
}
