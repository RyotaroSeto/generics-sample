package sample5

import "fmt"

func main() {
	var t *Tuple[int, string] = New(10, "hoge")
	fmt.Println(t.X, t.Y)
}

type Tuple[X, Y any] struct {
	X X
	Y Y
}

func New[X, Y any](x X, y Y) *Tuple[X, Y] {
	return &Tuple[X, Y]{
		X: x,
		Y: y,
	}
}
