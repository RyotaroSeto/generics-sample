// GOEXPERIMENT=rangefunc
package sample7

import (
	"fmt"
)

func main() {
	for c := range Alphabet {
		fmt.Printf("%c", c)
		if c == 'c' {
			break
		}
	}
}

func Alphabet(yield func(rune) bool) {
	for c := 'A'; c <= 'Z'; c++ {
		if !yield(c) {
			break
		}
	}
}
