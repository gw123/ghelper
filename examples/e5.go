package main

import (
	"fmt"
	"time"
)

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	s2 := make([]int, 5, 5)
	s2 = append(s2, 1, 2, 3)
	fmt.Println(s2)

	ch := make(chan interface{})
	go func() {
		for ; ;  {
			v := <-ch
			time.Sleep(time.Second)
			fmt.Println(v)
		}
		close(ch)
	}()
	ch<- 1
	ch<- 1
	ch<- 1

}
