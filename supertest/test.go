package main

import "fmt"

type SuperObj struct {
	objName string
}

func (this *SuperObj) SetName(name string) {
	this.objName = name
}

func (this *SuperObj) GetName() string {
	return this.objName
}

type SubObj struct {
	SuperObj
	objName string
}

func (this *SubObj) GetName() string {
	return this.objName
}

func main() {
	subObj := SubObj{}
	subObj.SuperObj.objName = "super01"
	subObj.objName = "sub_01"

	fmt.Println("subObj.GetName()\t", subObj.GetName())
	fmt.Println("subObj.SuperObj.GetName()\t", subObj.SuperObj.GetName())

	subObj.SetName("New name")
	fmt.Println("subObj.GetName()\t", subObj.GetName())
	fmt.Println("subObj.SuperObj.GetName()\t", subObj.SuperObj.GetName())
}
