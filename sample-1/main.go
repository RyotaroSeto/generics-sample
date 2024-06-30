package sample1

import "fmt"

func main() {
	ptr := Ptr(false)
	fmt.Println(ptr)
}

// func Ptr(v bool) *bool {
// 	return &v
// }

// used generics
func Ptr[T any](v T) *T {
	return &v
}
