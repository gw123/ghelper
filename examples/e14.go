package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for value := range ch {
			fmt.Println(value)
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	go func() {
		//defer func() {
		//	close(ch)
		//}()
		time.Sleep(time.Second * 5)
		ch <- 10
		close(ch)
	}()
	wg.Wait()
}
