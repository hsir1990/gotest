package main

import (
	"fmt"
	utils "gotest/util"
	"math/rand"
	"time"
)

var (
	//定义一个全局的匿名函数
	fun6 = func(n int, n2 int) int {
		return n + n2
	}
)

// . go中生成随机数的有两个包,分别是“math/rand”和“crypto/rand”,
// 前者实现了伪随机数生成器,
// 后者实现了用于加解密的跟安全的随机数生成器,当然,性能也就降下来了,毕竟鱼与熊掌不可兼得

func main() {
	// var i  *int   定义指针类型，
	// *i 获取i的指针的值
	//&a  获取a指针的地址
	//基本数据类型也就值类型 ， 包含 int系列  float系列 string Boolean  结构体struct 数组

	//值储存一般放到  栈中
	//引用使用的是地址，一般放在  堆中， 没有引用以后就被回收了
	// var f-d  int  变量名字不能这样定义  也不能用数字打头  3ii
	//可以这样定义  _ss  下划线可用  ss_ss

	var int1 int = 1
	int1++ //只能这么写，只能写在后面，   而且要单独使用，前面不能放等号，也不能直接比较  int1++ > 3  这样写不对

	var float1 float32 = 10 / 2
	fmt.Println(float1)
	var int2 int = 10 / 2 //省略小数点后面的
	fmt.Println(int2)

	//  <<=   ^=  |=  &=  >>=  这些复制运算符也可以使用

	//1即时真，0为假

	// & 同时为1，结果为1，否则为0
	// | 有一个为1，结果为1，否则为0
	//^ 当二进位不同时结果为1 ，相同是0
	//<< 二进制最低位补0， 乘以2，最高位省略
	//>> 二进制最高位补0，舍去最低位，相当于除2，符号位不变

	// var int3 int
	// fmt.Scanln(&int3)
	// fmt.Println("int3=", int3)

	// var bool1 bool
	// var float2 float64
	// var age byte
	// var name string
	// fmt.Scanf("%t %f %d %s", &bool1, &float2, &age, &name)

	//二进制
	var int4 int = 5
	fmt.Printf(" %b\n", int4)
	//8进制
	int4 = 031 //0开头的表示8进制
	int4 = 0x2 //0X开头的表示16进制

	int4 = -1 << 2
	fmt.Println("int4 << 2=", int4) //= -4
	int4 = -1 >> 2
	fmt.Println("int4 >> 2=", int4) //-1 //可能是整数的原因
	int4 = 1 >> 2
	fmt.Println("int4 >> 2=", int4) //1 //可能是整数的原因

	//不能是小数，必须是整数
	// var float2 float64 = -1.1
	// float2 = float2 >> 2  //无效操作：移位操作数float2（float64类型的变量）必须为整数
	// fmt.Println("float2 >> 2=", float2) //-1 //可能是整数的原因

	//位运算的时候
	//补数进行计算，计算完成以后，需要对补数进行 --反补--取反--得到正解

	if true {
		fmt.Println("你好")
	}

	var year int = 2019
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("this is run nian")
	} else {
		fmt.Println("this is not run nian")
	}

	var int5 int8 = 2
	switch int5 {
	case 5, 7:
		fmt.Println("true5")
	case 2:
		fmt.Println("true2")
		fallthrough //只穿透一层
	case 3:
		fmt.Println("true3")
	default:
		fmt.Println("other")
	} //true2   true3

	// 判断某个interface变量实际指向的变量类型
	var inter1 interface{}
	int6 := 10.0
	inter1 = int6
	switch i := inter1.(type) {
	case nil:
		fmt.Printf("x的类型~:%T \n", i)
	case int:
		fmt.Println("x的类型是:int")
	case float64:
		fmt.Println("x的类型是:float64")
	case func(int) float64:
		fmt.Println("x的类型是:func(int)")
	case bool, string:
		fmt.Println("x的类型是:bool或string")
	default:
		fmt.Println("x的类型是未知")

	}
	//具体数值不多，而且符合整形，浮点数，字符，字符串这几个类型 可以用switch
	//对于区间判断和结果为bool类型的判断，使用if

	//遍历字符串
	str1 := "你好abc~ok"
	for key, val := range str1 {
		fmt.Printf("index=%d,val=%c \n", key, val)
		//打印出下面的东西   //%c打印出字节
		// index=0,val=a
		// index=1,val=b
		// index=2,val=c
		// index=3,val=~
		// index=4,val=o
		// index=5,val=k

	}

	//将汉字字符串分解
	var str2 string = "你好aa"
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		// fmt.Println(str3[i]) //打印出来的是数字
		fmt.Printf("%c \n", str3[i]) //打印出来的是字符
	}

	//生成1-100的随机数
	for {
		// time.Now.Unix():返回从1979年到现在的秒数
		rand.Seed(time.Now().UnixNano()) //因为是伪随机数，我们需要加一个种子
		// rand.Intn(100)  [0,n)的伪随机int值
		rand1 := rand.Intn(100) + 1

		fmt.Println("rand1=", rand1)
		if rand1 != 10 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("continue=", i)
	}

	//goto
	for i := 0; i < 5; i++ {
		if i == 3 {
			goto label1
			fmt.Println("nihao")
		}
		fmt.Println(i)
	label1:
		fmt.Println("lable1=", i)

	}

	utils.Bao(1)

	res1, res2 := getSumAndSub(5, 3)
	fmt.Println(res1)
	fmt.Println(res2)
	res3 := recursion(5)
	fmt.Println(res3)

	//斐波纳契数
	fmt.Println("fbn(5)=", fbn(5))
	//猴子吃桃子
	fmt.Println("fbn(6)=", peach(1)) // 1534

	// 将值拷贝变成引用拷贝
	var int3 int = 0
	test(&int3)
	fmt.Println("引用拷贝0===", int3)

	//函数是一个类型，也可以作为参数
	func4 := func1(func3, 2, 3)
	fmt.Printf("func4类型==%T，fun1的类型==%T \n", func4, func1)
	// 匿名函数func5===
	func5 := func1(func(num1 int, num2 int) int {
		return num1 - num2
	}, 2, 3)

	fmt.Println("匿名函数func5===", func5)

	// 自定义函数类型
	func7 := func6(func(num1 int, num2 int) int {
		return num1 - num2
	}, 2, 3)

	fmt.Println("匿名函数func7===", func7)

	// go支持自定义数据类型,相当于给一个别名
	type myInt int
	var myInt1 myInt = 1
	fmt.Printf("myInt类型==%T,myInt值是==%d \n", myInt1, myInt1)
	var int9 int = 1
	if int9 == int(myInt1) {
		fmt.Printf("int虽然自定义了int类型，不过相当于生成了新的类型，需要转义才能比较")
	}

	//可变参数的使用
	sum1 := sum(1, 2, 3)
	fmt.Println("可变参数的验证===", sum1)

	//测试其中一个函数不定义类型，另一个需要定义,竟然可以通过
	sum2 := sum5(2, 3)
	fmt.Println("测试其中一个函数不定义类型，另一个需要定义sum2===", sum2)

	//当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。

	//每个源文件都可以包含一个init的函数

	//执行流程  全局函数 ===》 init函数===》main主函数

	//多文件执行是   util的全局变量==》util的init的函数 ===》main的全局变量==》main的init函数==》main的主函数

	//全局匿名函数的使用
	fun7 := fun6(2, 3)
	fmt.Println("全局匿名函数的使用fun7===", fun7)

	// 闭包累加器
	add := addUpper()
	fmt.Println("add(2)累加--", add(2))
	fmt.Println("add(3)累加--", add(3))

}

//返回多参数，  这个函数体不能写到main的主函数体中
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

//递归
func recursion(arg int) int {
	if arg == 1 {
		return arg
	}
	return arg * recursion(arg-1)
}

//斐波纳契数 1,1,2,3,5...
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

//猴子吃桃子，  第一天吃一半，并多吃一个，
//             第二天吃一半，并多吃一个，
//				当第10天，发现只有一个桃子，问最开始有多少个桃子
func peach(n int) int {
	if n == 10 {
		return 1
	}
	return (peach(n+1) + 1) * 2
}

//引用拷贝
func test(n *int) {
	*n = *n + 2
}

//函数作为参数
func func1(func2 func(int, int) int, num1 int, num2 int) int {
	return func2(num1, num2)
}
func func3(num1 int, num2 int) int {
	return num1 + num2
}

//函数重新定义
type myFun func(int, int) int

func func6(func2 myFun, num1 int, num2 int) int {
	return func2(num1, num2)
}

//go支持可变参数

func sum(n1 int, arg ...int) (sun int) { //这里面的arg是slice的切片//,同时后面的返回参数定义时需要加括号，也要加return 关键字
	sun += n1
	for i := 0; i < len(arg); i++ {
		sun += arg[i]
	}
	return
}

func sum5(n1, n2 float32) float32 {

	//调用的时候默认是float32了
	fmt.Printf("n1的类型---%T 、n", n1)
	return n1 + n2
}

// 闭包
func addUpper() func(int) int { //注意要返回一个函数
	i := 1
	return func(n int) int {
		i += n
		return i
	}
}

// &（取地址）和*（根据地址取值）

//指针类型，如：*int、*int64、*string等
// ptr := &v    // v的类型为T

// v:代表被取地址的变量，类型为T
//     ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。

//需要测试
// func main() {
//     var a *int
//     *a = 100 //应该是个地址，需要new分配内存
//     fmt.Println(*a)

//     var b map[string]int //需要定义make分配内存
//     b["测试"] = 100
//     fmt.Println(b)
// }
// 执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存

//new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
// func new(Type) *Type
// 其中，

//     1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
//     2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

// a := new(int)
//     b := new(bool)
//     fmt.Printf("%T\n", a) // *int
//     fmt.Printf("%T\n", b) // *bool
//     fmt.Println(*a)       // 0
//     fmt.Println(*b)       // false

// var a *int只是声明了一个指针变量a但是没有初始化

// 应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：
// var a *int
//     a = new(int)
//     *a = 10
//     fmt.Println(*a)

// make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
// b = make(map[string]int, 10)

// 1.二者都是用来做内存分配的。
//     2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//     3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

// var a int
// fmt.Println(&a)
// var p *int
// p = &a
// *p = 20
// fmt.Println(a)
