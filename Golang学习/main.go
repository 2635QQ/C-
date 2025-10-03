package main

import (
	"fmt"
)

// 定义变量

// func getUserinfo() (string, int) {
// 	return "张三", 18
// }

// 定义一个方法
func test() bool {
	fmt.Println("test...")
	return true
}

// var g = "全局变量" //全局变量
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

	// // 算数运算符  + - * /
	// var a = 2
	// var b = 8
	// fmt.Println(a+b)

	// // 对于除法运算， 如果运算的都是整数，那么在运算的时候，会去掉小数部分,保留整数部分

	// var c = 10
	// var d = 3
	// fmt.Println(c/d)

	// // 取余运算 ，余数 = 被除数 - （被除数/除数）*除数
	// fmt.Println(c%d)
	// fmt.Println(-10 % 3)   // -10 -(-10/3)*3 结果-1

	// fmt.Println(10 % -3)
	// goLang 中的自增。自减运算符，只能单独使用；不能和赋值运算符一起用

	//正确的写法
	// var a = 12
	// a++
	// fmt.Println(a)

	//逻辑运算符
	// && AND 与运算符，如果两边的操作数都是true ，那就是true ,否则为false
	// // || OR 或运算符，如果两边的操作数有一个True , 则为true，否则为false
	// // ! 逻辑NOT 运算符，如果条件为true ,则为false ,

	// //逻辑与和逻辑或 短路
	// var a = 10
	// if a > 9 && test(){
	// 	fmt.Println("执行")
	// }

	// // 交换a 和 b 的值;
	// var a = 34
	// var b = 10
	// t:= a
	// a = b
	// b = t

	// fmt.Printf(" a   的值 %v， b 的值 %v",a,b)

	// // 交换 a  和  b ,要求不能使用中间变量
	// var a = 32
	// var b = 12
	// a = a+b   // a = 32 + 12
	// b = a-b   // b= 32+12 -12 = 32

	// a = a-b  // (32+12)-32 = 12 ;// 这样就实现了交换
	// fmt.Printf("a=%v b=%v",a,b)

	// //  如果还有100天放假，是 xx 个星期 零xx 天
	// var days = 100
	// var week = days/7
	// var day = days%7
	// fmt.Printf("days = %v  week = %v  day = %v ",days, week,day)

	// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：
	// C = ( F-32 ) 除以 1.8,摄氏温度 （°C），华氏温度（°F），请求出华氏温度对应的摄氏温度

	// var F float32 = 100  // 华氏温度  这里需要注意的是，因为有除法，所以这个类型是定义成float 类型
	// C := (F-32)/1.8
	// fmt.Printf("转换后的摄氏温度为：C %v",C)

	// var a = 5 // 101
	// var b = 2 // 010

	// fmt.Println("a&b= ", a&b) // 000

	// fmt.Println("a|b=", a|b) // 111 值7

	// fmt.Println("a^b=", a^b) // 异或，有些不一样，则为1

	// fmt.Println("a<<b=", a<<b) // 10100   左移两位，在后面补0   值20
	// // 这个左移，其实也是 5* 2 的2次方，左移n 位就是乘以2 的n 次方

	// fmt.Println("a>>b= ",a>>b)  // 5/2 的2次方  值1

	//  flag := true

	//  if flag {
	// 	fmt.Println("flag = true")
	//  }

	// age := 30  // 这个age  是当前区域内部的全局变量
	// if age > 20 {
	// 	fmt.Println("成年人",age)
	// }
	// fmt.Println(age)

	// // if 语句的另一种写法
	// if age := 34; age > 20{   // 而这个if 是当前区域的局部变量
	// 	fmt.Println("是成年人", age)
	// }

	// 练习一个成绩等级的；  成绩大于90 输出A， 小于90 大于75 输出B，否则输出C

	// var socores = 93
	// if socores > 90{
	// 	fmt.Println("成绩大于90， A")
	// } else if socores >75{
	// 	fmt.Println("成绩为B")
	// }else{
	// 	fmt.Println("C")
	// }

	// //  求两个数的最大值
	// var a = 34
	// var b = 24
	// var max int
	// if a>b {
	// 	max = a
	// 	fmt.Println("最大值为 max ",max)
	// }else {
	// 	max = b
	// 	fmt.Println("最大值为 max ",max)
	// }

	// // for 循环练习
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)

	// }
	// 写 for 循环的时候要注意死循环
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++   // 因为上面没有写结束语句，所以这个在里面写上结束语句
	// }

	// i := 1   // 无限循环模式下，也就是没有初始语句，没有判断条件，没有结束语句
	// for {
	// 	if i<= 10{
	// 		fmt.Println(i)
	// 	}else{
	// 		break   // 使用 break 跳出循环；
	// 	}
	// 	i++  // 执行完上面的if 语句之后，加 1
	// }
	// // 且 Go 中没有 while 语句，可以使用for 循环代替，
	// // for 条件{ 循环体语句}

	// // 1、练习 ： 打印0-50 所有的偶数
	//  i := 0
	//  for ; i<=50 ;i++{
	// 	if i%2==0{
	// 		fmt.Println(i)
	// 	}
	//  }

	// // 2、练习；求 1+2+3+4....+100 的和
	// i := 1
	// sum :=0;  // sum 定义在外面
	// for ; i<=100;i++{
	// 	sum +=i;
	// }
	// fmt.Println("总和为%v :",sum)  // 可以放到外面，那么打印就是打印一次了

	// // 3、练习： 打印 1~100 之间所有是 9 的倍数的整数的个数及总和
	// i := 1
	// count := 0
	// sum := 0
	// for ; i<= 100 ;i++{
	// 	if i%9 == 0{
	// 		count +=1;
	// 		sum += i;
	// 	}
	// }
	// fmt.Printf("之间是9的倍数的整数个数有 %v 个, 其整数之和为%v,",count,sum)

	// // 4、 计算5 的阶乘
	// i:=1
	// res := 1
	// for ; i<=5;i++{

	// 	res *= i
	// }
	// fmt.Println(res)

	// // 打印一个矩形 下面是一种方法 ； 但还可以用for循环的嵌套；
	// for i := 1; i<= 12; i++{
	// 	fmt.Printf("*")
	// 	if i% 4== 0{
	// 		fmt.Println("")  // 是4的倍数的时候，打印一个换行
	// 	}
	// }

	// // for 循环的嵌套

	// var row = 3
	// var column = 3

	// for i := 0; i< row; i++{  // 这个是行
	// 	for j := 0; j < column;j++{   // 这个是列
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// // for 循环打印一个三角形
	// var row = 5
	// for i:=0;i<row ; i++{
	// 	for j := 1;j<= i;j++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")  // 控制打印5行
	// }

	// // 练习： 打印出9*9 乘法表
	// var row =9

	// for i:=1;i<=row;i++{
	// 	for j:=1;j<=i;j++{
	// 		fmt.Printf("%v*%v = %v ",j,i,i*j)
	// 	}
	// 	fmt.Println("")
	// }

	// // 打印字符串里面的内容，依次输出
	// var str ="你好，golang!"
	// for k,v := range str {
	// 	// fmt.Println("key= ",k,"val=",v)
	// 	fmt.Printf("key= %v, val= %c\n",k,v)  //要以字符输出，也就是%c

	// }
	// var arr = []string{"ahp","java","nodejs","golang"}
	// for i:=0;i<len(arr);i++{
	// 	fmt.Println(arr[i])
	// }
	// var arr = []string{"ahp","java","nodejs","golang"}
	// for k , val := range arr{
	// 	fmt.Println(val)
	// 	fmt.Println("key=",k)
	// }

	// 练习： 判断文件类型，如果后缀名是.html ,输入 text/html,
	// 如果后缀名 .css 输出 text/css ，
	// var extname = ".html"
	// switch extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// case "js":
	// 	fmt.Println("text/javascript")
	//     break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }

	// // 同时，这个变量的定义还可以放在 switch 里面进行定义
	// switch extname := ".html"; extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// }

	// //一个分支可以有多个值，多个case 值中间使用英文逗号分隔
	// var n = 5
	// switch n {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("这是奇数")
	// 	break
	// case 2, 4, 6, 8:
	// 	fmt.Println("偶数")
	// 	break
	// }

	// // case 使用分支语句进行 判断
	// var age = 30
	// switch {
	// case age < 24:   // 如果分支上使用的是表达式，那么switch后面不需要跟变量
	// 	fmt.Println("好好学习")
	// case age >= 24 && age <= 68:
	// 	fmt.Println("好好赚钱")
	// default :
	// fmt.Println("输入错误")
	// }

	// 
	// var age = 30
	// switch {
	// case age < 24:
	// 	fmt.Println("好好学习")
	// case age >= 24 && age <= 60:
	// 	fmt.Println("好好赚钱")
	// 	fallthrough
	// case age >60:
	// 	fmt.Println("注意身体")
	// default:
	// 	fmt.Println("输入错误")
	// }

	// // 1、表示当i = 2 的时候跳出当前循环
	// for i:= 1;i<=10;i++{
	// 	if i==2{
	// 		break  // 这个break 只是跳出当前循环
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("继续执行下一步")

	// 
	// for i:=0;i<2;i++{
	// 	for j:= 0;j<10;j++{
	// 		if j == 3{
	// 			break  // 其实是会跳出当前循环，这里层的for 循环不会被打印
	// 		}
	// 		fmt.Printf("i=%v j=%v\n",i,j)
	// 	}
	// 	// 直接跳出循环之后，进入到下一步，也就是执行下一轮的循环，因为这个里面没有其他句子了
	// }

	// label:
	// for i:=0;i<2;i++{
	// 	for j:= 0;j<10;j++{
	// 		if j == 3{
	// 			break label // 其实是会跳出当前循环，这里层的for 循环不会被打印
	// 		}
	// 		fmt.Printf("i=%v j=%v\n",i,j)
	// 	}
	// 	// 直接跳出循环之后，进入到下一步，也就是执行下一轮的循环，因为这个里面没有其他句子了
	// }

	// continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for 循环内使用
	// for i:=1;i<= 10;i++{
	// 	if i==3{
	// 		continue           //break  的时候，到 1 2 ；跳出当前循环
	// 	      //  而这个 continue 语句跳过当前循环，然后继续向下一轮的循环执行；所以没有打印3
	// 		  //  在跳过3之后，执行下一轮的需要执行的语句
	// 	}
	// 	fmt.Println(i)
	// }

	// continue 和 label 标签的一个配合使用
	// label2:
	// for i:=0;i<2; i++{
	// 	for j:=0;j< 10;j++{
	// 		if j==3 {
	// 			continue label2
	// 		}
	// 		fmt.Printf("i=%v j=%v\n",i,j)
	// 	}
	// }
// 数组
	// var arr1 [3]int 

	// var arr2 [4]int

	// var strArr [3]string

	// // 查看一下数组的类型
	// fmt.Printf("arr1:%T  arr2%T  atrArr:%T",arr1,arr2,strArr)

	// 数组的初始化 ：第一种方法； 声明数组
	// var arr1 [3]int
	// fmt.Println(arr1)

	// arr1[0]= 22
	// arr1[1] = 10
	// arr1[2]= 24

	// fmt.Println(arr1)

	// 数组的初始化 第二种方法：
	// var arr1 = [3]int{23,34,5}
	// fmt.Println(arr1)

	// // 让其自行推断长度
	// var arr1= [...]int{12,4354,656,333333336,
	// }
	// fmt.Println(len(arr1))


	// // 第四种方式 指定下标给数组，索引值的方式，用到的可能不是特别多
	// arr := [...]int {0:1,1:10,2:27,5:56}
	// fmt.Println(len(arr))

	// fmt.Println(arr)

	// // 数组的循环遍历  for for range
	// var arr = [3]int {23,34,5}
	// for i:=0;i<len(arr);i++{
	// 	fmt.Println(arr[i])
	// }

	// arr1 := [...]string{"php","nodejs","golang","js"}
	// for k,v := range arr1{
	// 	fmt.Printf("key:%v value:%v\n",k,v)

	// }

	// // 1、 求出一个数组里面元素的和以及这些元素的平均值，分别用for 和 for -range 实现
	// var arr = [...] int {32 ,12,45,5}
	// var sum = 0 
	// for i:=0;i< len(arr);i++{
	// 	sum += arr[i]
	// }
	// // fmt.Println("数组的和 ",sum)
	// fmt.Printf("arr 数组元素的和是：%v 平均值：%.2f",sum,float64(sum)/float64(len(arr)))

	// 请求出一个数组的最大值 ，并得到对应的下标

	// var arr = [...]int {1,12,25,65,11}

	// var max = arr[0]
	// var index = 0

	// for i:=0;i<len(arr);i++{
	// 	if arr[i]>max{
	// 		max=arr[i]
	// 		index = i
	// 	}
	// }
	// fmt.Printf("最大值：%v , 最大值对应的索引值%v ",max,index)

	// // 从数组中 找出和为8 的两个元素的下标；
	// var arr = [...]int{1,3,5,7,8}
	// for i:=0;i<len(arr);i++{
	// 	for j:=i+1;j<len(arr);j++{
	// 		if arr[i] + arr[j] == 8{
	// 			fmt.Printf("(%v,%v)",i,j)
	// 		}
	// 	}
	// }

	//数组的值类型和引用类型; 下面这种能够将arr1 拷贝给arr2 的是数值类型
	// var arr1 = [...]int{1,2,4}
	// arr2 := arr1
	// fmt.Println(arr1,arr2)
	// arr1[0]=13
	// fmt.Println(arr1,arr2)

	//引用类型： 切片 定义的时候 【】 里面是空的表示不指定数组的长度
	// var arr1 = []int {1,2,3}
	// arr2 := arr1
	// arr1[0]=11
	// fmt.Println(arr1,arr2)

	// 多维数组 var arr = [3]int{1,2,3}  一维数组
	var arr = [3][2]string{
		{"北京","上海"},
		{"广州","深圳"},
		{"成都","重庆"},
	}
	fmt.Println(arr[0])


	// 循环遍历数组
	for _, v1:=range arr{
		for _,v2 := range v1{
			fmt.Println(v2)
		}
	}

	// 或者使用两层for 循环
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr[i]);j++{
			fmt.Println(arr[i][j])
		}
	}

	// 支持数组推导的方式，
	var arr1 = [...][2]string{
		{"北京","上海"},
		{"广州","深圳"},
	}
	fmt.Println(arr1)







}
