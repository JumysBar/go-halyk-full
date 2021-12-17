package main

const N = 10000

type T struct {
	x int
}

//go:noinline
func f(t *T) {
	t.x = 0
	for i := 0; i < N; i++ {
		t.x += i
	}
}

//go:noinline
func g(t *T) {
	var x = 0
	for i := 0; i < N; i++ {
		x += i
	}
	t.x = x
}

func main() {
	t := &T{}
	f(t)
	g(t)
}
