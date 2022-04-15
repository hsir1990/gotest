package main

import (
	_ "fmt"
	_ "math/rand"
	_ "runtime"
	_ "time"
)

func main() {
	//面试13
	// var myMap map[int]string
	// //myMap = make(map[int]string, 10)//只有make以后才不报错
	// myMap[0] = "golang"
	// //面试11
	// type A struct {
	// 	Num int
	// }
	// type B struct {
	// 	Num int
	// }
	// var a A
	// var b B
	// // fmt.Println(a == b) //不能做比较  报错invalid operation: a == b (mismatched types A and B)  // A和B不是一个类型

	// a = A(b) // ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
	// fmt.Println(a, b)

	// 	//面试10 select
	// 	// 创建2个管道
	// 	int_chan := make(chan int, 1)
	// 	string_chan := make(chan string, 1)
	// 	go func() {
	// 		//time.Sleep(2 * time.Second)
	// 		int_chan <- 1
	// 	}()
	// 	go func() {
	// 		string_chan <- "hello"
	// 	}()
	// 	select {
	// 	case value := <-int_chan:
	// 		fmt.Println("int:", value)
	// 	case value := <-string_chan:
	// 		fmt.Println("string:", value)
	// 	}
	// 	fmt.Println("main结束")

	// // // 	string: hello
	// // // main结束
	// // 或者
	// // int: 1
	// // main结束

	// //select 一次只能选择一个

	// 面试9 timer
	// // 1.timer基本使用
	// timer1 := time.NewTimer(2 * time.Second)
	// t1 := time.Now()
	// fmt.Printf("t1:%v\n", t1)
	// t2 := <-timer1.C
	// fmt.Printf("t2:%v\n", t2)

	// // 	t1:2022-03-17 11:32:20.2580423 +0800 CST m=+0.004010401
	// // t2:2022-03-17 11:32:22.2582508 +0800 CST m=+2.004218901

	// 	// 2.验证timer只能响应1次
	// 	timer2 := time.NewTimer(time.Second)
	// 	for {
	// 	<-timer2.C
	// 	fmt.Println("时间到")
	// 	}
	// // 	时间到
	// // fatal error: all goroutines are asleep - deadlock!

	// 3.timer实现延时的功能
	//(1)
	//time.Sleep(time.Second)
	//(2)
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2秒到")
	//(3)
	//<-time.After(2*time.Second)
	//fmt.Println("2秒到")

	// // 4.停止定时器
	// timer4 := time.NewTimer(2 * time.Second)
	// go func() {
	// 	<-timer4.C
	// 	fmt.Println("定时器执行了")
	// }()
	// b := timer4.Stop()
	// if b {
	// 	fmt.Println("timer4已经关闭")
	// }

	// //timer4已经关闭

	// 5.重置定时器
	// timer5 := time.NewTimer(3 * time.Second)
	// timer5.Reset(1 * time.Second)
	// fmt.Println(time.Now())
	// fmt.Println(<-timer5.C)
	//2022-03-17 11:28:32.6393863 +0800 CST m=+0.004034401
	// 2022-03-17 11:28:33.640384 +0800 CST m=+1.005032101

	// for {
	// }
	//  // 1.获取ticker对象
	//  ticker := time.NewTicker(1 * time.Second)
	//  i := 0
	//  // 子协程
	//  go func() {
	// 	 for {
	// 		 //<-ticker.C
	// 		 i++
	// 		 fmt.Println(<-ticker.C)
	// 		 if i == 5 {
	// 			 //停止
	// 			 ticker.Stop()
	// 		 }
	// 	 }
	//  }()
	//  for {
	//  }
	//  2022-03-17 11:38:53.0633575 +0800 CST m=+1.004703901
	// 2022-03-17 11:38:54.0630149 +0800 CST m=+2.004361301
	// 2022-03-17 11:38:55.0627367 +0800 CST m=+3.004083201
	// 2022-03-17 11:38:56.0633966 +0800 CST m=+4.004743101
	// 2022-03-17 11:38:57.0627553 +0800 CST m=+5.004101801

	// 	// 面试8 Goroutine池
	//  // 需要2个管道
	//     // 1.job管道
	//     jobChan := make(chan *Job, 128)
	//     // 2.结果管道
	//     resultChan := make(chan *Result, 128)
	//     // 3.创建工作池
	//     createPool(64, jobChan, resultChan)
	//     // 4.开个打印的协程
	//     go func(resultChan chan *Result) {
	//         // 遍历结果管道打印
	//         for result := range resultChan {
	//             fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
	//                 result.job.RandNum, result.sum)
	//         }
	//     }(resultChan)
	//     var id int
	//     // 循环创建job，输入到管道
	//     for {
	//         id++
	//         // 生成随机数
	//         r_num := rand.Int()
	//         job := &Job{
	//             Id:      id,
	//             RandNum: r_num,
	//         }
	//         jobChan <- job
	//     }

	// // 	job id:2827026 randnum:6305581339897553207 result:89
	// // job id:2827027 randnum:3016094518478363371 result:79
	// // job id:2827028 randnum:2302019554749561176 result:77
	// // job id:2827029 randnum:7184152774698280978 result:103
	// // job id:2827030 randnum:8750094336704324887 result:88
	// // job id:2827031 randnum:4970689211519738031 result:84
	// // job id:2827032 randnum:7541462261557285868 result:92
	// // job id:2827033 randnum:6529036923274841958 result:93
	// // job id:2827034 randnum:1505574606295606430 result:74
	// // job id:2827035 randnum:389025848969403138 result:90
	// // job id:2827036 randnum:4519465854605333529 result:87
	// // job id:2827037 randnum:9174186234974291120 result:80
	// // job id:2827038 randnum:2074139606246420748 result:75
	// // job id:2827039 randnum:8633768963163929835 result:105
	// // job id:2827040 randnum:4975341993075920005 result:82
	// // job id:2827041 randnum:1121316154217137469 result:65
	// // job id:2827042 randnum:2922257509989915683 result:101
	// // job id:2827043 randnum:7282684191410259282 result:81
	// // job id:2827044 randnum:8646541333201342530 result:63
	// // job id:2827045 randnum:863639991280242677 result:92
	// // job id:2827046 randnum:317095901630152187 result:68
	// // job id:2827047 randnum:5272209963828030412 result:73
	// //等等

	// // 面试7
	// c := make(chan int)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()
	// for {
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Println("main结束")

	// // 	0
	// // 1
	// // 2
	// // 3
	// // 4
	// // main结束

	// //面试6
	// go func() {
	//     defer fmt.Println("A.defer")
	//     func() {
	//         defer fmt.Println("B.defer")
	//         // 结束协程
	//         runtime.Goexit()
	//         defer fmt.Println("C.defer")
	//         fmt.Println("B")
	//     }()
	//     fmt.Println("A")
	// }()
	// for {
	// }
	//面试5
	// go func(s string) {
	//     for i := 0; i < 2; i++ {
	//         fmt.Println(s)
	//     }
	// }("world")
	// // 主协程
	// for i := 0; i < 2; i++ {
	//     // 切一下，再次分配任务
	//     runtime.Gosched()
	//     fmt.Println("hello")
	// }

	// // world
	// // world
	// // hello
	// // hello

	//面试1
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

	// 	// //面试题4
	// 	 // 合起来写
	// 	 go func() {
	//         i := 0
	//         for {
	//             i++
	//             fmt.Printf("new goroutine: i = %d\n", i)
	//             time.Sleep(time.Second)
	//         }
	//     }()
	//     i := 0
	//     for {
	//         i++
	//         fmt.Printf("main goroutine: i = %d\n", i)
	//         time.Sleep(time.Second)
	//         if i == 2 {
	//             break
	//         }
	//     }

	// // 	$ go run test.go
	// // main goroutine: i = 1
	// // new goroutine: i = 1
	// // main goroutine: i = 2
	// // new goroutine: i = 2
	// // new goroutine: i = 3

}

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
// // 面试8Goroutine池
// type Job struct {
//     // id
//     Id int
//     // 需要计算的随机数
//     RandNum int
// }

// type Result struct {
//     // 这里必须传对象实例
//     job *Job
//     // 求和
//     sum int
// }

// // 创建工作池
// // 参数1：开几个协程
// func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
//     // 根据开协程个数，去跑运行
//     for i := 0; i < num; i++ {
//         go func(jobChan chan *Job, resultChan chan *Result) {
//             // 执行运算
//             // 遍历job管道所有数据，进行相加
//             for job := range jobChan {
//                 // 随机数接过来
//                 r_num := job.RandNum
//                 // 随机数每一位相加
//                 // 定义返回值
//                 var sum int
//                 for r_num != 0 {
//                     tmp := r_num % 10
//                     sum += tmp
//                     r_num /= 10
//                 }
//                 // 想要的结果是Result
//                 r := &Result{
//                     job: job,
//                     sum: sum,
//                 }
//                 //运算结果扔到管道
//                 resultChan <- r
//             }
//         }(jobChan, resultChan)
//     }
// }
