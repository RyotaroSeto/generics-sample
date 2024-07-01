// GOEXPERIMENT=rangefunc
package sample8

import (
	"fmt"
	"slices"
)

func main() {
	seq := Map(slices.All([]int{10, 20}), func(v int) string {
		return fmt.Sprintf("0x%x", v)
	})
	for i, v := range seq {
		fmt.Println(i, v)
	}

}

func Map[V1, V2 any](seq iter.Seq2[int, V1], f func(V1) V2) iter.Seq2[int, V2] {
	return func(yield func(int, V2) bool) {
		for i, v := range seq {
			if !yield(i, f(v)) {
				break
			}
		}
	}
}
