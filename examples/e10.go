package main

import "fmt"

func main()  {
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)

	/***
	在这个程序中，a 的值是 5 并且 a 在语法上是泛化的（它既可以表示浮点数 5.0，也可以表示整数 5，甚至可以表示没有虚部的复数 5 + 0i），
	因此 a 可以赋值给任何与之类型兼容的变量。像 a 这种数值常量的默认类型可以想象成是通过上下文动态生成的。
	var intVar int = a 要求 a 是一个 int，那么 a 就变成一个 int 常量。
	var complex64Var complex64 = a 要求 a 是一个 complex64，
	那么 a 就变成一个 complex64 常量。这很容易理解：）
	 */
	fmt.Println()
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)


	//fmt.Println()
	//const a2 int = 5
	//var intVar2 int = a2
	//var int32Var2 int32 = a2
	//var float64Var2 float64 = a2
	//var complex64Var2 complex64 = a2
	//fmt.Println("intVar",intVar2, "\nint32Var", int32Var2, "\nfloat64Var", float64Var2, "\ncomplex64Var",complex64Var2)


