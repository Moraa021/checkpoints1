package main

import "github.com/01-edu/z01"

// PrintNumber handles numbers greater than 9 without using strconv
func PrintNumber(n int) {
	if n > 9 {
		PrintNumber(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func CountRepeats(s string) {
	if s == "" {
		return
	}
	counter := 1

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			counter++
		} else {
			// Print the character itself
			z01.PrintRune(rune(s[i]))

			// If it repeated, print the counter digits
			if counter > 1 {
				PrintNumber(counter)
			}
			// Reset the counter for the next unique character
			counter = 1
		}
	}
	// Print a newline at the end of the string run
	z01.PrintRune('\n')
}

func main() {
	CountRepeats("aaabbbccc") // Prints: a3b3c3
	CountRepeats("abc")       // Prints: abc
	CountRepeats("abbccc")    // Prints: ab2c3
}
