package sample3

import "fmt"

func main() {
	ns := []int{1, 2, 3, 4}
	ms := Filter(ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms)
}

// func Filter(s []int, f func(i, v int) bool) []int {
// 	var s2 []int
// 	for i, v := range s {
// 		if f(i, v) {
// 			s2 = append(s2, v)
// 		}
// 	}
// 	return s2
// }

// used generics
func Filter[T any](s []T, f func(i int, v T) bool) []T {
	var s2 []T
	for i, v := range s {
		if f(i, v) {
			s2 = append(s2, v)
		}
	}
	return s2
}
