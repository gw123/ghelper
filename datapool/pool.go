package datapool

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()
	// 不使用对象池
	for i := 0; i < 10000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	// 使用对象池
	for i := 0; i < 10000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}

func main2() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)

	c := p.Get().(int)

	fmt.Println(a, b,c)
}
