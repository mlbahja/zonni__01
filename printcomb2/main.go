package main

import "github.com/01-edu/z01"

func printcomb2() {
	for i := ('0' - 48); i <= 98; i++ {
		for j := i + 1; j < 99; j++ {
			z01.PrintRune((i / 10) + '0')
			z01.PrintRune((i % 10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune((j / 10) + '0')
			z01.PrintRune((j % 10) + '0')
			if j != 99 || i != 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	printcomb2()
}
