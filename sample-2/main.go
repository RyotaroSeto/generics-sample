package sample2

import "fmt"

func main() {
	var sum int
	Apply([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum)
}

// func Apply(s []int, f func(i, v int)) {
// 	for i, v := range s {
// 		f(i, v)
// 	}
// }

// used generics
func Apply[T any](s []T, f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}
