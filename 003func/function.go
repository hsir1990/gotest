/*
 * @Author: Hsir
 * @Date: 2022-03-02 18:07:07
 * @LastEditTime: 2022-03-03 18:34:44
 * @LastEditors: Do not edit
 * @Description: In User Settings Edit
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//大写定义全局，整个系统有效
var Name string = "大松"

// Name2 := "大宋"
// 报错因为上面的东西相当于定于为
// var Name2 string
// Name2 = "大宋"

func main() {
	suffix := makeSuffix(".jpg")
	fmt.Println("suffix===>", suffix("nihoa.jp1g"))
	fmt.Println("suffix===>", suffix("nihao"))

	//在函数内当执行defer时，暂时不执行，会将defer后面的语句压入独立的栈中（defer）
	defer fmt.Println("defer最后执行")
	defer fmt.Println("defer后入栈的先执行")

	//1  len()计算的是字节的多少，汉字一般一个字占3字节
	var str1 string = "hello北"
	fmt.Println(len(str1)) //打印8

	//2 所有用到slice 切片 []rune(str)

	var str2 string = "qie你"
	var str3 []rune = []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符是==%c \n", str3[i]) //%c可以打印单个字符
	}

	//3字符串转整数
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转化错误", err)
	} else {
		fmt.Println("转换结果", n)
	}
	//4整数转字符串
	str4 := strconv.Itoa(123)
	fmt.Printf("str4=%v，str4=%T  \n", str4, str4)

	//5字符串转[]byte :
	var btye1 []byte = []byte("hello go")
	fmt.Printf("byte=%v \n", btye1) //byte=[104 101 108 108 111 32 103 111]

	//6  []btye转字符串
	str5 := string([]byte{97, 98, 99})
	fmt.Printf("[]byte的字符串 == %v \n", str5)

	//7   10进制转2，8，16   ，strconv.FormatInt(123,2) 返回对应字符串
	var str6 string = strconv.FormatInt(123, 2)
	fmt.Printf("str6的2进制 == %v \n", str6) //str6的2进制 == 1111011
	str6 = strconv.FormatInt(123, 16)
	fmt.Printf("str6的16进制 == %v \n", str6) //str6的16进制 == 7b

	//8查找字符串是否在指定的字符串中 strings.Contains("seefood", "foo") //true
	var b bool = strings.Contains("seafood", "foo")
	fmt.Printf("b=%v \n", b)

	//9统计一个字符串有几个指定的字串  strings.Count("cehed","e") //2
	var b1 int = strings.Count("adfd", "d")
	fmt.Printf("b1=%v \n", b1)

	//10不区分大小的字符串比较 strings.EqualFold("abc","Abc")  //true
	var b2 bool = strings.EqualFold("abc", "ABC")
	fmt.Printf("EqualFold==%v \n", b2)
	fmt.Println("aa == Aa结果", "aa" == "Aa") //false

	//11返回字符串第一次出现的index值，如果没有返回-1  strings.Index("ladfabc0","abc") //4
	index := strings.Index("ladfabc0", "abc")
	fmt.Printf("索引==%v \n", index)

	//12返回子串在字符串最后一个出现的index，如果没有返回 -1  strings.LastIndex("ladfabc0","abc") //4
	index = strings.LastIndex("ladfabc0", "abc")
	fmt.Printf("索引LastIndex==%v \n", index)

	//13 将指定的子串替换成另外一个子串 strings.replace("go go hello","go语言",n)  n可以指定你希望替换几个，  如果 n=-1 表示全部替换
	str7 := "go go hello"
	var str8 string = strings.Replace(str7, "go", "go语言", -1)
	fmt.Printf("str7=%v,替换后str8=%v \n", str7, str8)

	//14按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组  strings.Split("hello,word,ok", ",")
	var strArr []string = strings.Split("hello,word,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Println(i, "=", strArr[i])
	}
	fmt.Printf("strArr==%v \n", strArr)

	//15 将字符串的字幕进行大小写的转换， strings.ToLower("Go")  //go   strings.ToUpper("Go") //GO
	str8 = "golang Hello"

	str8 = strings.ToLower(str8)
	str8 = strings.ToUpper(str8)
	fmt.Printf("str8==%v \n", str8)

	//16将字符串左右两边的空格去掉 string.TrimSpace("  sdfad df sdf  ")
	str8 = strings.TrimSpace("  sdfad df sdf  ")
	fmt.Println("\"" + str8 + "\"")
}

//编写一个makeSuffix(suffix string) 可以接收一个文件后缀名（比如 .jpg），
// 并返回一个闭包，调用闭包，可以传入一个文件名如果有则直接返回，没有则添加.jpg
//string.HasSuffix,该函数可以判断某个字符串是否有指定的后缀   strings.HasSuffix("nihao.jpg",".jpg")
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
