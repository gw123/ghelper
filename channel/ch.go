package channel

import (
	"fmt"
	"time"
)

//正常关闭一个channel
func test2() {
	ch1 := make(chan int, 10)
	ch1 <- 1

	go func() {
		time.Sleep(time.Second * 3)
		close(ch1)
	}()

	for ; ; {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("run", data)
			} else {
				fmt.Println("run over", data)
				return
			}
		}
		time.Sleep(time.Second)
	}
}


//测试channel 中有数据未处理,这时关闭close.未处理的数据channel是如何处理的
func testCloseChWithData() {
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 1
	ch1 <- 1
	ch1 <- 1

	go func() {
		time.Sleep(time.Second * 1)
		close(ch1)
	}()

	for ; ; {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("run", data)
			} else {
				fmt.Println("run over", data)
				return
			}
			break
		}
		time.Sleep(time.Second)
	}
}

