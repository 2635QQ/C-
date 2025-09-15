package main

import (
	"fmt"
)

//定义变量
func main() {
	// fmt.Println("Hello, World!")//自带换行
	// fmt.Print("Hello, World!	Hello, World!") //不换行
    // fmt.Printf ("Hello, World!")//格式化输出


	// var a = "aaa"  //定义变量

	// fmt.Println(a)  //会输出变量

	// fmt.Printf("%v",a) //格式化输出a

	//定义变量  1
	// var a int = 10

	// var b int = 3

	// var c int = 5

	// fmt.Println("a=",a,"b=",b,"c=",c)  //自动换行

	// fmt.Printf("a=%v b=%v c=%v\n",a,b,c)

	// fmt.Printf("a=%v b=%v c=%v",a,b,c)

// 类型推导方式定义变量

a :=10
b :=3
c :=5

fmt.Printf("a=%v b=%v c=%v",a,b,c)

fmt.Printf("a=%v a的类型是%T",a,a)
	
}