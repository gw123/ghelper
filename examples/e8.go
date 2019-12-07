package main

import "fmt"

type person struct {
	attr map[string]interface{}
	addrs []string
}

func test1()  {
	p := new(person)
	p.attr["name"] = "gw123"
	fmt.Println(p.attr)
}

func test2()  {
	p := new(person)
	p.addrs = []string{"123"}
	fmt.Println(p.addrs)
}

//map 需要初始化
func main()  {
	test2()
}

