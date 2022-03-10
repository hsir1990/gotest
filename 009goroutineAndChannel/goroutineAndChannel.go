package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁，
	//sync 是包: synchornized 同步
	//Mutex : 是互斥
	lock sync.Mutex
)

//使用协程+管道  ===>???
// test1 函数就是计算 n!, 让将这个结果放入到 myMap
func test1(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//这里我们将 res 放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res //concurrent map writes?
	//解锁
	lock.Unlock()
}

type Cat struct {
	Name string
	Age  int
}

func main() {
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
	fmt.Println("num=", num)
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

	//不同goroutine之间如何通讯
	// 1）全局变量的互斥锁
	// 2）使用管道channel来解决

	//使用全局变量加锁同步改进程序
	//因为没有对全局变量 m 加锁，因此会出现资源争夺问题，代码会出现错误，提示concurrent map writes
	//解决方案：加入互斥锁
	//我们的数的阶乘很大，结果会越界，可以将求阶乘改成  sum+=uint64(i)
	//代码改进

	// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
	// 最后显示出来。要求使用goroutine完成

	// 思路
	// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
	// 2. 我们启动的协程多个，统计的将结果放入到 map中
	// 3. map 应该做出一个全局的.

	// 我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 20; i++ {
		go test1(i)
	}

	//休眠10秒钟【第二个问题 】
	//time.Sleep(time.Second * 5)

	//这里我们输出结果,变量这个结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
	// num= 12
	// map[8]=40320
	// map[16]=20922789888000
	// map[18]=6402373705728000
	// map[7]=5040
	// map[1]=1
	// map[17]=355687428096000
	// map[14]=87178291200
	// map[4]=24
	// map[15]=1307674368000
	// map[3]=6
	// map[11]=39916800
	// map[6]=720
	// map[10]=3628800
	// map[12]=479001600
	// map[13]=6227020800
	// map[19]=121645100408832000
	// map[2]=2
	// map[5]=120

	//为什么需要channel
	// 1)前面使用全局变量加锁同步来解决goroutine的通讯，但不完美
	// 2）主线程在等待所有goroutine全部完成的时间很难确定，我们这里设置10秒，仅仅是估算
	// 3）如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有goroutine处于工作状态，这时也会随主线程的退出而销毁
	// 4）通过全局变量加锁同步来实现通讯，也并不利用多个协程对全局变量的读写操作。
	// 5）上面种种分析都在呼唤一个新的通讯机制 -- chanel

	//channel 的基本介绍
	// 1）chanel本质就是一个数据结构--队列
	// 2）数据是先进先出【FIFO：first in first out】
	// 3)线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的 ，多协程操作同一个管道时，不会发生资源竞争问题
	// 4）channel有类型的，一个string 的 channel 只能存放string

	//定义/声明 channel

	//var 变量名 chan 数据类型
	//举例
	// var intChan chan int (intChan 用于存在int数据)
	// var mapChan chan map[int]string (mapChan 用于存放map[int]string 类型)
	// var perChan chan Person
	// var perChan2 chan *Person

	//说明
	// channel 是引用类型
	// channel 必须初始化才能写入数据，即make后才能使用
	// 管道是有类型，intChan 只能写入 整数 int

	// 管道的初始化，写入数据到管道，从管道读取数据及基本的注意事项

	//演示一下管道的使用
	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2. 看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3. 向管道写入数据
	intChan <- 10
	num6 := 211
	intChan <- num6
	intChan <- 50
	// //如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 98 //注意点, 当我们给管写入数据时，不能超过其容量

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	//5. 从管道中读取数据

	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock

	num3 := <-intChan
	num4 := <-intChan

	//num5 := <-intChan

	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

	//channel使用的注意事项
	// 1) channel中只能存放指定的数据类型
	// 2）channel的数据放满后，就不能再放入了
	// 3）如果从channel取出数据后，可以继续放入
	// 4）在没有使用协程的情况下，如果channel数据取完了，再取，就会报dead lock

	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan <- cat

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么?

	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat) //newCat=main.Cat , newCat={小花猫 4}

	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name=%v", newCat.Name)

	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name) //newCat.Name=小花猫main()

	//channel 的遍历和关闭
	//使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写入数据了，但是仍然是可以从该channel读取数据
	//
	intChan1 := make(chan int, 3)
	intChan1 <- 100
	intChan1 <- 200
	close(intChan1) // close
	//这是不能够再写入数到channel
	//intChan1<- 300
	fmt.Println("okook~")
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan1
	fmt.Println("n1=", n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}

	//遍历管道不能使用普通的 for 循环
	// for i := 0; i < len(intChan2); i++ {

	// }
	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)

	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	//请完成goroutine和channel 协同工作的案例
	// 1开启一个writeData 协程，向管道intChan中写入50个整数
	//2 开启一个readData协程，从管道intChan中读取writeData写入的数据
	//3注意 writeData和readData操作的是同一个管道
	//4 主线程需要等待writeData和readData协程都完成工作才能退出【管道】
	//创建两个管道
	intChan3 := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan3)
	go readData(intChan3, exitChan)

	//time.Sleep(time.Second * 10)
	//就是为了等待readData协程完成
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	// writeData  1
	// writeData  2
	// writeData  3
	// writeData  4
	// writeData  5
	// writeData  6
	// writeData  7
	// writeData  8
	// writeData  9
	// writeData  10
	// writeData  11
	// readData 读到数据=1
	// readData 读到数据=2
	// readData 读到数据=3
	// readData 读到数据=4
	// readData 读到数据=5
	// readData 读到数据=6
	// readData 读到数据=7
	// readData 读到数据=8
	// readData 读到数据=9
	// readData 读到数据=10
	// readData 读到数据=11
	// readData 读到数据=12
	// writeData  12
	// writeData  13
	// writeData  14
	// writeData  15
	// writeData  16
	// writeData  17
	// writeData  18
	// writeData  19
	// writeData  20
	// writeData  21
	// writeData  22
	// writeData  23
	// readData 读到数据=13
	// readData 读到数据=14
	// readData 读到数据=15
	// readData 读到数据=16
	// readData 读到数据=17
	// readData 读到数据=18
	// readData 读到数据=19
	// readData 读到数据=20
	// readData 读到数据=21
	// readData 读到数据=22
	// readData 读到数据=23
	// readData 读到数据=24
	// writeData  24
	// writeData  25
	// writeData  26
	// writeData  27
	// writeData  28
	// writeData  29
	// writeData  30
	// writeData  31
	// writeData  32
	// writeData  33
	// writeData  34
	// writeData  35
	// readData 读到数据=25
	// readData 读到数据=26
	// readData 读到数据=27
	// readData 读到数据=28
	// readData 读到数据=29
	// readData 读到数据=30
	// readData 读到数据=31
	// readData 读到数据=32
	// readData 读到数据=33
	// readData 读到数据=34
	// readData 读到数据=35
	// readData 读到数据=36
	// writeData  36
	// writeData  37
	// writeData  38
	// writeData  39
	// writeData  40
	// writeData  41
	// writeData  42
	// writeData  43
	// writeData  44
	// writeData  45
	// writeData  46
	// writeData  47
	// readData 读到数据=37
	// readData 读到数据=38
	// readData 读到数据=39
	// readData 读到数据=40
	// readData 读到数据=41
	// readData 读到数据=42
	// readData 读到数据=43
	// readData 读到数据=44
	// readData 读到数据=45
	// readData 读到数据=46
	// readData 读到数据=47
	// readData 读到数据=48
	// writeData  48
	// writeData  49
	// writeData  50
	// readData 读到数据=49
	// readData 读到数据=50

	//阻塞
	//如果编辑器（运行），发现一个管道只有写，而没有读，则该管道，会堵塞
	//写管道和读管道的频率不一致，无所谓
	//问题：如果注销掉go readData(intChan, exitChan),程序会怎么样
	//答：如果只是向管道写入数据，而没有读取，就会出现堵塞而 dead lock，原因是intChan容量是10，
	//而代码writeData会写入50个数据，因此会阻塞在writeData的 ch <- i

	//需求
	//要求统计1-200000的数字中，哪些是素数？这个问题在本章开篇就提出了，现在我们有goroutine和channel的知识后，就可以完成了【测试数据： 80000】
	//分析思路
	//传统的方法，就是使用一个循环，循环的判断各个数是不是素数【ok】
	//使用并发/并行的方式，将统计素数的任务分配给多个（4个）goroutine去完成，完成任务时间短。

	intChan4 := make(chan int, 1000)
	primeChan := make(chan int, 20000) //放入结果
	//标识退出的管道
	exitChan1 := make(chan bool, 8) // 4个

	start := time.Now().Unix()

	//开启一个协程，向 intChan4放入 1-8000个数
	go putNum(intChan4)
	//开启4个协程，从 intChan4取出数据，并判断是否为素数,如果是，就
	//放入到primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan4, primeChan, exitChan1)
	}

	//这里我们主线程，进行处理
	//直接
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan1
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)

		//当我们从exitChan1 取出了4个结果，就可以放心的关闭 prprimeChan
		close(primeChan)
	}()

	//遍历我们的 primeChan ,把结果取出
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		//fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")

	//使用go协程后，执行的速度，比普通方法提高至少4倍
	// 素数 79901
	// 素数 79907
	// 素数 79903
	// 素数 79943
	// 有一个primeNum 协程因为取不到数据，退出
	// 素数 79939
	// 有一个primeNum 协程因为取不到数据，退出
	// 素数 79967
	// 有一个primeNum 协程因为取不到数据，退出
	// 素数 79979
	// 素数 79973
	// 有一个primeNum 协程因为取不到数据，退出
	// 有一个primeNum 协程因为取不到数据，退出
	// 素数 79987
	// 有一个primeNum 协程因为取不到数据，退出
	// 素数 79997
	// 有一个primeNum 协程因为取不到数据，退出
	// 素数 79999
	// 有一个primeNum 协程因为取不到数据，退出
	// 使用协程耗时= 0
	// main线程退出

	//channel使用细节和注意事项
	//1）channel可以声明为只读，或者只写性质

	//管道可以声明为只读或者只写

	//1. 在默认情况下下，管道是双向
	//var chan1 chan int //可读可写

	// //2 声明为只写
	// var chan2 chan<- int
	// chan2 = make(chan int, 3)
	// chan2 <- 20
	// //num := <-chan2 //error

	// fmt.Println("chan2=", chan2)

	// //3. 声明为只读
	// var chan3 <-chan int
	// num7 := <-chan3
	// //chan3<- 30 //err
	// fmt.Println("num7", num7)

	//channel 只读和只写的最佳实践案例

	//使用select可以解决从管道取数据的阻塞问题

	//使用select可以解决从管道取数据的阻塞问题

	//1.定义一个管道 10个数据int
	intChan6 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan6 <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock

	//问题，在实际开发中，可能我们不好确定什么关闭该管道.
	//可以使用select 方式可以解决
	//label:
	for {
		select {
		//注意: 这里，如果intChan6一直没有关闭，不会一直阻塞而deadlock
		//，会自动到下一个case匹配
		case v := <-intChan6:
			fmt.Printf("从intChan6读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			return
			//break label
		}
	}

	// 	从stringChan读取的数据hello0
	// 从intChan6读取的数据0
	// 从intChan6读取的数据1
	// 从stringChan读取的数据hello1
	// 从intChan6读取的数据2
	// 从stringChan读取的数据hello2
	// 从stringChan读取的数据hello3
	// 从stringChan读取的数据hello4
	// 从intChan6读取的数据3
	// 从intChan6读取的数据4
	// 从intChan6读取的数据5
	// 从intChan6读取的数据6
	// 从intChan6读取的数据7
	// 从intChan6读取的数据8
	// 从intChan6读取的数据9
	// 都取不到了，不玩了, 程序员可以加入逻辑

	// goroute中使用recover,解决协程中出现panic，导致程序崩溃问题
	//说明：如果我们起一个协程，但是这个协程出现了panic，如果我们没有捕获这个panic，就会造成整个
	//程序崩溃，这时我们可以在goroutine中使用recover来捕获panic，进行处理，这样即使这个协程发生的
	//问题，但是主线程仍然不受影响，可以继续执行。

	go sayHello()
	go test2()


	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}


	//time.Sleep(time.Second * 10) //休眠10秒，防止主线程自己关闭
	go test() //开启一个线程
	for i := 0; i < 10; i++ {
		fmt.Println("main() hello,world " + strconv.Itoa(i))
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


//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
//函数
func test2() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}




//向 intChan放入 1-8000个数
func putNum(intChan chan int) {

	for i := 1; i <= 80000; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)
}

// 从 intChan取出数据，并判断是否为素数,如果是，就
// 	//放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for 循环
	// var num int
	var flag bool //
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan //intChan 取不到..

		if !ok {
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			fmt.Println("素数", num)
			//将这个数就放入到primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入true
	exitChan <- true

}

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i //
		fmt.Println("writeData ", i)
		//time.Sleep(time.Second)
	}
	close(intChan) //关闭
}

//read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		// time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)

}

//编写一个函数，每隔一秒输出“hello，world”
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
