package sample4

import "fmt"

func main() {
	var ss []string = Map([]int{10, 20}, func(n int) string {
		return fmt.Sprintf("0x%x", n)
	})
	fmt.Println(ss) // [0xa 0x14]
}

// func Map(s []int, f func(int) string) []string {
// 	s2 := make([]string, len(s))
// 	for i, v := range s {
// 		s2[i] = f(v)
// 	}
// 	return s2
// }

func Map[T1 any, T2 any](s []T1, f func(T1) T2) []T2 {
	s2 := make([]T2, len(s))
	for i, v := range s {
		s2[i] = f(v)
	}
	return s2
}
