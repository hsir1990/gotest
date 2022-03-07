/*
 * @Author: Hsir
 * @Date: 2022-03-04 14:24:55
 * @LastEditTime: 2022-03-07 10:26:54
 * @LastEditors: Do not edit
 * @Description: In User Settings Edit
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	//定义数组
	var arr [5]float64
	//给值
	arr[0] = 1.023
	arr[1] = 5.0
	arr[2] = 3.0
	arr[3] = 9.0
	arr[4] = 7.0
	fmt.Printf(" arr=%v, arr[1]= %.2f\n, &arr=%p, &arr[1]= %p \n", arr, arr[0], &arr, &arr[0]) //%p是地址
	//1.上面数组地址获取可以通过  &arr  获取
	//2.数组第一个元素的地址，就是数组的首地址  &arr==&arr[1]
	//3数组的个元素之间间隔地址是依据数组的类型决定的   比如  int64 间隔 8位，  int32 间隔 4位

	// 4  种初始化数组的方式
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{1, 2, 3}
	// 这种写法也规定数组的写法[...]
	var arr3 = [...]int{1, 2, 3}
	var arr4 = [...]int{1: 1, 0: 5, 2: 7}

	//类型推导
	arrStrings := [...]string{1: "dd", 0: "aa", 2: "cc"}
	arrStrings2 := [...]string{"dd", "aa", "cc"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arrStrings)
	fmt.Println(arrStrings2)

	//1.数据一旦定义，长度和类型就不能改变 ，不能越界，不能改变类型
	//2.同时数组种的类型可以是引用或者数值类型，但是不能混用
	//3.数组创建后，如果没有赋值，会有默认值
	//4.go的数组属于值传递，属于值拷贝，如果想改变原来数组的值，需要用引用传递（指针的方式）

	//切片属于引用类型  slice

	var intArr [5]int = [5]int{1, 22, 3, 33, 77}
	//1.声明/定义一个切片
	//intArr[1:3]表示slice 引用到intArr这个数组，   引用intArr数组的起始下标为1，最后的下标为3(但是不包含3)
	slice1 := intArr[1:3]                                                                  //intArr[0:4]可以简写成intArr[:]
	fmt.Printf("slice1=%v,slice1的个数=%d,slice1的容量=%d \n", slice1, len(slice1), cap(slice1)) //切片的容量可以动态变化
	//slice=[22 3],slice的个数=2,slice的容量=4
	//这种方式原来的数组会被引用，可以修改原来的数组
	slice1[0] = 66
	fmt.Println("被改变的intArr [5]int", intArr) // [1 66 3 33 77]

	//slice从底层来说，其实就是一个数据结构（struct结构体）
	type slice struct {
		ptr *[2]int
		len int
		cap int
	}

	//2 第二种方式，用make来创建切片
	//var 切片名 []type = make([]type,len,[cap])  /如果分配了cap  则 cap>=len
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[2] = 20.0
	fmt.Printf("slice2=%v,slice2的个数=%d,slice2的容量=%d \n", slice2, len(slice2), cap(slice2))
	//slice2=[0 10 20 0 0],slice2的个数=5,slice2的容量=10
	//方法二与方法一的区别在于 方法一程序员事先可见，方法二由底层进行维护,不可见

	//3 第三种方式，直接就指定具体数组，使用原理和make的方式相似
	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Printf("slice3=%v,slice3的个数=%d,slice3的容量=%d \n", slice3, len(slice3), cap(slice3))
	//slice3=[tom jack mary],slice3的个数=3,slice3的容量=3

	//注意切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片使用

	//使用append内置函数，对切片进行动态追加
	slice3 = append(slice3, "tos", "abc")
	slice2 = append(slice2, slice2...)
	fmt.Println(slice3)
	fmt.Println(slice2)
	//append底层操作是 创建一个新的切片，将slice3的值拷贝到新的切片上，然后可以在复制到slice3上，等于是生成了新的切片

	//copy内置函数完成拷贝  copy(slice5, slice4)两个参数都是切片，copy之后两个值都互不影响
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 []int = make([]int, 5, 10)
	copy(slice5, slice4)
	fmt.Println(slice4) // [1 2 3 4 5]
	fmt.Println(slice5) //[1 2 3 4 5]

	//string 底层是一个byte数组，因此也可以进行切片处理
	var str string = "abcdef"
	strSlice := str[2:]
	fmt.Println("strSlice=", strSlice)          //strSlice= cdef
	fmt.Printf("strSlice的切片类型=%T", strSlice[1]) //strSlice的切片类型=uint8

	//string 是不可变的，  var str string = "abcdef"  str[0]='z'//报错

	//如果修改字符串，可以先将string==》[]byte 或者 []rune  ==》修改==》重写转成string
	var arr5 []uint8 = []uint8(str)
	arr5[1] = 's'
	fmt.Println("str1=", str) //str1= abcdef  原来的没有变

	str = string(arr5)
	fmt.Println("str=", str) //str= ascdef

	var arr6 []rune = []rune(str)
	arr6[1] = '南'
	fmt.Println("str1=", str) //str1= ascdef 原来的没有变

	str = string(arr6)
	fmt.Println("str=", str) //str= a南cdef

	var int7 [5]int = [5]int{11, 3, 7, 9, 7}
	BubbleSort(&int7)

	//定义声明二维数组
	var arr7 [3][6]int
	arr7[1][2] = 1
	arr7[2][3] = 2
	fmt.Println("arr7==", arr7)

	//另外一种初始化二位数组
	var arr8 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr8=", arr8)

	//4种方式
	// var arr8 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	// var arr8 [2][3]int = [...][3]int{{1,2,3},{4,5,6}}
	// var arr8 = [2][3]int{{1,2,3},{4,5,6}}
	// var arr8 = [...][3]int{{1,2,3},{4,5,6}}

	//map 是key-value的数据结构，又称为字段或者关联数组
	//其中key可以很多种类型，比如 bool ，数字，string，指针，channel ，
	//还可以是只包含前面几个类型的接口，结构体，数组
	//通常为 int 和 string
	// 注意 slice ，map还有function不可以，因为这几个没法用 == 来判断
	//value和key能用的类型基本上一样

	// map声明的举例
	//var a map[int]string
	//var a map[string]map[string]int

	// 声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
	//方式一
	var a map[string]string
	a = make(map[string]string, 10) //make的作用是给map分配数据空间
	a["no1"] = "松江"
	a["no2"] = "武松"
	a["no1"] = "无用" //不能重复，出现重复，会以最后一个为主
	a["no3"] = "鲁智深"
	fmt.Println(a)
	//map中没有顺序可言
	//方式二
	map1 := make(map[string]string)
	map1["no1"] = "北京"
	map1["no2"] = "天津"
	map1["no3"] = "上海"
	fmt.Println(map1)
	//方式三
	var map2 map[string]string = map[string]string{
		"no1": "松江",
		"no2": "武松",
		"no3": "鲁智深",
	}
	//可添加
	map2["no4"] = "无用"
	fmt.Println(map2)

	//map的排序
	map_2 := make(map[int]int, 10)
	map_2[10] = 100
	map_2[1] = 10
	map_2[3] = 2
	map_2[7] = 55
	fmt.Println(map_2)

	// 通过切片来排序
	// var int_2 []int = make([]int, 2)
	var int_2 []int
	for key, _ := range map_2 {
		int_2 = append(int_2, key)
	}
	sort.Ints(int_2)
	fmt.Println(int_2)

	for _, value := range int_2 {
		fmt.Printf("map_2[%v] =%v \n", value, map_2[value])
	}

	//map是引用类型，可以自动扩容，里面也可以是结构体
	//需要先定义Stu
	type Stu struct {
		Name    string
		age     int
		address string
	}
	students := make(map[string]Stu, 10)
	str1 := Stu{"tom", 18, "上海"}
	str2 := Stu{"to1m", 18, "上海"}
	students["no1"] = str1
	students["no2"] = str2
	fmt.Println(students)
}

//冒泡
func BubbleSort(arr *[5]int) {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := i; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp

			}
		}
	}
	fmt.Println(*arr)
}
