package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type School struct {
	Name string
}
type Student struct {
	Name string
	Age  int
	*School
}

func main() {
	stu1 := Student{
		"gw123",
		1,
		&School{
			Name: "xytschool",
		},
	}

	stu2 := Student{
		"gw123",
		1,
		&School{
			Name: "xytschool",
		},
	}

	stu3 := &stu2
	fmt.Println("stu1", stu1)
	fmt.Println("stu2", stu2)
	fmt.Println("stu1 == stu2----------------------------------", stu1 == stu2)
	fmt.Println("reflect.DeepEqual(stu1, stu2)-----------------", reflect.DeepEqual(stu1, stu2))
	fmt.Println("unsafe.Pointer(&stu1) == unsafe.Pointer(&stu2)", unsafe.Pointer(&stu1) == unsafe.Pointer(&stu2))
	fmt.Println("unsafe.Pointer(&stu1) == unsafe.Pointer(stu3) ", unsafe.Pointer(&stu1) == unsafe.Pointer(stu3))
	fmt.Println("unsafe.Pointer(&stu2) == unsafe.Pointer(stu3) ", unsafe.Pointer(&stu2) == unsafe.Pointer(stu3))
	fmt.Println()

	v1 := []int{1, 2, 3}
	v2 := []int{1, 2, 3}
	fmt.Println("v1", v1)
	fmt.Println("v2", v2)
	fmt.Println("&v1 == &v2------------------------------------", &v1 == &v2)
	//fmt.Println(v1 == v2) 语法错误
	fmt.Println("reflect.DeepEqual(v1, v2)---------------------", reflect.DeepEqual(v1, v2))
	fmt.Println()
}
