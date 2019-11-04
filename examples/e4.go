package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func t1() {
	a := 1
	b := 2
	defer func() {
		calc("1", a, calc("10", a, b))
	}()
	a = 0
	defer func() {
		calc("2", a, calc("20", a, b))
	}()
	b = 1
}

func t2() {
	a := 1
	b := 2
	calc("1", a, calc("10", a, b))
	a = 0
	calc("2", a, calc("20", a, b))
	b = 1
}

func main() {
	t1()
	fmt.Println("@@@@@@@@@@@@@@@@@@")
	t2()
}

/*
 defer 执行顺序 后面的加的defer 先执行
 并且defer 记录调用时刻传入的参数
 10  1 2 3
 20  0 2 2
 2   0 2 2
 1   1 3 4
*/
