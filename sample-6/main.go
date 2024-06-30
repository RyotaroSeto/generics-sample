// GOEXPERIMENT=rangefunc
package sample6

import (
	"fmt"
	"log"
)

func main() {
	for v := range seq1 {
		log.Println(v)
	}
	for v := range seq2 {
		log.Println(v)
	}
}

func seq1(yield func(int) bool) {
	fmt.Println("in seq1")
}

func seq2(yield func(int) bool) {
	yield(100)
	yield(200)
}
