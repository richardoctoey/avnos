package number2

import "fmt"

func Reverser(i int) int {
	reversed := 0
	for i > 0 {
		lastNumber := i % 10
		reversed = reversed * 10 + lastNumber
		i = i/10
	}
	return reversed
}

func Reverse() {
	fmt.Println("Number 2:  Input:", 1234, "  >>> ", Reverser(1234))
	fmt.Println("Number 2:  Input:", 54321, " >>> ", Reverser(54321))
}