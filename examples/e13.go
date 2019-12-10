package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

const (
	a byte = 100 // int to byte
	//这里编译不过去
	// b int = 1e20 // float64 to int, overflows
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	fmt.Println(1 << 2)

	fmt.Println(KB, MB, GB, TB)
}
