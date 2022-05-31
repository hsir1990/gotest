func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()  //按栈的形式，先进后出,后进的先出
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	fmt.Println("打印1")
	panic("触发异常")  //panic会中断执行，反应到最上层，所以最后处理，
	fmt.Println("打印2")  //被panic中断

}

// 打印1
// 打印后
// 打印中
// 打印前
// panic: 触发异常



type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu  //stu是一个变量，所有打印出来的都是一样的值，也就是最后一个赋值的值 //&str 具有迷惑性，他是变量的地址
 		fmt.Printf("%v,%p",m[stu.Name],m[stu.Name]) 
		 //&{zhou 24},0xc000004078&{li 23},0xc000004078&{wang 22},0xc000004078
	}

	fmt.Println(m)//map[li:0xc000004078 wang:0xc000004078 zhou:0xc000004078]

	for key,value := range m{
		fmt.Println(key," ",value)
		// zhou   &{wang 22}
		// li   &{wang 22}
		// wang   &{wang 22}

	}
}


// func GOMAXPROCS(n int) int
// GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置。 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。本函数在调度程序优化后会去掉。
// func NumCPU() int
// NumCPU返回本地机器的逻辑CPU个数。
runtime.GOMAXPROCS(1)  //设置最大的cpu数
// type WaitGroup
// type WaitGroup struct {
//     // 包含隐藏或非导出字段
// }
// WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
wg := sync.WaitGroup{} //同步锁
// func (*WaitGroup) Add
// func (wg *WaitGroup) Add(delta int)
// Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。

// func (*WaitGroup) Done
// func (wg *WaitGroup) Done()
// Done方法减少WaitGroup计数器的值，应在线程的最后执行。

// func (*WaitGroup) Wait
// func (wg *WaitGroup) Wait()
// Wait方法阻塞直到WaitGroup计数器减为0。
wg.Add(20)
for i := 0; i < 10; i++ {
	//i想对匿名函数来说，是函数外面的变量
	go func() {  //进入了其他线程，所以变相成了异步  //这个可以看成一组线程，for出了这么多线程
		fmt.Println("i: ", i)
		wg.Done()
	}()

	func() {
		fmt.Println("i: ", i)
	}()
}
for i := 0; i < 10; i++ {
	go func(i int) {
		fmt.Println("i: ", i)
		wg.Done()
	}(i)
}
wg.Wait()
i:  0
i:  1
i:  2
i:  3
i:  4
i:  5
i:  6
i:  7
i:  8
i:  9
i:  9
i:  10
i:  10
i:  10
i:  10
i:  10
i:  10
i:  10
i:  10
i:  10
i:  10
i:  0
i:  1
i:  2
i:  3
i:  4
i:  5
i:  6
i:  7
i:  8


type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB() //这个很明显是进入了p，才调用的showB，所以调用自己showB方法
}
func (p *People) ShowB() {
	fmt.Println("showB") 
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
// showA
// showB


func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:  //在这里面重复使用value不报错
		fmt.Println(value)
	case value := <-string_chan://在这里面重复使用value不报错
		panic(value)
	}
}
//select的调度是随机的，内部有算法，可能是打印1，也可能先报错
//没有用for一般只执行一次


var a  int = 3
  // 以下有额外内存分配吗？
  var i interface{} = a
答案：小整数转换为接口值不再需要进行内存分配。小整数是指 0 到 255 之间的数。

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	//defer的函数参数，在定义的时候就传入了栈中并执行了参数中的函数，所以现在 a=1 b =2 //有点快照的意思
	//按栈的形式，先进后出
	defer calc("1", a, calc("10", a, b))   //里面的函数在进站的时候，就已经执行了
	a = 0
	defer calc("2", a, calc("20", a, b))  //a =0 b=2
	b = 1
}

10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4


func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)  //等于是插入了3个元素
	fmt.Println(s)
}

//[0 0 0 0 0 1 2 3]


type UserAges struct {
	ages map[string]int
	// type Mutex
	// type Mutex struct {
	// 	// 包含隐藏或非导出字段
	// }
	// Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {

	// func (*Mutex) Lock
	// func (m *Mutex) Lock()
	// Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。

	// func (*Mutex) Unlock
	// func (m *Mutex) Unlock()
	// Unlock方法解锁m，如果m未加锁会导致运行时错误。锁和线程无关，可以由不同的线程加锁和解锁。

	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//add加了锁，第二个Get函数没有加锁，所以这个肯定有并发的问题，等于写加了锁，读没有加



package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {  //通过指针的形式引用，多态使用也要用指针
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//因为student在Speak函数中是一个指针引用，所以它赋值也要用地址的方式实现多态   var peo People = &Stduent{}
	var peo People = Stduent{}  
	think := "bitch"
	fmt.Println(peo.Speak(think))
}



package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
//打印 BBBBBB


//指针对象被赋值以后，它打印的值，可能是<nil>，但不是初始化时候的nil了
func main(){
	var peo People
	fmt.Printf("%T, %v\n",peo,peo)  //<nil>, <nil>
	if peo == nil {
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
	//yes
	var stu1 *Student  //这个地方使用指针，才能通过
	var peo1 People = stu1
	fmt.Printf("%T, %v\n",peo1,peo1) //*main.Student, <nil>
	if peo1 == nil {
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
	//no
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	//BBBBBBB
}

<nil>, <nil>
yes
*main.Student, <nil>
no
BBBBBBB



用path.join的目的是 ，Linux的目录是反斜杠，而win是正斜杠，用这个方法就统一了
打印行号是非常浪费资源的