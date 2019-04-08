package number1

import (
	"fmt"
	"strconv"
	"strings"
)

func BinaryGap(i int64) int {
	stringBinary := strconv.FormatInt(i, 2)
	counter := 0
	for _, elm := range stringBinary {
		if elm == '1' {
			counter+=1
		}
	}
	if counter == 1 {
		return 0
	}
	splittedBinary := strings.Split(stringBinary, "1")
	longest := 0
	for _, elm := range splittedBinary {
		if len(elm) > longest {
			longest = len(elm)
		}
	}

	return longest
}

func MaxBinaryGap(){
	fmt.Println("Number 1: ", BinaryGap(9))
	fmt.Println("Number 1: ", BinaryGap(529))
	fmt.Println("Number 1: ", BinaryGap(15))
	fmt.Println("Number 1: ", BinaryGap(32))
}
