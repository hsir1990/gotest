/*
 * @Author: Hsir
 * @Date: 2022-03-04 14:24:55
 * @LastEditTime: 2022-03-15 14:38:01
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
	fmt.Println(slice5) //[5 10 3 4 5]

	// slice s1 : [8 9]
	// slice s2 : [0 1 2 3 4]

	// copy(s2, s1)
	// copied slice s1 : [8 9]
	// copied slice s2 : [8 9 2 3 4]
	//应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。

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
	fmt.Println(map2) //map[no1:松江 no2:武松 no3:鲁智深 no4:无用]

	//delete(map,"key")   delete 是一个内置函数，如果key 存在，就删除key-value，
	//如果key不存在，不操作，但是也不会报错

	delete(map2, "no4")
	fmt.Println(map2)
	//没有专门的一个方法一次删除所有的，可以遍历一下key去逐个删除，
	//或者make一个新的，将原来的成为垃圾

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

	//例子   map 中的值是 map
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["age"] = "18"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "tom1"
	studentMap["stu02"]["age"] = "19"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu01"])
	fmt.Println(studentMap["stu01"]["name"])

	//map遍历只能使用for rang

	//map的长度
	fmt.Println(len(studentMap))

	//map切片  slice of map
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "15"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "牛魔王1"
		monsters[1]["age"] = "16"
	}

	newMonter := map[string]string{
		"name": "牛魔王2",
		"age":  "17",
	}

	monsters = append(monsters, newMonter)

	fmt.Println(monsters)

	//map每次输出都是无序的
	//不过int的key，可以通过sort.Ints(keys) 去搞
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

// 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
// 8.指针数组 [n]*T，数组指针 *[n]T。
//   7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。

// 也可以这样
// var str = [5]string{3: "hello world", 4: "tom"}
// var arr2 = [...]int{1, 2, 3, 4, 5, 6}// 使用索引号初始化元素。

// d := [...]struct {
// 	name string
// 	age  uint8
// }{
// 	{"user1", 10}, // 可省略元素类型。
// 	{"user2", 20}, // 别忘了最后一行的逗号。
// }

// b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。

// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

//  // 若想做一个真正的随机数，要种子
//     // seed()种子默认是1
//     //rand.Seed(1)
//     rand.Seed(time.Now().Unix())
// b[i] = rand.Intn(1000)

// 需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案

// 、、如果 slice == nil，那么 len、cap 结果都等于 0。

// c := append(a, b...)

// arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
//         fmt.Println("原数组：",arr)
// append ：向 slice 尾部添加数据，返回新的 slice 对象。

//         fmt.Println("对数组进行截取：")
//         //如果指定max，max的值最大不能超过截取对象（数组、切片）的容量
//         s1:=arr[2:5:9] //max:9  low：2  high;5  len:5-2(len=high-low)  cap:9-2(cap=max-low)
//         fmt.Printf("数组截取之后的类型为：%T,    数据是：%v;长度：%d;容量：%d\n",s1,s1, len(s1), cap(s1))

//         //如果没有指定max，max的值为截取对象（切片、数组）的容量
//         s2:=s1[1:7]  //max:7  low：1  high;7  len:7-1(len=high-low)  cap:7-1(cap=max-low)
//         fmt.Println("对切片进行截取：")
//         fmt.Printf("对切片进行截取之后的数据是：%v,长度:%d； 容量%d\n",s2, len(s2), cap(s2))

// arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
//         fmt.Println("原数组：",arr)

//         fmt.Println("对数组进行截取：")
//         //如果指定max，max的值最大不能超过截取对象（数组、切片）的容量
//         s1:=arr[2:5:9] //max:9  low：2  high;5  len:5-2(len=high-low)  cap:9-2(cap=max-low)
//         fmt.Printf("数组截取之后的类型为：%T,    数据是：%v;长度：%d;容量：%d\n",s1,s1, len(s1), cap(s1))
//    //int[]  [3 4 5]  3   9
//         //如果没有指定max，max的值为截取对象（切片、数组）的容量
//         s2:=s1[1:7]  //max:7  low：1  high;7  len:7-1(len=high-low)  cap:7-1(cap=max-low)
//         fmt.Println("对切片进行截取：")
//         fmt.Printf("对切片进行截取之后的数据是：%v,长度:%d； 容量%d\n",s2, len(s2), cap(s2))
//  // int[]   [4 5 6 7 8 9]  6  6

//         //利用数组创建切片，切片操作的是同一个底层数组
//         s1[0]=8888
//         s2[0]=6666
//         fmt.Println("操作之后的数组为：",arr)
// 		// [1 2 8888 6666  5 6 7 8 9 10]
//         /*
//         切片对数组的截取  最终都是切片操作的底层数组（通过指针操作原数组）
//         */

// data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8

// str := "hello world"
// s1 := str[0:5]
// fmt.Println(s1) //hello

// s2 := str[6:]
// fmt.Println(s2) //world

// 数组or切片转字符串：
// strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)

// 切片本身并不是动态数组或者数组指针。它内部实现的数据结构通过指针引用底层数组，设定相关属性将数据读写操作限定在指定的区域内。切片本身是一个只读对象，其工作机制类似数组指针的一种封装

// 切片（slice）是对数组一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。切片提供了一个与指向数组的动态窗口。

// 并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，而且小数组在栈上拷贝的消耗也未必比 make 消耗大。
//分情况使用数组和切片

// type slice struct {
//     array unsafe.Pointer   //定义任意类型的指针  //一般都自己定义好 *[]int
//     len   int
//     cap   int
// }

// unsafe.Pointer
// type Pointer
// type Pointer *ArbitraryType
// Pointer类型用于表示任意类型的指针。有4个特殊的只能用于Pointer类型的操作：

// 1) 任意类型的指针可以转换为一个Pointer类型值
// 2) 一个Pointer类型值可以转换为任意类型的指针
// 3) 一个uintptr类型值可以转换为一个Pointer类型值
// 4) 一个Pointer类型值可以转换为一个uintptr类型值
// 因此，Pointer类型允许程序绕过类型系统读写任意内存。使用它时必须谨慎。

// 如果想从 slice 中得到一块内存地址，可以这样做：

// s := make([]byte, 200)
// ptr := unsafe.Pointer(&s[0])
// 如果反过来呢？从 Go 的内存地址中构造一个 slice。

// var ptr unsafe.Pointer
// var s1 = struct {
//     addr uintptr
//     len int
//     cap int
// }{ptr, length, length}
// s := *(*[]byte)(unsafe.Pointer(&s1))  //不明白这是什么意思
// 构造一个虚拟的结构体，把 slice 的数据结构拼出来。

// 当然还有更加直接的方法，在 Go 的反射中就存在一个与之对应的数据结构 SliceHeader，我们可以用它来构造一个 slice

//     var o []byte
//     sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
//     sliceHeader.Cap = length
//     sliceHeader.Len = length
//     sliceHeader.Data = uintptr(ptr)

//reflect.SliceHeader
// type SliceHeader
// type SliceHeader struct {
//     Data uintptr
//     Len  int
//     Cap  int
// }
// SliceHeader代表一个运行时的切片。它不保证使用的可移植性、安全性；它的实现在未来的版本里也可能会改变。而且，Data字段也不能保证它指向的数据不会被当成垃圾收集，因此程序必须维护一个独立的、类型正确的指向底层数据的指针。

// unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
// unsafe.Pointer 可以和 普通指针 进行相互转换；
// unsafe.Pointer 可以和 uintptr 进行相互转换。
// 举例
// 通过一个例子加深理解，接下来尝试用指针的方式给结构体赋值。
// package main

// import (
//  "fmt"
//  "unsafe"
// )

// type W struct {
//  b int32
//  c int64
// }

// func main() {
//  var w *W = new(W)
//  //这时w的变量打印出来都是默认值0，0
//  fmt.Println(w.b,w.c)

//  //现在我们通过指针运算给b变量赋值为10
//  b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
//  *((*int)(b)) = 10
//  //此时结果就变成了10，0
//  fmt.Println(w.b,w.c)
// }
// uintptr(unsafe.Pointer(w)) 获取了 w 的指针起始值
// unsafe.Offsetof(w.b) 获取 b 变量的偏移量
// 两个相加就得到了 b 的地址值，将通用指针 Pointer 转换成具体指针 ((*int)(b))，通过 * 符号取值，然后赋值。*((*int)(b)) 相当于把 (*int)(b) 转换成 int 了，最后对变量重新赋值成 10，这样指针运算就完成了。

//nil 切片的指针指向 nil

// 1.1.4. 切片扩容
// 当一个切片的容量满了，就需要扩容了。怎么扩，策略是什么？

// 如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。

// 一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一

// 过打印的结果，我们可以看到，在这种情况下，扩容以后并没有新建一个新的数组，扩容前后的数组都是同一个，这也就导致了新的切片修改了一个值，也影响到了老的切片了。。并且 append() 操作也改变了原来数组里面的值。一个 append() 操作影响了这么多地方，如果原数组上有多个切片，那么这些切片都会被影响！无意间就产生了莫名的 bug！

// 这种情况，由于原数组还有容量可以扩容，所以执行 append() 操作以后，会在原数组上直接操作，所以这种情况下，扩容以后的数组还是指向原来的数组。
// 情况二其实就是在扩容策略里面举的例子，在那个例子中之所以生成了新的切片，是因为原来数组的容量已经达到了最大值，再想扩容， Go 默认会先开一片内存区域，把原来的值拷贝过来，然后再执行 append() 操作。这种情况丝毫不影响原数组。

// 所以建议尽量避免情况一，尽量使用情况二，避免 bug 产生。

//copy与append两种深拷贝方式，copy性能更好，建议使用copy。
// append会创建一个新的切片

// 判断某个键是否存在
// Go语言中有个判断map中键是否存在的特殊写法，格式如下:

//     value, ok := map[key]
// 举个例子：

// func main() {
//     scoreMap := make(map[string]int)
//     scoreMap["张三"] = 90
//     scoreMap["小明"] = 100
//     // 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
//     v, ok := scoreMap["张三"]
//     if ok {
//         fmt.Println(v)
//     } else {
//         fmt.Println("查无此人")
//     }
// }

// 最通俗的话说Map是一种通过key来获取value的一个数据结构，其底层存储方式为数组，在存储时key不能重复，当key重复时，value进行覆盖，我们通过key进行hash运算（可以简单理解为把key转化为一个整形数字）然后对数组的长度取余，得到key存储在数组的哪个下标位置，最后将key和value组装为一个结构体，放入数组下标处，看下图：

// length = len(array) = 4
// hashkey1 = hash(xiaoming) = 4
// index1  = hashkey1% length= 0
// hashkey2 = hash(xiaoli) = 6
// index2  = hashkey2% length= 2

//  hash冲突的常见解决方法
// 开放定址法：也就是说当我们存储一个key，value时，发现hashkey(key)的下标已经被别key占用，那我们在这个数组中空间中重新找一个没被占用的存储这个冲突的key，那么没被占用的有很多，找哪个好呢？常见的有线性探测法，线性补偿探测法，随机探测法，这里我们主要说一下线性探测法

// 线性探测，字面意思就是按照顺序来，从冲突的下标处开始往后探测，到达数组末尾时，从数组开始处探测，直到找到一个空位置存储这个key，当数组都找不到的情况下回扩容（事实上当数组容量快满的时候就会扩容了）；查找某一个key的时候，找到key对应的下标，比较key是否相等，如果相等直接取出来，否则按照顺寻探测直到碰到一个空位置，说明key不存在。如下图：首先存储key=xiaoming在下标0处，当存储key=xiaowang时，hash冲突了，按照线性探测，存储在下标1处，（红色的线是冲突或者下标已经被占用了） 再者key=xiaozhao存储在下标4处，当存储key=xiaoliu是，hash冲突了，按照线性探测，从头开始，存储在下标2处 （黄色的是冲突或者下标已经被占用了）

// 拉链法：何为拉链，简单理解为链表，当key的hash冲突时，我们在冲突位置的元素上形成一个链表，通过指针互连接，当查找时，发现key冲突，顺着链表一直往下找，直到链表的尾节点，找不到则返回空，如下图：

// 开放定址（线性探测）和拉链的优缺点

// 由上面可以看出拉链法比线性探测处理简单

// 线性探测查找是会被拉链法会更消耗时间

// 线性探测会更加容易导致扩容，而拉链不会

// 拉链存储了指针，所以空间上会比线性探测占用多一点

// 拉链是动态申请存储空间的，所以更适合链长不确定的
