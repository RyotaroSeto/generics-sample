// GOEXPERIMENT=rangefunc
package sample7

import (
	"fmt"
)

func main() {
	for c := range Alphabet {
		fmt.Println("%c", c)
		if c == 'c' {
			break
		}
	}
}

func Alphabet(yield func(int) bool) {
	fmt.Println("in seq1")
}
