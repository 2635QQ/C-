package main

import (
	"fmt"
	"unsafe"
)

// 定义变量

func getUserinfo() (string, int) {
	return "张三", 18
}

var g = "全局变量" //全局变量
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

	// a :=10
	// b :=3
	// c :=5

	// fmt.Printf("a=%v b=%v c=%v",a,b,c)

	// fmt.Printf("a=%v a的类型是%T",a,a)

	// -----------------------------9月16号fmt

	// fmt.Println("1.var 变量的声明")
	// // var 变量名称 变量类型

	// var username string  //变量名称首字符不能是数字
	// fmt.Println("username=",username)  //变量声明没有初始化为空

	// var a1="张三"
	// fmt.Println("a1=",a1)  //关键字，保留字，不能做变量名

	//go 语言变量的定义以及初始化
	// var username string
	// username ="张三"

	// fmt.Println("username=",username)

	// var username2 string ="李思"   //第二种初始化的方式
	// fmt.Println("username2=",username2)

	//  var username3 ="王五" // 第三种初始化的方式
	//  fmt.Println("username3=",username3)

	// var username = "张三"  //同一个作用域内不支持重复声明
	// var age = 20
	// var sex = "男"

	// fmt.Println(username,age,sex)

	// username ="李四"  //这个代表赋值，是可以的

	// fmt.Println(username,age,sex)

	//一次定义多个变量

	// 1、 var 变量名称，变量名称，类型
	// 2、var {
	//  变量名称 类型

	//  变量名称 类型

	// }

	// var a1, a2 string

	// a1 = "aa" //赋值的类型必须和定义的类型是一致的；
	// a2 = "bbb"

	// fmt.Println("a1=", a1, "a2=", a2)

	// var (
	// 	username string  //这样声明的多个类型是不一致的
	// 	age      int
	// 	sex      string
	// )
	// username = "张三"
	// age = 20
	// sex = "男"

	// fmt.Println(username, age, sex)
	// var (
	// 	username = "张三"
	// 	age      = 20
	// )

	// fmt.Println(username, age)

	//短变量声明法  也就是用 ：= 来声明，只能在函数内部使用，不能作为全局变量使用

	// 原来的  var username = "张三"
	// username := "张三"

	// fmt.Println(username)

	// fmt.Printf("类型：%T",username)

	//使用短变量声明法一次声明多个变量，并初始化;只能声明局部变量

	// a, b, c := 12, 13, "C"
	// fmt.Printf("a的类型：%T , b的类型: %T  c的类型 :%T\n", a, b, c)

	// fmt.Println(g)

	// var username, age = getUserinfo()

	// fmt.Println(username, age) //这样可以获取到张三 18；

	//匿名变量 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量 （anonymous variable）
	//匿名变量 用一个下划线 _ 表示 ：

	// //也就是
	//  var username, _ = getUserinfo()

	// fmt.Println(username)
	// // 匿名变量之间不存在重复声明  因为匿名变量不占用命名空间，不会分配内容；

	// var _, age = getUserinfo()

	// fmt.Println(age)

	//常量
	// const pi = 3.14159

	// fmt.Println(pi)

	// const (
	// 	A = "A"

	// 	B = "B"
	// )

	// fmt.Println(A, B)

	// const (
	// 	n1 = 100
	// 	n2
	// 	n3
	// )

	// fmt.Println(n1, n2, n3)

	// //const 和iota ，计数器，一起使用
	// const a = iota
	// fmt.Println(a)

	// const (
	// 	n1 = iota
	// 	n2
	// 	n3
	// )

	// fmt.Println(n1, n2, n3)

	// //还可以跳过某个值使用，也就是、
	// const (s1 = iota
	// 	_
	// 	s2
	// 	s3
	// )

	// fmt.Println(s1,s2,s3)
	//iota 可以作为中间插队使用
	// const(
	// 	n1 = 2  // 0
	// 	n2 = 100 //100
	// 	n3 = iota
	// 	n4
	// )
	// fmt.Println(n1,n2,n3,n4)

	//多个iota 定义在一行

	// const (
	// 	n1, n2 = iota + 2, iota + 2 // 1 2
	// 	n3, n4                      // 2 3
	// 	n5, n6                      // 3 4
	// )
	// fmt.Println(n1, n2, n3, n4, n5, n6)

	var a int8 = 36

	fmt.Printf("num = %v 类型 ：%T\n", a, a)
	fmt.Println(unsafe.Sizeof(a))

}
