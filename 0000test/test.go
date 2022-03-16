package main

import (
	"fmt"
)

// // 面试题1
// type student struct {
// 	name string
// 	age  int
// }

// // // 面试题2
// type student struct {
// 	id   int
// 	name string
// 	age  int
// }

// func demo(ce []student) {
// 	//切片是引用传递，是可以改变值的
// 	ce[1].age = 999
// 	// ce = append(ce, student{3, "xiaowang", 56})
// 	// return ce
// }
// //面试题3
// type People interface {
//     Speak(string) string
// }

// type Student struct{}

// func (stu *Stduent) Speak(think string) (talk string) {
//     if think == "sb" {
//         talk = "你是个大帅比"
//     } else {
//         talk = "您好"
//     }
//     return
// }
func main() {
	// var a *int
	// *a = 100  //应该是个地址// 或者new一下
	// fmt.Println(*a)

	// var b map[string]int  //需要定义make分配内存
	// b["测试"] = 100
	// fmt.Println(b)

	// 	// 面试题1
	// 	m := make(map[string]*student)
	// 	stus := []student{
	// 		{name: "pprof.cn", age: 18},
	// 		{name: "测试", age: 23},
	// 		{name: "博客", age: 28},
	// 	}

	// 	for _, stu := range stus {
	// 		m[stu.name] = &stu  //循环过程中，stu变量只声明了一次，所以stu地址即&stu是不变的，值是变化的。所以&stu始终不变
	// 	}
	// 	for k, v := range m {
	// 		fmt.Println(k, "=>", v.name)
	// 	}
	// 	// 	pprof.cn => 博客
	// 	// 测试 => 博客
	// 	// 博客 => 博客
	// // for range 每次产生的 key 和 value 其实是对应的 stus 里面值的拷贝，不是对应的 stus 里面的值的引用，所以出现了这种问题。
	// // stu 是 stus 在for循环中申请的一个局部变量，每次循环都会拷贝 stus 中对应的值 stu。迭代遍历之后，stu 每次会被重新赋值，而在 m 这个 map 中记录的 value 只不过是 stu 的内存地址。
	//可能是因为每次定义数据，用的是同一个地址，然后地址相同

	// // 重新申请一个变量，即可解决
	// //     for _, stu := range stus {
	// //         s:=stu
	// //         m[stu.name] = &s
	// //     }
	// // 面试题2
	// var ce []student //定义一个切片类型的结构体
	// ce = []student{
	// 	student{1, "xiaoming", 22},
	// 	student{2, "xiaozhang", 33},
	// }
	// fmt.Println(ce)
	// demo(ce)
	// fmt.Println(ce)

	// 	[{1 xiaoming 22} {2 xiaozhang 33}]
	// [{1 xiaoming 22} {2 xiaozhang 999}]
	// //面试题3
	// var peo People = Student{}
	// think := "bitch"
	// fmt.Println(peo.Speak(think))
}
