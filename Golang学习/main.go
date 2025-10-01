package main

import (
	"fmt"
	
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

	// var a int8 = 36

	// fmt.Printf("num = %v 类型 ：%T\n", a, a)
	// fmt.Println(unsafe.Sizeof(a))

	// var a int32 = 15

	// fmt.Printf("num=%v 类型：%T\n",a,a)
	// fmt.Println(unsafe.Sizeof(a))

	// var a1 int32 = 10
	// var a2 int64 = 21

	// fmt.Println(int64(a1) + a2) // 把a1转换成64位

	// var n1 int16 = 130

	// fmt.Println(int8(n1))   //

	//数字字面量语法  %d 表示10进制输出  %b 表示二进制输出

	// num := 10

	// fmt.Printf("num=%v  原样输出\n",num)

	// fmt.Printf("num=%d 十进制输出\n",num)   // 10进制输出

	// fmt.Printf("num=%b 二进制输出\n",num)   //二进制输出

	// fmt.Printf("num=%o 八进制输出\n",num)

	// fmt.Printf("num=%x 十六进制输出\n",num)

	// var a float32 = 3.128765

	// fmt.Printf("a=%v 类型：%T，  %f \n",a,a,a)

	// //若是想要保留2位小数，可以使用%.2f来进行操作；

	// fmt.Printf("a= %.2f\n",a)

	// //Golang 科学计数法表示浮点类型;表示3.14*10的2次方
	// var f1 float32 = 3.14e2

	// fmt.Printf("f1=%v --- %T",f1,f1)

	// var f2 float32 = 3.14e-2  //表示3.14除以10的2次方
	// fmt.Printf("f2=%v --- %T",f2,f2)

	//由于会出现浮点数精度丢失的问题，所以会进行下载引入第三方包，
	// 也就是  "github.com/shopspring/decimal"

	// 这个包需要先进行下载 go get -u github.com/shopspring/decimal

	// 使用加法运算；
	// var num1 float64 = 3.1
	// var num2 float64 = 4.2
	// d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))

	// fmt.Println(d1)

	// // 减法
	// d2 := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(num2))

	// fmt.Println(d2)

	// var flag = true
	// fmt.Printf("%v---%T\n", flag, !flag)

	// var b bool   //布尔类型默认是false
	// fmt.Printf("%v----%T\n",b,b)

	// //string 类型默认值也是 空
	// var a string
	// fmt.Printf("%v",a)

	// //int 类型默认值为0,float 类型默认是0

	// //布尔类型无法参与数值运算，也无法与其他类型进行转换；

	//关于字符串
	// var str1 string
	// var str2 = "hello"
	// str3 := "world"

	// fmt.Printf("%v--- %v----- %v\n", str1, str2, str3)

	// fmt.Printf("%T--- %T----- %T\n", str1, str2, str3)

	//字符串转义字符
	// str1 := "this \nis str"

	// fmt.Println(str1)

	// str2 := "C:\\Golang\\go.exe"
	// fmt.Println(str2)   //输出\ ，要用转义字符输出 \\

	// str3 := "C:GO\"bin"

	// fmt.Println(str3)

	//一次定义多行字符串，要用反引号，```
	// str1 := `this is str
	// line2  is str
	// line3 str
	// `
	// fmt.Println(str1)

	// len (str) 求长度
	// var str1 ="Nihao"
	// fmt.Println(len(str1))

	// var str2 ="你好"
	// fmt.Println(len(str2))  //中文占3个字节
	//拼接字符串 可以用 +  或者 fmt.Sprintf 拼接字符串

	// str1 :="hello"
	// str2 :="你好"
	// fmt.Println(str1 + str2)
	// str3 := fmt.Sprintf("%v  %v",str1, str2)
	// fmt.Println(str3)

	//string.Split 分割字符串   strings 需要引入strings包
	// var str1 = "123-456"
	// arr := strings.Split(str1,"-")

	// fmt.Println(arr)

	// // stings.Join  把切片连接成字符串
	// str2 := strings.Join(arr,"*")  // 第一个参数是要拼接的数组，或者切片，第二个是* ，用来拼接的字符串

	// fmt.Println(str2)

	// array := []string{"aa","java","golang"}
	// fmt.Println(array)

	// str3 := strings.Join(array,"-")

	// fmt.Println(str3)

	// strings.contains  判断是否包含某个字符串

	// str1 := "this is str"
	// str2 := "this"

	// flag := strings.Contains(str1,str2) //判断str1 是否包含 str2

	// fmt.Println(flag)

	// strings.HasPrefix, strings.HasSuffix   前缀/后缀判断

	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.HasPrefix(str1,str2) //前缀判断，是有的；
	// fmt.Println(flag)

	// // 后缀判断
	// str3 := "str"
	// flag2 := strings.HasSuffix(str1,str3)
	// fmt.Println(flag2)

	// //查找字串出现的位置，  strings.Index() , strings.LastIndex()
	// str1 := "this is str"
	// str2 := "is"

	// num := strings.Index(str1,str2)
	// fmt.Println(num)
	// num2 := strings.LastIndex(str1,str2)

	// fmt.Println(num2)

	// 字符类型的，通过遍历字符串元素获得字符；
	// var a = 'a'   //golang 中定义字符，字符属于int 类型

	// fmt.Printf("值： %v  类型：%T ",a,a)

	// 也就是直接输出byte 字符 的时候，输出的是这个字符对应的码值

	// 如果需要原样输出字符，就是用%c
	// fmt.Printf("值： %c  类型 ： %T",a,a)

	//定义 字符串，输出字符串内容

	// var str ="this"

	// fmt.Printf("值： %v   原样输出 %c,类型 %T",str[2],str[2],str[2])

	//unsafe 没法查看string 类型数据所占用的存储空间

	// fmt.Println(len(str))

	// var a ='国'
	// fmt.Printf("值：%v   类型：%T\n",a,a)  // Unicode 编码 10进制,zhi :

	// fmt.Printf("值： %c   类型 ：%T",a,a)  // 原样输出字符

	// // 通过循环输出字符串里面的字符
	// str := "你好 golang"
	// for i:= 0; i < len(str); i++{   //for 循环相当于 byte 类型
	// 	fmt.Printf("%v(%c)",str[i],str[i])
	// }

	// // range 循环相当于 rune 类型 还适合处理 UTF-8 类型
	// s := "你好 golang"
	// for _,v:= range s {  // 可以原样输出中文字符这种
	// 	fmt.Printf("%v(%c)",v,v)
	// }

	// // 修改字符串
	// s1 := " big"
	// byteStr := []byte(s1) // 把字符串转换成 byte 数组
	// byteStr[0] ='p'  // 然后进行修改
	// fmt.Println(string(byteStr))  // 最后再强制转换为字符串

	// s2 := "你好 golang"
	// runeStr := []rune(s2)
	// runeStr[0]= '我'
	// fmt.Println(string(runeStr))

	// var a int8 = 20
	// var b int16 = 40
	// fmt.Println(int16(a) + b)

	// var a float32 = 20
	// var b float64 = 40
	// fmt.Println(float64(a) + b)

	// Sprintf 拼接字符，还可以用来string转换

	/*

	   // Sprintf 可以将任何类型转换为字符串类型 string 类型
	   	var i int = 22
	   	str := fmt.Sprintf("%d",i)
	   	fmt.Printf("值：%v 类型 ： %T\n",str,str)

	   	var f float64 = 3.14
	   	str2 := fmt.Sprintf("%f",f)
	   	fmt.Printf("值：%v 类型 ： %T\n",str2, str2)

	   	var  t bool = true
	   	str3 := fmt.Sprintf("%t",t)
	   	fmt.Printf("值 ：  %v 类型： %T\n",str3,str3)

	   	var b byte = 'a'
	   	str4 := fmt.Sprintf("%c",b)  //%c 原样输出
	   	fmt.Printf("值： %v 类型 ： %T\n",str4,str4)
	*/

	// 第二种，通过 strconv 包进行转换，需要进入 strconv 包  "strconv"

	/*
	
	var i int = 23
	str := strconv.FormatInt(int64(i),10)
	fmt.Printf("值：%v 类型 ：%T\n",str,str)

	var f float32 = 29.8
	str2 := strconv.FormatFloat(float64(f),'f',2,64)
	fmt.Printf("值： %v ， 类型 %T ",str2,str2)

	var t bool = false    // 但是这个意义不大
	str3 := strconv.FormatBool(t)
	fmt.Printf("值：%v 类型 ：%T\n",str3,str3)


	a := 'a'
	str4 := strconv.FormatUint(uint64(a),10)  // 表示以10进制格式化输出
	fmt.Printf("值：%v 类型 ：%T\n",str4,str4)
	*/

	/*
	
	// 将 string 类型转换为 数值类型  首先是整型
	str := "12345"
	fmt.Printf("%v -- %T",str,str)

	// 还可以使用 strconv 包进行 ParseInt 转换
	num, _:= strconv.ParseInt(str,10,64)
	fmt.Printf("%v -- %T",num ,num)

	num2 , _ := strconv.ParseFloat(str,64) // 转换为 Float 类型
	fmt.Printf("%v  -- %T",num2,num2)

	*/

	// 算数运算符  + - * /
	var a = 2
	var b = 8
	fmt.Println(a+b)

	// 对于除法运算， 如果运算的都是整数，那么在运算的时候，会去掉小数部分,保留整数部分

	var c = 10 
	var d = 3
	fmt.Println(c/d)

	// 取余运算 ，余数 = 被除数 - （被除数/除数）*除数
	fmt.Println(c%d)
	fmt.Println(-10 % 3)   // -10 -(10/3)*3 









}
