package main

import (
	"fmt"
	"time"
)

func main()  {
	go func() {
		defer func() {
			v := recover()
			if v !=nil{
				fmt.Println("panic 2", v)
			}
		}()

		go func() {
			defer func() {
				v := recover()
				if v !=nil{
					fmt.Println("panic 1", v)
				}
			}()
			go R1()
		}()

	}()

	select {

	}
}

func R1()  {
	defer func() {
		v := recover()
		if v !=nil{
			fmt.Println("panic R1", v)
		}
	}()
	fmt.Println("R1")
	go R2()
	time.Sleep(time.Second)
	panic("test")
	fmt.Println("R1------------")
}

func R2() {
	for ;; {
		time.Sleep(time.Second)
		fmt.Println("R2")
	}
}

/**
recover 可以恢复一个panic的协程 , 注意recover只能在defer中使用, 而且只能恢复当前recover所在的协程.

The recover built-in function allows a program to manage behavior of a
panicking goroutine. Executing a call to recover inside a deferred
function (but not any function called by it) stops the panicking sequence
by restoring normal execution and retrieves the error value passed to the
call of panic. If recover is called outside the deferred function it will
not stop a panicking sequence. In this case, or when the goroutine is not
panicking, or if the argument supplied to panic was nil, recover returns
nil. Thus the return value from recover reports whether the goroutine is
panicking.
 */