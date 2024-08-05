package printmemory

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"
	for i, nbr := range arr {
		z01.PrintRune(rune(base[nbr/16]))
		z01.PrintRune(rune(base[nbr%16]))

		if ((i+1)%4 == 0 && i != 0) || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	// Print ASCII graphic characters
	for _, nbr := range arr {
		if nbr >= 33 && nbr <= 126 {
			z01.PrintRune(rune(nbr))
		} else {
			z01.PrintRune('.')
		}
	}

	// Pad the remaining length with dots if needed
	for i := len(arr); i < 10; i++ {
		z01.PrintRune('.')
	}
	z01.PrintRune('\n')
}
