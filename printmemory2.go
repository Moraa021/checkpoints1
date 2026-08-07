package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexChars := "0123456789abcdef"

	// 1. Print hexadecimal representation
	for i, b := range arr {
		// Convert byte to hex by splitting it into two 4-bit parts
		upperNibble := b >> 4
		lowerNibble := b & 0x0f

		z01.PrintRune(rune(hexChars[upperNibble]))
		z01.PrintRune(rune(hexChars[lowerNibble]))

		// Format grid alignment
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	// If last line is not complete, add a newline
	if len(arr)%4 != 0 {
		z01.PrintRune('\n')
	}

	// 2. Print ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
