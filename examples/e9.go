package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

type Teacher struct {
	name string
}

//测试golang类型判断 , %T %v

func main() {
	var v interface{}
	v = &Student{}
	switch msg := v.(type) {
	case *Student:
		fmt.Printf("*Student, %T : %v\n", msg, msg)
	case Student:
		fmt.Printf("Student, %T : %v\n", msg, msg)
	case Teacher:
		fmt.Printf("Teacher, %T : %v\n", msg, msg)
	case *Teacher:
		fmt.Printf("*Teacher, %T : %v\n", msg, msg)
	}

	fmt.Println()
	fmt.Println("reflect.TypeOf(v).String()", reflect.TypeOf(v).String())
	fmt.Println("reflect.TypeOf(v).Kind()", reflect.TypeOf(v).Kind())
	fmt.Println("reflect.TypeOf(v).Name()", reflect.TypeOf(v).Name())

	num := 100
	fmt.Println()
	fmt.Println("reflect.TypeOf(v).String()", reflect.TypeOf(num).String())
	fmt.Println("reflect.TypeOf(v).Kind()", reflect.TypeOf(num).Kind())
	fmt.Println("reflect.TypeOf(v).Name()", reflect.TypeOf(num).Name())
}
