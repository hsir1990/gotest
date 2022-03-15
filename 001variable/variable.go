/*
 * @Author: Hsir
 * @Date: 2022-02-24 16:03:32
 * @LastEditTime: 2022-03-15 09:57:04
 * @LastEditors: Do not edit
 * @Description: In User Settings Edit
 */
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

var (
	v2     = 3
	v3, v4 = "11", 22
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

//     const (
//         n1 = 100
//         n2
//         n3
//     )
// 上面示例中，常量n1、n2、n3的值都是100。

// iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

// 举个例子：

//     const (
//             n1 = iota //0
//             n2        //1
//             n3        //2
//             n4        //3
//         )

// const (
// 	n1 = iota //0
// 	n2 = 100  //100
// 	n3 = iota //2
// 	n4        //3
// )
// const n5 = iota //0

// const (
// 	n1 = iota //0
// 	n2        //1
// 	_
// 	n4        //3
// )

// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）

//     const (
//             _  = iota
//             KB = 1 << (10 * iota)
//             MB = 1 << (10 * iota)
//             GB = 1 << (10 * iota)
//             TB = 1 << (10 * iota)
//             PB = 1 << (10 * iota)
//         )
// 多个iota定义在一行

//     const (
//             a, b = iota + 1, iota + 2 //1,2
//             c, d                      //2,3
//             e, f                      //3,4
//         )

// 1.1.3. 复数
// complex64和complex128

// 复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位

func main() {
	var v1 = 12.2
	fmt.Println(v1)
	v5, v6 := 12, 12.3

	var name, sex = "pprof.cn", 1

	// func main() {
	// 	x, _ := foo()
	// 	_, y := foo()
	// 	fmt.Println("x=", x)
	// 	fmt.Println("y=", y)
	// }
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 (在Lua等编程语言里，匿名变量也被叫做哑元变量。)

	// v2 := 3 //重复定义出错
	var v7 int = 12
	var v8 float32
	v9 := .005
	v9 = 5.11e3

	fmt.Println(v2, v3, v4, v5, v6, v7, v8, v9, name, sex)

	var byt1 string = "nihao"
	var byt2 byte = 'a'
	//  byt2  = "a"//注意单引号
	var byt3 uint = 'a'
	var byt4 rune = '倍'
	var byt5 int32 = '倍'
	var byt6 int64 = '倍'
	var byt7 byte = '\n'
	var byt8 byte = '\n'
	// byt2 = byt3  //类型不匹配报错
	// byt3 = byt2//类型不匹配报错
	//输出对应的是utf8的编码，utf8中镶嵌这ascll
	fmt.Println(byt1, byt2, byt3, byt4, byt5, byt6, byt7, byt8) //nihao 97 97 20493 20493 20493 10
	var str1 int = 22269
	fmt.Printf("str1= %c \n", str1)
	str2 := 10 + 'a'
	fmt.Printf("str2 = %c \n", str2)
	//字符类型的本质是
	//储存：  字符-->对应码值-->二进制-->存储-->
	//读取: 二进制-->码值-->字符-->读取

	var b bool = true
	if b {
		// unsafe.Sizeof返回类型的大小
		fmt.Printf("b的类型=%T 大小=%v 值是=%v", b, unsafe.Sizeof(b), b)
	}

	str3 := "hell0"
	// str3[0] = 'a'//报错
	fmt.Println(str3, str3[0])

	// //只能重新切来改变
	// s := []rune(str3)
	// s[0] = 'G' 这样可以

	var str4 string = "abc \n asd"
	var str5 string = `sdfa
	sfadf
	asdf
	asdfas
	dfasdf
	asdf
	sa`
	str5 = "aa" + "bb" +
		"cc"
	fmt.Println(str4, str5)

	//初始值
	var str6 int     //0
	var str7 float32 //0
	var str8 bool    //false
	var str9 string  //""空
	fmt.Println(str6)
	fmt.Println(str7)
	fmt.Println(str8)
	fmt.Println(str9)

	//语法转换
	var n1 int = 10
	var n2 float64 = 10.1
	n2 = float64(n1)
	var n3 float64 = 5555555555.5
	n1 = int(n3) //会溢出但是不报错

	var n4 int32 = 12
	var n5 int8
	var n6 int8
	// n5 = 128 + int8(n4) //编译不通过
	n6 = 127 + int8(n4) //编译通过，凡是结果不是127+12，按溢出处理

	//1转换srting
	// var sdf string = "11"
	// sdf = n1

	fmt.Println(n2, n5, n6)

	//转成字符
	var n7 int = 10
	var n8 float64 = 12.2
	var n9 bool = true
	var myChar byte = 'a'
	var n11 string
	n11 = fmt.Sprintf("%d", n7)
	n11 = fmt.Sprintf("%f", n8)
	n11 = fmt.Sprintf("%t", n9)
	n11 = fmt.Sprintf("%c", myChar)

	n11 = strconv.FormatInt(int64(n7), 10)
	n11 = strconv.FormatFloat(n8, 'f', 10, 64)
	n11 = strconv.FormatBool(n9)
	n11 = strconv.FormatUint(uint64(myChar), 10)
	fmt.Printf("n11的值= %v,n11类型=%T，n11的类型的大小=%v", n11, n11, unsafe.Sizeof(n11)) //n11的值= 97,n11类型=string，n11的类型的大小=16

	// 字符转其他类型
	n11 = "12"
	var n12 int64
	var n13 int
	n12, _ = strconv.ParseInt(n11, 10, 64)
	n13 = int(n12) //需要在转换一下
	n11 = "12.2"
	var n14 float64
	n14, _ = strconv.ParseFloat(n11, 64)

	var n15 bool
	n15, _ = strconv.ParseBool(n11)
	fmt.Println(n12, n13, n14, n15)

}
