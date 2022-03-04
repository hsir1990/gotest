/*
 * @Author: Hsir
 * @Date: 2022-03-02 18:07:07
 * @LastEditTime: 2022-03-04 14:23:57
 * @LastEditors: Do not edit
 * @Description: In User Settings Edit
 */
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
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

	//17将字符串左右两边的指定的字符去掉
	//strings.Trim("! hello! ","!") //["hello"] 将左右两边的 ！和“ ”去掉
	str8 = strings.Trim("! h el! lo! ", " !") //h el! lo
	fmt.Printf("str8 = %q \n", str8)
	//%q该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

	//18将字符串左边指定的字符去掉：strings.TrimLeft("! hello! ","! ") //["hello"]//将左边！和 “ ”
	//19将字符串左边指定的字符去掉：strings.TrimRight("! hello! ","! ") //["hello"]//将左边！和 “ ”
	//20判断字符串是否以指定的字符串开头： strings.HasPrefix("ftp://192.168.10.1","ftp") //true

	var b3 bool = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b3=%v \n", b3)

	//21判断字符串是否以指定的字符串结束： string.HasSuffix("sdfjkl.jpg","jpg") //true

	//1. 日期 time的类型==time.Time，
	//值是==2022-03-03 22:11:13.5509589 +0800 CST m=+0.002122001
	now := time.Now() //time.Time
	fmt.Printf("time的类型==%T，值是==%v \n", now, now)
	//time的类型==time.Time，值是==2022-03-03 22:11:13.5509589 +0800 CST m=+0.002122001

	//2. 通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v \n", now.Year())
	fmt.Printf("月=%v \n", now.Month())
	fmt.Printf("月=%v \n", int(now.Month()))
	fmt.Printf("日=%v \n", now.Day())
	fmt.Printf("时=%v \n", now.Hour())
	fmt.Printf("分=%v \n", now.Minute())
	fmt.Printf("秒=%v \n", now.Second())

	//格式化日期时间
	//方式1：就是使用Printf 或者 SPrintf
	fmt.Printf("当1前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当2前年月日 %d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dataStr = %v \n", dateStr)

	//方式二 使用time.Format()方法完成
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	//"2006-01-02 15:04:05" 是固定的数字

	// 时间常量

	// 100 * time.Millisecond

	//纳秒 Nanosecond Duration = 1
	//微秒 Microsecond
	//毫秒 Millisecond
	//秒 Second
	//分 Minute
	//时 Hour
	i := 1
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 10)
		if i == 10 {
			break
		}
	}

	// 7.time的Unix 和UnixNano的方法
	fmt.Printf("unix时间戳=%v  unixnano时间戳=%v \n", now.Unix(), now.UnixNano())

	//内置函数
	//1. len用来求长度 比如 string array slice map channel

	//2.new 用来分配内存，主要用来分配 值类型 ，比如 int,float32,struct...返回的是指针
	num1 := 100
	fmt.Printf("num1的类型%T，num1的值=%v，num1的地址=%v \n", num1, num1, &num1)
	num2 := new(int) //  可以类比成 *int  等于是new出一个地址，然后，这个地址再指向一个值，*num2 = 100，
	//num2的类型%T =》 *int
	//num2的值 = 地址 0xc04204c098 (这个地址是系统分配的)
	//num2的地址%v = 地址 0xc04206a020 （这个地址是系统分配的）
	//num2指向的值 = 100
	*num2 = 100
	fmt.Printf("num2的类型 %T，num2的值 =%v, num2的地址%v \n,num2这个指针，指向的值=%v",
		num2, num2, &num2, *num2)

	//3.make 用来分配内存，主要用来分配引用类型，比如 channel,map,slice.

	//错误处理
	//go中可以抛出一个panic，然后在defer中通过recover 捕获这个异常，然后正常使用
	// 使用defer+recover 来捕获异常
	panicTest()

	// 用error.New("错误说明")自定义生成错误，并用panic这个内置对象，接受一个interface{}类型的值（也就是任何值了）
	//作为参数，可以接收error类型的变量，输出错误信息，并推出程序
	panicTest2()

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

func panicTest() {
	//使用defer+recover 来捕获异常
	defer func() {
		err := recover() //recover()内置函数，可以捕获异常
		if err != nil {
			fmt.Println("")
			fmt.Println("err=", err)
			fmt.Println("")
			fmt.Println("发送邮件给admin@zongheng.com")
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("")
	fmt.Println("res=", res)
	fmt.Println("")
}

//函数去读取以配置文件init.conf 的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func panicTest1(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New(("读取文件有误"))
	}
}

func panicTest2() {
	err := panicTest1("config1.ini")

	if err != nil {
		fmt.Println("panicTest1()执行。。。。")
		//如果读取文件发送错误，就输出这个错误
		panic(err)
	}
	fmt.Println("panicTest2()继续执行。。。。")
}
