package main

import (
	"fmt"
	"runtime"
	"time"
)

type MyInt int

func OnFinalizer(v interface{}) {
	fmt.Printf("%T ,%v \n", v, v)
}

func Call() {
	a1 := "hello"
	b2 := "hello"
	var b3 MyInt = 1000
	b4 := new(MyInt)
	*b4 = 200


	runtime.SetFinalizer(&a1, OnFinalizer)
	runtime.SetFinalizer(&b2, OnFinalizer)
	runtime.SetFinalizer(&b3, OnFinalizer)
	runtime.SetFinalizer(b4, OnFinalizer)

	var rd MyInt = MyInt(999)
	r := &rd
	runtime.SetFinalizer(r, OnFinalizer)
}

func main() {
	Call()
	time.Sleep(time.Second)
	runtime.GC()
	for {
		time.Sleep(time.Second)
	}
}
