package main

import "fmt"

type Student struct {
	name string
}

func main()  {
	var v interface{}
	v = &Student{}
	switch msg := v.(type) {
	case *Student:
		fmt.Println("*Student")
	case Student:
		fmt.Println("Student")
	}
}
