package main

import "fmt"

type Stu struct {
	Name string
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}


	//v2 := stus[0]
	//for _, stu := range stus {
	//	v := stu
	//	v2 = stu
	//	m[stu.Name] = &v
	//	fmt.Printf("stu:%p , v:%p  v2:%p\n", &stu, &v, &v2)
	//}

	//for i := 0; i < len(stus); i++ {
	//	m[stus[i].Name] = &stus[i]
	//}

	for i, stu := range stus {
		m[stu.Name] = &stus[i]
		fmt.Printf("stu:%+v , v:%+v  \n", stu, stus[i])
	}


	for key, v := range m {
		fmt.Printf("key: %s , addr : %p ,name :%s, Age :%d \n", key, v, v.Name, v.Age)
	}
}




func t2() {
	m := make(map[string]*Stu)
	m["gw"] = &Stu{Name: "gw123"}
	m["gw1"] = &Stu{Name: "gw123"}
	m["gw2"] = &Stu{Name: "gw123"}

	for key, v := range m {
		fmt.Printf("addr: %p , key : %s , naneAddr: %p, name :%s\n", v, key, &v.Name, v.Name)
	}
}

func main() {
	//t2()

	pase_student()
}
