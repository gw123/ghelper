package main

import "fmt"

type Cat struct {
	Name string
}

func Test(cat2 *Cat) {
	*cat2 = Cat{
		Name: "zhengsheng",
	}
	cat2Addr := &cat2
	fmt.Println("cat2Addr", cat2Addr)
	return
}

func Test2(cat2 *Cat) {
	fmt.Printf("cat1Addr %p\n", cat2)
	cat2 = &Cat{Name: "zw"}
	fmt.Printf("cat1Addr %p\n", cat2)
	return
}

type CatPtr *Cat

func Test3(cat2 *CatPtr) {
	fmt.Printf("cat1Addr %p\n", cat2)
	var temp CatPtr = &Cat{Name: "zw"}
	cat2 = &temp
	fmt.Printf("cat1Addr %p\n", cat2)
	return
}

func main() {
	var cat1 CatPtr = &Cat{
		Name: "gw",
	}
	//fmt.Printf("cat1Addr %p\n", &cat)

	//Test2(&cat)

	fmt.Printf("CatPtr %p\n", cat1)
	Test3(&cat1)
	fmt.Printf("CatPtr %p\n", cat1)
	fmt.Println(cat1)
	//fmt.Println("++++++++++++++++ ")
	//Test(cat1)
	//fmt.Println(cat)
}
