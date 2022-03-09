package main
import(
	"fmt"
	"strconv"
	"time"
	"runtime"
)
func main(){
	//1.进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位
//2.线程是进程的一个执行实例，是程序执行的最小单元，它是进程更小的能独立运行的基本单位
//3.一个进程可以创建核销多个线程，同一个进程中的多个线程可以并发执行
//4.一个程序至少有一个进程，一个进程至少有一个线程



//多线程程序在单核上运行，就是并发
//多线程在多核上运行，就是并行

//go协程和go主线程
// go主线程（有程序员直接称为线程/也可以理解成进程）；一个go线程上，可以起多个协程，
// 你可以这样理解，协程是轻量级的线程【编译器做优化】


// go协程的特点
//1 有独立的栈空间
//2 共享程序堆空间
//3 调度由用户控制
//4 协程是轻量级的线程





//小结
//1. 主线程是一个物理线程，直接作用在cpu上，是重量级的，非常耗费cpu资源
//2. 协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小。
//3. golang的协程机制是重要的特点，可以轻松的开启上万个协程。其他编程语言的并发机制是一般
//基于线程的，开启过多的线程，资源消耗费大，这里就突显Golang在并发上的优势了


//goroutine的调度模型
//MPG模式基本介绍
// M 操作系统的主线程（是物理线程）
// P 协程执行需要的上下文
// G 协程


//一个cpu有一个进程，多个线程，多个线程都可以当作主线程，主线程上有可以掉起多个协程

//设置golang运行的cpu数
//获取当前系统cpu的数量
num := runtime.NumCPU()
//我们这里设置num-1在cpu运行go程序
runtime.GOMAXPROCS(num)
fmt.Println("num=",num)
//go1.8后，默认让程序运行在多个核上，可以不用设置了
//go1.8前，还是要设置一下，可以更高效的利益cpu

//channel管道
//计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来，用goroutine完成

//思路
//1 使用gorutine来完成，效率高，但是会出现并发，并行安全问题
//2 这里就提出了不同goroutine 如何通讯的问题


//代码实现
//1 使用goroutine 来完成（看看使用goroutine并发完成会造成什么问题，然后我们会去解决）
//2 在运行某个程序时，如何知道是否存在资源竞争问题。
// 方法很简单，在编译该程序时，增加一个参数 -race 即可




















go test() //开启一个线程
for i:= 0;i<10;i++{
	fmt.Println("main() hello,world "+ strconv.Itoa(i))
	time.Sleep(time.Second)
}
// 输出的效果说明，main这个主线程和test协程同时执行
// main() hello,world 0
// test() hello,world 0
// test() hello,world 1
// main() hello,world 1
// test() hello,world 2
// main() hello,world 2
// main() hello,world 3
// test() hello,world 3
// test() hello,world 4
// main() hello,world 4
// main() hello,world 5
// test() hello,world 5
// test() hello,world 6
// main() hello,world 6
// test() hello,world 7
// main() hello,world 7
// test() hello,world 8
// main() hello,world 8
// test() hello,world 9
// main() hello,world 9
}
//编写一个函数，每隔一秒输出“hello，world”
func test(){
	for i:= 0;i<10;i++{
		fmt.Println("test() hello,world "+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}