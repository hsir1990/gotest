/*
 * @Author: Hsir
 * @Date: 2022-02-24 16:03:32
 * @LastEditTime: 2022-03-01 17:30:19
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

func main() {
	var v1 = 12.2
	fmt.Println(v1)
	v5, v6 := 12, 12.3

	// v2 := 3 //重复定义出错
	var v7 int = 12
	var v8 float32
	v9 := .005
	v9 = 5.11e3

	fmt.Println(v2, v3, v4, v5, v6, v7, v8, v9)

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
