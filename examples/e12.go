package main

import "fmt"

const (
	a1 = 1
	a2 = 4
	a3
	a4
	a5
	a6
)


const (
	b1 = 1
	b2 = 1
	b3 = 1
	b4 = iota
	b5
	b6
)

const (
	c1 = iota
	c2
	c3
	c4 = 100
	c5 = iota
	c6
)

func main() {
	fmt.Println(a1, a2, a3, a4, a5, a6)
	fmt.Println(b1, b2, b3, b4, b5, b6)
	fmt.Println(c1, c2, c3, c4, c5, c6)
}
