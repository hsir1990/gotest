package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)   //可以定义，也可以定义并且赋值
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
	//select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
	//select 是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。 select 随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
	// 每个case都必须是一个通信
	// 所有channel表达式都会被求值
	// 所有被发送的表达式都会被求值
	// 如果任意某个通信可以进行，它就执行；其他被忽略。
	// 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	// 否则：
	// 如果有default子句，则执行该语句。
	// 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

	// 与switch语句可以选择任何使用相等比较的条件相比，select由比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作

	// select { //不停的在这里检测
	// 	case <-chanl : //检测有没有数据可以读
	// 	//如果chanl成功读取到数据，则进行该case处理语句
	// 	case chan2 <- 1 : //检测有没有可以写
	// 	//如果成功向chan2写入数据，则进行该case处理语句

	// 	//假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
	// 	default:
	// 	//如果以上都没有符合条件，那么则进行default处理流程
	// 	}

	// 	在一个select语句中，Go会按顺序从头到尾评估每一个发送和接收的语句。

	// 如果其中的任意一个语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。 如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况： ①如果给出了default语句，那么就会执行default的流程，同时程序的执行会从select语句后的语句中恢复。 ②如果没有default语句，那么select语句将被阻塞，直到至少有一个case可以进行下去。

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

// 超时判断
// //比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
// var resChan = make(chan int)
// // do request
// func test() {
//     select {
//     case data := <-resChan:
//         doData(data)
//     case <-time.After(time.Second * 3):
//         fmt.Println("request time out")
//     }
// }

// func doData(data int) {
//     //...
// }

// func (Time) After
// func (t Time) After(u Time) bool
// 如果t代表的时间点在u之后，返回真；否则返回假。

// func After
// func After(d Duration) <-chan Time
// After会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C。

// Example
// select {
// case m := <-c:
//     handle(m)
// case <-time.After(5 * time.Minute):
//     fmt.Println("timed out")
// }

// 2.退出
// //主线程（协程）中如下：
// var shouldQuit=make(chan struct{})
// fun main(){
//     {
//         //loop
//     }
//     //...out of the loop
//     select {
//         case <-c.shouldQuit:
//             cleanUp()
//             return
//         default:
//         }
//     //...
// }

// //再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
// close(shouldQuit)
// 3.判断channel是否阻塞
// //在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
// ch := make (chan int, 5)
// //...
// data：=0
// select {
// case ch <- data:
// default:
//     //做相应操作，比如丢弃data。视需求而定
// }

//两种引用类型 map、channel 是指针包装，而不像 slice 是 struct。

// 1.1.1. Goto、Break、Continue
//      1.三个语句都可以配合标签(label)使用
//     2.标签名区分大小写，定以后若不使用会造成编译错误
//     3.continue、break配合标签(label)可用于多层循环跳出
//     4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同

// //panic：
// 1、内置函数
//     2、假如函数F中书写了panic语句，  会终止其后要执行的代码  ，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
//     3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
//     4、直到goroutine整个退出，并报告错误

// 	//recover：
// 	1、内置函数
//     2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
//     3、一般的调用建议
//         a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
//         b). 可以获取通过panic传递的error

// 注意:

//     1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
//     2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
//     3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

// package main

// func main() {
//     test()
// }

// func test() {
//     defer func() {
//         if err := recover(); err != nil {
//             println(err.(string)) // 将 interface{} 转型为具体类型。
//         }
//     }()

//     panic("panic error!")
// }
// 输出结果：

//     panic error!
// 由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。

//     func panic(v interface{})
//     func recover() interface{}

// 	向已关闭的通道发送数据会引发panic

// package main

// import (
//     "fmt"
// )

// func main() {
//     defer func() {
//         if err := recover(); err != nil {
//             fmt.Println(err)
//         }
//     }()

//     var ch chan int = make(chan int, 10)
//     close(ch)
//     ch <- 1
// }
// 输出结果：

//     send on closed channel
// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。

// package main

// import "fmt"

// func test() {
//     defer func() {
//         fmt.Println(recover())
//     }()

//     defer func() {
//         panic("defer panic")
//     }()

//     panic("test panic")
// }

// func main() {
//     test()
// }
// 输出:

//     defer panic
// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。

// package main

// import "fmt"

// func test() {
//     defer func() {
//         fmt.Println(recover()) //有效
//     }()
//     defer recover()              //无效！
//     defer fmt.Println(recover()) //无效！
//     defer func() {
//         func() {
//             println("defer inner")
//             recover() //无效！
//         }()
//     }()

//     panic("test panic")
// }

// func main() {
//     test()
// }
// 输出:

//     defer inner
//     <nil>
//     test panic

// 	同一函数内 多个panic捕捉最后一个,多个recover ，只有一个有效

// //使用延迟匿名函数或下面这样都是有效的。

// package main

// import (
//     "fmt"
// )

// func except() {
//     fmt.Println(recover())
// }

// func test() {
//     defer except()
//     panic("test panic")
// }

// func main() {
//     test()
// }
// 输出结果：

//     test panic
// 	使用延迟匿名函数或下面这样都是有效的。

// 	package main

// 	import (
// 		"fmt"
// 	)

// 	func except() {
// 		fmt.Println(recover())
// 	}

// 	func test() {
// 		defer except()
// 		panic("test panic")
// 	}

// 	func main() {
// 		test()
// 	}
// 	输出结果：

// 		test panic
// 	如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执 。

// 	package main

// 	import "fmt"

// 	func test(x, y int) {
// 		var z int

// 		func() {
// 			defer func() {
// 				if recover() != nil {
// 					z = 0
// 				}
// 			}()
// 			panic("test panic")
// 			z = x / y
// 			return
// 		}()

// 		fmt.Printf("x / y = %d\n", z)
// 	}

// 	func main() {
// 		test(2, 1)
// 	}
// 	输出结果：

// 		x / y = 0  //这样不用中断
// 	除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。

// 	type error interface {
// 		Error() string
// 	}
// 	标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型。

// 	package main

// 	import (
// 		"errors"
// 		"fmt"
// 	)

// 	var ErrDivByZero = errors.New("division by zero")

// 	func div(x, y int) (int, error) {
// 		if y == 0 {
// 			return 0, ErrDivByZero
// 		}
// 		return x / y, nil
// 	}

// 	func main() {
// 		defer func() {
// 			fmt.Println(recover())
// 		}()
// 		switch z, err := div(10, 0); err {
// 		case nil:
// 			println(z)
// 		case ErrDivByZero:
// 			panic(err)
// 		}
// 	}
// 	输出结果：

// 		division by zero
// 	Go实现类似 try catch 的异常处理

// 	package main

// 	import "fmt"

// 	func Try(fun func(), handler func(interface{})) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				handler(err)
// 			}
// 		}()
// 		fun()
// 	}

// 	func main() {
// 		Try(func() {
// 			panic("test panic")
// 		}, func(err interface{}) {
// 			fmt.Println(err)
// 		})
// 	}
// 	输出结果：

// 		test panic
// 	如何区别使用 panic 和 error 两种方式?

// 	惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。

// 并发介绍
// 进程和线程
//     A. 进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
//     B. 线程是进程的一个执行实体,是CPU调度和分派的基本单位,它是比进程更小的能独立运行的基本单位。
//     C.一个进程可以创建和撤销多个线程;同一个进程中的多个线程之间可以并发执行。
// 并发和并行
//     A. 多线程程序在一个核的cpu上运行，就是并发。
//     B. 多线程程序在多个核的cpu上运行，就是并行。

// 协程和线程
// 协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
// 线程：一个线程上可以跑多个协程，协程是轻量级的线程。
// goroutine 只是由官方实现的超级"线程池"。
// 每个实力4~5KB的栈内存占用和由于实现机制而大幅减少的创建和销毁开销是go高并发的根本原因。

// 并发不是并行：
// 并发主要由切换时间片来实现"同时"运行，并行则是直接利用多核实现多线程的运行，go可以设置使用核数，以发挥多核计算机的能力。

// goroutine 奉行通过通信来共享内存，而不是共享内存来通信。

// 1. Goroutine
// 在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？

// Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

// 在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。

// 使用goroutine
// Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。

// 一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

// 启动单个goroutine
// 启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个go关键字。

// 举个例子如下：

// func hello() {
//     fmt.Println("Hello Goroutine!")
// }
// func main() {
//     hello()
//     fmt.Println("main goroutine done!")
// }
// 这个示例中hello函数和下面的语句是串行的，执行的结果是打印完Hello Goroutine!后打印main goroutine done!。

// 接下来我们在调用hello函数前面加上关键字go，也就是启动一个goroutine去执行hello这个函数。

// func main() {
//     go hello() // 启动另外一个goroutine去执行hello函数
//     fmt.Println("main goroutine done!")
// }
// 这一次的执行结果只打印了main goroutine done!，并没有打印Hello Goroutine!。为什么呢？

// 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。

// 当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，main函数所在的goroutine就像是权利的游戏中的夜王，其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。

// 所以我们要想办法让main函数等一等hello函数，最简单粗暴的方式就是time.Sleep了。

// func main() {
//     go hello() // 启动另外一个goroutine去执行hello函数
//     fmt.Println("main goroutine done!")
//     time.Sleep(time.Second)
// }
// 执行上面的代码你会发现，这一次先打印main goroutine done!，然后紧接着打印Hello Goroutine!。

// 首先为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。

// 启动多个goroutine
// 在Go语言中实现并发就是这样简单，我们还可以启动多个goroutine。让我们再来一个例子： （这里使用了sync.WaitGroup来实现goroutine的同步）

// var wg sync.WaitGroup

// func hello(i int) {
//     defer wg.Done() // goroutine结束就登记-1
//     fmt.Println("Hello Goroutine!", i)
// }
// func main() {

//     for i := 0; i < 10; i++ {
//         wg.Add(1) // 启动一个goroutine就登记+1
//         go hello(i)
//     }
//     wg.Wait() // 等待所有登记的goroutine都结束
// }
// 多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。

// 注意
// 如果主协程退出了，其他任务还执行吗（运行下面的代码测试一下吧）
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // 合起来写
//     go func() {
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
// }
// 1.1.1. goroutine与线程
// 可增长的栈
// OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这个大。所以在Go语言中一次创建十万左右的goroutine也是可以的。

// goroutine调度
// GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

// 1.G很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
// 2.P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
// 3.M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；
// P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。

// P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

// 单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。

// 1. runtime包
// 1.1.1. runtime.Gosched()
// 让出CPU时间片，重新等待安排任务(大概意思就是本来计划的好好的周末出去烧烤，但是你妈让你去相亲,两种情况第一就是你相亲速度非常快，见面就黄不耽误你继续烧烤，第二种情况就是你相亲速度特别慢，见面就是你侬我侬的，耽误了烧烤，但是还馋就是耽误了烧烤你还得去烧烤)

// package main

// import (
//     "fmt"
//     "runtime"
// )

// func main() {
//     go func(s string) {
//         for i := 0; i < 2; i++ {
//             fmt.Println(s)
//         }
//     }("world")
//     // 主协程
//     for i := 0; i < 2; i++ {
//         // 切一下，再次分配任务
//         runtime.Gosched()
//         fmt.Println("hello")
//     }
// }

// world
// world
// hello
// hello
// 	func Gosched
// func Gosched()
// Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。

// 1.1.2. runtime.Goexit()
// 退出当前协程(一边烧烤一边相亲，突然发现相亲对象太丑影响烧烤，果断让她滚蛋，然后也就没有然后了)

// package main

// import (
//     "fmt"
//     "runtime"
// )

// func main() {
//     go func() {
//         defer fmt.Println("A.defer")
//         func() {
//             defer fmt.Println("B.defer")
//             // 结束协程
//             runtime.Goexit()
//             defer fmt.Println("C.defer")
//             fmt.Println("B")
//         }()
//         fmt.Println("A")
//     }()
//     for {
//     }
// }

//B.defer
// A.defer

// func Goexit
// func Goexit()
// Goexit终止调用它的go程。其它go程不会受影响。Goexit会在终止该go程前执行所有defer的函数（不包含fmt.Println()）。

// 在程序的main go程调用本函数，会终结该go程，而不会让main返回。因为main函数没有返回，程序会继续执行其它的go程。如果所有其它go程都退出了，程序就会崩溃。
// 1.1.3. runtime.GOMAXPROCS
// Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

// Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

// 我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果，这里举个例子：

// func a() {
//     for i := 1; i < 10; i++ {
//         fmt.Println("A:", i)
//     }
// }

// func b() {
//     for i := 1; i < 10; i++ {
//         fmt.Println("B:", i)
//     }
// }

// func main() {
//     runtime.GOMAXPROCS(1)
//     go a()
//     go b()
//     time.Sleep(time.Second)
// }
// 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。 将逻辑核心数设为2，此时两个任务并行执行，代码如下。

// func a() {
//     for i := 1; i < 10; i++ {
//         fmt.Println("A:", i)
//     }
// }

// func b() {
//     for i := 1; i < 10; i++ {
//         fmt.Println("B:", i)
//     }
// }

// func main() {
//     runtime.GOMAXPROCS(2)
//     go a()
//     go b()
//     time.Sleep(time.Second)
// }
// Go语言中的操作系统线程和goroutine的关系：

// 1.一个操作系统线程对应用户态多个goroutine。
// 2.go程序可以同时使用多个操作系统线程。
// 3.goroutine和OS线程是多对多的关系，即m:n。

//1 Channel
// 1. Channel
// 1.1.1. channel
// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

// 虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

// Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

// Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

// 1.1.2. channel类型
// channel是一种类型，一种引用类型。声明通道类型的格式如下：

//     var 变量 chan 元素类型
// 举几个例子：

//     var ch1 chan int   // 声明一个传递整型的通道
//     var ch2 chan bool  // 声明一个传递布尔型的通道
//     var ch3 chan []int // 声明一个传递int切片的通道
// 1.1.3. 创建channel
// 通道是引用类型，通道类型的空值是nil。

// var ch chan int
// fmt.Println(ch) // <nil>
// 声明的通道后需要使用make函数初始化之后才能使用。

// 创建channel的格式如下：

//     make(chan 元素类型, [缓冲大小])
// channel的缓冲大小是可选的。

// 举几个例子：

// ch4 := make(chan int)
// ch5 := make(chan bool)
// ch6 := make(chan []int)
// 1.1.4. channel操作
// 通道有发送（send）、接收(receive）和关闭（close）三种操作。

// 发送和接收都使用<-符号。

// 现在我们先使用以下语句定义一个通道：

// ch := make(chan int)
// 发送
// 将一个值发送到通道中。

// ch <- 10 // 把10发送到ch中
// 接收
// 从一个通道中接收值。

// x := <- ch // 从ch中接收值并赋值给变量x
// <-ch       // 从ch中接收值，忽略结果
// 关闭
// 我们通过调用内置的close函数来关闭通道。

//     close(ch)
// 关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

// 关闭后的通道有以下特点：

//     1.对一个关闭的通道再发送值就会导致panic。
//     2.对一个关闭的通道进行接收会一直获取值直到通道为空。
//     3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//     4.关闭一个已经关闭的通道会导致panic。
// 1.1.5. 无缓冲的通道

// 无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：

// func main() {
//     ch := make(chan int)
//     ch <- 10
//     fmt.Println("发送成功")
// }
// 上面这段代码能够通过编译，但是执行的时候会出现以下错误：

//     fatal error: all goroutines are asleep - deadlock!

//     goroutine 1 [chan send]:
//     main.main()
//             .../src/github.com/pprof/studygo/day06/channel02/main.go:8 +0x54
// 为什么会出现deadlock错误呢？

// 因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。

// 上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？

// 一种方法是启用一个goroutine去接收值，例如：

// func recv(c chan int) {
//     ret := <-c
//     fmt.Println("接收成功", ret)
// }
// func main() {
//     ch := make(chan int)
//     go recv(ch) // 启用goroutine从通道接收值
//     ch <- 10
//     fmt.Println("发送成功")
// }
// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。

// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

// 1.1.6. 有缓冲的通道
// 解决上面问题的方法还有一种就是使用有缓冲区的通道。

// 我们可以在使用make函数初始化通道的时候为其指定通道的容量，例如：

// func main() {
//     ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
//     ch <- 10
//     fmt.Println("发送成功")
// }
// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

// 我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

// 1.1.7. close()
// 可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）

// package main

// import "fmt"

// func main() {
//     c := make(chan int)
//     go func() {
//         for i := 0; i < 5; i++ {
//             c <- i
//         }
//         close(c)
//     }()
//     for {
//         if data, ok := <-c; ok {
//             fmt.Println(data)
//         } else {
//             break
//         }
//     }
//     fmt.Println("main结束")
// }

// 0
// 1
// 2
// 3
// 4
// main结束

// 1.1.8. 如何优雅的从通道循环取值
// 当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了呢？

// 我们来看下面这个例子：

// // channel 练习
// func main() {
//     ch1 := make(chan int)
//     ch2 := make(chan int)
//     // 开启goroutine将0~100的数发送到ch1中
//     go func() {
//         for i := 0; i < 100; i++ {
//             ch1 <- i
//         }
//         close(ch1)
//     }()
//     // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//     go func() {
//         for {
//             i, ok := <-ch1 // 通道关闭后再取值ok=false
//             if !ok {
//                 break
//             }
//             ch2 <- i * i
//         }
//         close(ch2)
//     }()
//     // 在主goroutine中从ch2中接收值打印
//     for i := range ch2 { // 通道关闭后会退出for range循环
//         fmt.Println(i)
//     }
// }
// 从上面的例子中我们看到有两种方式在接收值的时候判断通道是否被关闭，我们通常使用的是for range的方式。

// 1.1.9. 单向通道
// 有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。

// Go语言中提供了单向通道来处理这种情况。例如，我们把上面的例子改造如下：

// func counter(out chan<- int) {
//     for i := 0; i < 100; i++ {
//         out <- i
//     }
//     close(out)
// }

// func squarer(out chan<- int, in <-chan int) {
//     for i := range in {
//         out <- i * i
//     }
//     close(out)
// }
// func printer(in <-chan int) {
//     for i := range in {
//         fmt.Println(i)
//     }
// }

// func main() {
//     ch1 := make(chan int)
//     ch2 := make(chan int)
//     go counter(ch1)
//     go squarer(ch2, ch1)
//     printer(ch2)
// }
// 其中，

//     1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
//     2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
// 在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的。

// 1.1.10. 通道总结
// channel常见的异常总结，如下图：

// 通道总结

// 注意:关闭已经关闭的channel也会引发panic。

// 1. Goroutine池
// 1.1.1. worker pool（goroutine池）
// 本质上是生产者消费者模型
// 可以有效控制goroutine数量，防止暴涨
// 需求：
// 计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
// 随机生成数字进行计算
// 控制台输出结果如下：

// package main

// import (
//     "fmt"
//     "math/rand"
// )

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

// func main() {
//     // 需要2个管道
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
//打印
// 	job id:2827026 randnum:6305581339897553207 result:89
// job id:2827027 randnum:3016094518478363371 result:79
// job id:2827028 randnum:2302019554749561176 result:77
// job id:2827029 randnum:7184152774698280978 result:103
// job id:2827030 randnum:8750094336704324887 result:88
// job id:2827031 randnum:4970689211519738031 result:84
// job id:2827032 randnum:7541462261557285868 result:92
// job id:2827033 randnum:6529036923274841958 result:93
// job id:2827034 randnum:1505574606295606430 result:74
// job id:2827035 randnum:389025848969403138 result:90
// job id:2827036 randnum:4519465854605333529 result:87
// job id:2827037 randnum:9174186234974291120 result:80
// job id:2827038 randnum:2074139606246420748 result:75
// job id:2827039 randnum:8633768963163929835 result:105
// job id:2827040 randnum:4975341993075920005 result:82
// job id:2827041 randnum:1121316154217137469 result:65
// job id:2827042 randnum:2922257509989915683 result:101
// job id:2827043 randnum:7282684191410259282 result:81
// job id:2827044 randnum:8646541333201342530 result:63
// job id:2827045 randnum:863639991280242677 result:92
// job id:2827046 randnum:317095901630152187 result:68
// job id:2827047 randnum:5272209963828030412 result:73
//等等

// 1. 定时器
// 1.1.1. 定时器

// type Timer
// type Timer struct {
//     C <-chan Time
//     // 内含隐藏或非导出字段
// }
// Timer类型代表单次时间事件。当Timer到期时，当时的时间会被发送给C，除非Timer是被AfterFunc函数创建的。

// func NewTimer
// func NewTimer(d Duration) *Timer
// NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。

// Timer：时间到了，执行只执行1次
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // 1.timer基本使用
//     //timer1 := time.NewTimer(2 * time.Second)
//     //t1 := time.Now()
//     //fmt.Printf("t1:%v\n", t1)
//     //t2 := <-timer1.C
//     //fmt.Printf("t2:%v\n", t2)

// 	t1:2022-03-17 11:32:20.2580423 +0800 CST m=+0.004010401
// t2:2022-03-17 11:32:22.2582508 +0800 CST m=+2.004218901

//     // 2.验证timer只能响应1次
//     //timer2 := time.NewTimer(time.Second)
//     //for {
//     // <-timer2.C
//     // fmt.Println("时间到")
//     //}

// // 	时间到
// // fatal error: all goroutines are asleep - deadlock!

//     // 3.timer实现延时的功能
//     //(1)
//     //time.Sleep(time.Second)
//     //(2)
//     //timer3 := time.NewTimer(2 * time.Second)
//     //<-timer3.C
//     //fmt.Println("2秒到")
//     //(3)
//     //<-time.After(2*time.Second)
//     //fmt.Println("2秒到")

//     // 4.停止定时器
//     //timer4 := time.NewTimer(2 * time.Second)
//     //go func() {
//     // <-timer4.C
//     // fmt.Println("定时器执行了")
//     //}()
//     //b := timer4.Stop()
//     //if b {
//     // fmt.Println("timer4已经关闭")
//     //}

//timer4已经关闭

//     // 5.重置定时器
//     timer5 := time.NewTimer(3 * time.Second)
//     timer5.Reset(1 * time.Second)
//     fmt.Println(time.Now())
//     fmt.Println(<-timer5.C)

//2022-03-17 11:28:32.6393863 +0800 CST m=+0.004034401
// 2022-03-17 11:28:33.640384 +0800 CST m=+1.005032101

//     for {
//     }
// }
// Ticker：时间到了，多次执行
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // 1.获取ticker对象
//     ticker := time.NewTicker(1 * time.Second)
//     i := 0
//     // 子协程
//     go func() {
//         for {
//             //<-ticker.C
//             i++
//             fmt.Println(<-ticker.C)
//             if i == 5 {
//                 //停止
//                 ticker.Stop()
//             }
//         }
//     }()
//     for {
//     }
// }

// 2022-03-17 11:38:53.0633575 +0800 CST m=+1.004703901
// 2022-03-17 11:38:54.0630149 +0800 CST m=+2.004361301
// 2022-03-17 11:38:55.0627367 +0800 CST m=+3.004083201
// 2022-03-17 11:38:56.0633966 +0800 CST m=+4.004743101
// 2022-03-17 11:38:57.0627553 +0800 CST m=+5.004101801

// 1. select
// 1.1.1. select多路复用
// 在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现：

// for{
//     // 尝试从ch1接收值
//     data, ok := <-ch1
//     // 尝试从ch2接收值
//     data, ok := <-ch2
//     …
// }
// 这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

// select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：

//     select {
//     case <-chan1:
//        // 如果chan1成功读到数据，则进行该case处理语句
//     case chan2 <- 1:
//        // 如果成功向chan2写入数据，则进行该case处理语句
//     default:
//        // 如果上面都没有成功，则进入default处理流程
//     }
// select可以同时监听一个或多个channel，直到其中一个channel ready
// package main

// import (
//    "fmt"
//    "time"
// )

// func test1(ch chan string) {
//    time.Sleep(time.Second * 5)
//    ch <- "test1"
// }
// func test2(ch chan string) {
//    time.Sleep(time.Second * 2)
//    ch <- "test2"
// }

// func main() {
//    // 2个管道
//    output1 := make(chan string)
//    output2 := make(chan string)
//    // 跑2个子协程，写数据
//    go test1(output1)
//    go test2(output2)
//    // 用select监控
//    select {
//    case s1 := <-output1:
//       fmt.Println("s1=", s1)
//    case s2 := <-output2:
//       fmt.Println("s2=", s2)
//    }
// }
// 如果多个channel同时ready，则随机选择一个执行
// package main

// import (
//    "fmt"
// )

// func main() {
//    // 创建2个管道
//    int_chan := make(chan int, 1)
//    string_chan := make(chan string, 1)
//    go func() {
//       //time.Sleep(2 * time.Second)
//       int_chan <- 1
//    }()
//    go func() {
//       string_chan <- "hello"
//    }()
//    select {
//    case value := <-int_chan:
//       fmt.Println("int:", value)
//    case value := <-string_chan:
//       fmt.Println("string:", value)
//    }
//    fmt.Println("main结束")
// }

// // // 	string: hello
// // // main结束
// // 或者
// // int: 1
// // main结束

// //select 一次只能选择一个

// 可以用于判断管道是否存满
// package main

// import (
//    "fmt"
//    "time"
// )

// // 判断管道有没有存满
// func main() {
//    // 创建管道
//    output1 := make(chan string, 10)
//    // 子协程写数据
//    go write(output1)
//    // 取数据
//    for s := range output1 {
//       fmt.Println("res:", s)
//       time.Sleep(time.Second)
//    }
// }

// func write(ch chan string) {
//    for {
//       select {
//       // 写数据
//       case ch <- "hello":
//          fmt.Println("write hello")
//       default:
//          fmt.Println("channel full")
//       }
//       time.Sleep(time.Millisecond * 500)
//    }
// }

// 1. 并发安全和锁
// 有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

// 举个例子：

// var x int64
// var wg sync.WaitGroup

// func add() {
//     for i := 0; i < 5000; i++ {
//         x = x + 1
//     }
//     wg.Done()
// }
// func main() {
//     wg.Add(2)
//     go add()
//     go add()
//     wg.Wait()
//     fmt.Println(x)
// }
// 上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。

// 1.1.1. 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：

// var x int64
// var wg sync.WaitGroup
// var lock sync.Mutex

// func add() {
//     for i := 0; i < 5000; i++ {
//         lock.Lock() // 加锁
//         x = x + 1
//         lock.Unlock() // 解锁
//     }
//     wg.Done()
// }
// func main() {
//     wg.Add(2)
//     go add()
//     go add()
//     wg.Wait()
//     fmt.Println(x)
// }
// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。

// 1.1.2. 读写互斥锁
// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

// 读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

// 读写锁示例：

// var (
//     x      int64
//     wg     sync.WaitGroup
//     lock   sync.Mutex
//     rwlock sync.RWMutex
// )

// func write() {
//     // lock.Lock()   // 加互斥锁
//     rwlock.Lock() // 加写锁
//     x = x + 1
//     time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
//     rwlock.Unlock()                   // 解写锁
//     // lock.Unlock()                     // 解互斥锁
//     wg.Done()
// }

// func read() {
//     // lock.Lock()                  // 加互斥锁
//     rwlock.RLock()               // 加读锁
//     time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
//     rwlock.RUnlock()             // 解读锁
//     // lock.Unlock()                // 解互斥锁
//     wg.Done()
// }

// func main() {
//     start := time.Now()
//     for i := 0; i < 10; i++ {
//         wg.Add(1)
//         go write()
//     }

//     for i := 0; i < 1000; i++ {
//         wg.Add(1)
//         go read()
//     }

//     wg.Wait()
//     end := time.Now()
//     fmt.Println(end.Sub(start))
// }
// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

// 1. Sync
// 1.1.1. sync.WaitGroup
// 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：

// 方法名	功能
// (wg * WaitGroup) Add(delta int)	计数器+delta
// (wg *WaitGroup) Done()	计数器-1
// (wg *WaitGroup) Wait()	阻塞直到计数器变为0
// sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

// 我们利用sync.WaitGroup将上面的代码优化一下：

// var wg sync.WaitGroup

// func hello() {
//     defer wg.Done()
//     fmt.Println("Hello Goroutine!")
// }
// func main() {
//     wg.Add(1)
//     go hello() // 启动另外一个goroutine去执行hello函数
//     fmt.Println("main goroutine done!")
//     wg.Wait()
// }
// 需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针。

// 1.1.2. sync.Once
// 说在前面的话：这是一个进阶知识点。

// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。

// sync.Once只有一个Do方法，其签名如下：

// func (o *Once) Do(f func()) {}
// 注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。

// 加载配置文件示例
// 延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践。因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

// var icons map[string]image.Image

// func loadIcons() {
//     icons = map[string]image.Image{
//         "left":  loadIcon("left.png"),
//         "up":    loadIcon("up.png"),
//         "right": loadIcon("right.png"),
//         "down":  loadIcon("down.png"),
//     }
// }

// // Icon 被多个goroutine调用时不是并发安全的
// func Icon(name string) image.Image {
//     if icons == nil {
//         loadIcons()
//     }
//     return icons[name]
// }
// 多个goroutine并发调用Icon函数时不是并发安全的，现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问内存的顺序。loadIcons函数可能会被重排为以下结果：

// func loadIcons() {
//     icons = make(map[string]image.Image)
//     icons["left"] = loadIcon("left.png")
//     icons["up"] = loadIcon("up.png")
//     icons["right"] = loadIcon("right.png")
//     icons["down"] = loadIcon("down.png")
// }
// 在这种情况下就会出现即使判断了icons不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化icons的时候不会被其他的goroutine操作，但是这样做又会引发性能问题。

// 使用sync.Once改造的示例代码如下：

// var icons map[string]image.Image

// var loadIconsOnce sync.Once

// func loadIcons() {
//     icons = map[string]image.Image{
//         "left":  loadIcon("left.png"),
//         "up":    loadIcon("up.png"),
//         "right": loadIcon("right.png"),
//         "down":  loadIcon("down.png"),
//     }
// }

// // Icon 是并发安全的
// func Icon(name string) image.Image {
//     loadIconsOnce.Do(loadIcons)
//     return icons[name]
// }
// sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

// 1.1.3. sync.Map
// Go语言中内置的map不是并发安全的。请看下面的示例：

// var m = make(map[string]int)

// func get(key string) int {
//     return m[key]
// }

// func set(key string, value int) {
//     m[key] = value
// }

// func main() {
//     wg := sync.WaitGroup{}
//     for i := 0; i < 20; i++ {
//         wg.Add(1)
//         go func(n int) {
//             key := strconv.Itoa(n)
//             set(key, n)
//             fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//             wg.Done()
//         }(i)
//     }
//     wg.Wait()
// }
// 上面的代码开启少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误。

// 像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

// var m = sync.Map{}

// func main() {
//     wg := sync.WaitGroup{}
//     for i := 0; i < 20; i++ {
//         wg.Add(1)
//         go func(n int) {
//             key := strconv.Itoa(n)
//             m.Store(key, n)
//             value, _ := m.Load(key)
//             fmt.Printf("k=:%v,v:=%v\n", key, value)
//             wg.Done()
//         }(i)
//     }
//     wg.Wait()
// }

// 1. 原子操作(atomic包)
// 1.1.1. 原子操作
// 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供。

// 1.1.2. atomic包
// 方法	解释
// func LoadInt32(addr int32) (val int32)
// func LoadInt64(addr `int64) (val int64)<br>func LoadUint32(addruint32) (val uint32)<br>func LoadUint64(addruint64) (val uint64)<br>func LoadUintptr(addruintptr) (val uintptr)<br>func LoadPointer(addrunsafe.Pointer`) (val unsafe.Pointer)	读取操作
// func StoreInt32(addr *int32, val int32)
// func StoreInt64(addr *int64, val int64)
// func StoreUint32(addr *uint32, val uint32)
// func StoreUint64(addr *uint64, val uint64)
// func StoreUintptr(addr *uintptr, val uintptr)
// func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)	写入操作
// func AddInt32(addr *int32, delta int32) (new int32)
// func AddInt64(addr *int64, delta int64) (new int64)
// func AddUint32(addr *uint32, delta uint32) (new uint32)
// func AddUint64(addr *uint64, delta uint64) (new uint64)
// func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)	修改操作
// func SwapInt32(addr *int32, new int32) (old int32)
// func SwapInt64(addr *int64, new int64) (old int64)
// func SwapUint32(addr *uint32, new uint32) (old uint32)
// func SwapUint64(addr *uint64, new uint64) (old uint64)
// func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
// func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)	交换操作
// func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
// func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
// func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
// func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
// func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
// func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)	比较并交换操作
// 1.1.3. 示例
// 我们填写一个示例来比较下互斥锁和原子操作的性能。

// var x int64
// var l sync.Mutex
// var wg sync.WaitGroup

// // 普通版加函数
// func add() {
//     // x = x + 1
//     x++ // 等价于上面的操作
//     wg.Done()
// }

// // 互斥锁版加函数
// func mutexAdd() {
//     l.Lock()
//     x++
//     l.Unlock()
//     wg.Done()
// }

// // 原子操作版加函数
// func atomicAdd() {
//     atomic.AddInt64(&x, 1)
//     wg.Done()
// }

// func main() {
//     start := time.Now()
//     for i := 0; i < 10000; i++ {
//         wg.Add(1)
//         // go add()       // 普通版add函数 不是并发安全的
//         // go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
//         go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
//     }
//     wg.Wait()
//     end := time.Now()
//     fmt.Println(x)
//     fmt.Println(end.Sub(start))
// }
// atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。

// 1. GMP 原理与调度
// 1.1.1. 一、Golang “调度器” 的由来？
// (1) 单进程时代不需要调度器
// 我们知道，一切的软件都是跑在操作系统上，真正用来干活 (计算) 的是 CPU。早期的操作系统每个程序就是一个进程，知道一个程序运行完，才能进行下一个进程，就是 “单进程时代”

// 一切的程序只能串行发生。

// 早期的单进程操作系统，面临 2 个问题：

// 单一的执行流程，计算机只能一个任务一个任务处理。

// 进程阻塞所带来的 CPU 时间浪费。

// 那么能不能有多个进程来宏观一起来执行多个任务呢？

// 后来操作系统就具有了最早的并发能力：多进程并发，当一个进程阻塞的时候，切换到另外等待执行的进程，这样就能尽量把 CPU 利用起来，CPU 就不浪费了。

// (2) 多进程 / 线程时代有了调度器需求

// 在多进程 / 多线程的操作系统中，就解决了阻塞的问题，因为一个进程阻塞 cpu 可以立刻切换到其他进程中去执行，而且调度 cpu 的算法可以保证在运行的进程都可以被分配到 cpu 的运行时间片。这样从宏观来看，似乎多个进程是在同时被运行。

// 但新的问题就又出现了，进程拥有太多的资源，进程的创建、切换、销毁，都会占用很长的时间，CPU 虽然利用起来了，但如果进程过多，CPU 有很大的一部分都被用来进行进程调度了。

// 怎么才能提高 CPU 的利用率呢？

// 但是对于 Linux 操作系统来讲，cpu 对进程的态度和线程的态度是一样的。

// 很明显，CPU 调度切换的是进程和线程。尽管线程看起来很美好，但实际上多线程开发设计会变得更加复杂，要考虑很多同步竞争等问题，如锁、竞争冲突等。

// (3) 协程来提高 CPU 利用率
// 多进程、多线程已经提高了系统的并发能力，但是在当今互联网高并发场景下，为每个任务都创建一个线程是不现实的，因为会消耗大量的内存 (进程虚拟内存会占用 4GB [32 位操作系统], 而线程也要大约 4MB)。

// 大量的进程 / 线程出现了新的问题

// 高内存占用
// 调度的高消耗 CPU
// 好了，然后工程师们就发现，其实一个线程分为 “内核态 “线程和” 用户态 “线程。

// 一个 “用户态线程” 必须要绑定一个 “内核态线程”，但是 CPU 并不知道有 “用户态线程” 的存在，它只知道它运行的是一个 “内核态线程”(Linux 的 PCB 进程控制块)。

// 这样，我们再去细化去分类一下，内核线程依然叫 “线程 (thread)”，用户线程叫 “协程 (co-routine)”.

// ​ 看到这里，我们就要开脑洞了，既然一个协程 (co-routine) 可以绑定一个线程 (thread)，那么能不能多个协程 (co-routine) 绑定一个或者多个线程 (thread) 上呢。

// ​ 之后，我们就看到了有 3 中协程和线程的映射关系：

// N:1 关系

// N 个协程绑定 1 个线程，优点就是协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速。但也有很大的缺点，1 个进程的所有协程都绑定在 1 个线程上

// 缺点：

// 某个程序用不了硬件的多核加速能力

// 一旦某协程阻塞，造成线程阻塞，本进程的其他协程都无法执行了，根本就没有并发的能力了。

// 1:1 关系

// 1 个协程绑定 1 个线程，这种最容易实现。协程的调度都由 CPU 完成了，不存在 N:1 缺点，

// 缺点：

// 协程的创建、删除和切换的代价都由 CPU 完成，有点略显昂贵了。

// M:N 关系

// M 个协程绑定 1 个线程，是 N:1 和 1:1 类型的结合，克服了以上 2 种模型的缺点，但实现起来最为复杂。

// ​ 协程跟线程是有区别的，线程由 CPU 调度是抢占式的，协程由用户态调度是协作式的，一个协程让出 CPU 后，才执行下一个协程。

// (4) Go 语言的协程 goroutine
// Go 为了提供更容易使用的并发方法，使用了 goroutine 和 channel。goroutine 来自协程的概念，让一组可复用的函数运行在一组线程之上，即使有协程阻塞，该线程的其他协程也可以被 runtime 调度，转移到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。

// Go 中，协程被称为 goroutine，它非常轻量，一个 goroutine 只占几 KB，并且这几 KB 就足够 goroutine 运行完，这就能在有限的内存空间内支持大量 goroutine，支持了更多的并发。虽然一个 goroutine 的栈只占几 KB，但实际是可伸缩的，如果需要更多内容，runtime 会自动为 goroutine 分配。

// Goroutine 特点：

// 占用内存更小（几 kb）
// 调度更灵活 (runtime 调度)
// (5) 被废弃的 goroutine 调度器
// ​好了，既然我们知道了协程和线程的关系，那么最关键的一点就是调度协程的调度器的实现了。

// Go 目前使用的调度器是 2012 年重新设计的，因为之前的调度器性能存在问题，所以使用 4 年就被废弃了，那么我们先来分析一下被废弃的调度器是如何运作的？

// 大部分文章都是会用 G 来表示 Goroutine，用 M 来表示线程，那么我们也会用这种表达的对应关系。

// 下面我们来看看被废弃的 golang 调度器是如何实现的？

// M 想要执行、放回 G 都必须访问全局 G 队列，并且 M 有多个，即多线程访问同一资源需要加锁进行保证互斥 / 同步，所以全局 G 队列是有互斥锁进行保护的。

// 老调度器有几个缺点：

// 创建、销毁、调度 G 都需要每个 M 获取锁，这就形成了激烈的锁竞争。
// M 转移 G 会造成延迟和额外的系统负载。比如当 G 中包含创建新协程的时候，M 创建了 G’，为了继续执行 G，需要把 G’交给 M’执行，也造成了很差的局部性，因为 G’和 G 是相关的，最好放在 M 上执行，而不是其他 M’。
// 系统调用 (CPU 在 M 之间的切换) 导致频繁的线程阻塞和取消阻塞操作增加了系统开销。
// 1.1.2. 二、Goroutine 调度器的 GMP 模型的设计思想
// 面对之前调度器的问题，Go 设计了新的调度器。

// 在新调度器中，出列 M (thread) 和 G (goroutine)，又引进了 P (Processor)。

// Processor，它包含了运行 goroutine 的资源，如果线程想运行 goroutine，必须先获取 P，P 中还包含了可运行的 G 队列。

// (1) GMP 模型
// 在 Go 中，线程是运行 goroutine 的实体，调度器的功能是把可运行的 goroutine 分配到工作线程上。

// 全局队列（Global Queue）：存放等待运行的 G。
// P 的本地队列：同全局队列类似，存放的也是等待运行的 G，存的数量有限，不超过 256 个。新建 G’时，G’优先加入到 P 的本地队列，如果队列满了，则会把本地队列中一半的 G 移动到全局队列。
// P 列表：所有的 P 都在程序启动时创建，并保存在数组中，最多有 GOMAXPROCS(可配置) 个。
// M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，P 队列为空时，M 也会尝试从全局队列拿一批 G 放到 P 的本地队列，或从其他 P 的本地队列偷一半放到自己 P 的本地队列。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。
// Goroutine 调度器和 OS 调度器是通过 M 结合起来的，每个 M 都代表了 1 个内核线程，OS 调度器负责把内核线程分配到 CPU 的核上执行。

// 有关 P 和 M 的个数问题

// 1、P 的数量：

// 由启动时环境变量 $GOMAXPROCS 或者是由 runtime 的方法 GOMAXPROCS() 决定。这意味着在程序执行的任意时刻都只有 $GOMAXPROCS 个 goroutine 在同时运行。
// 2、M 的数量:

// go 语言本身的限制：go 程序启动时，会设置 M 的最大数量，默认 10000. 但是内核很难支持这么多的线程数，所以这个限制可以忽略。
// runtime/debug 中的 SetMaxThreads 函数，设置 M 的最大数量
// 一个 M 阻塞了，会创建新的 M。
// M 与 P 的数量没有绝对关系，一个 M 阻塞，P 就会去创建或者切换另一个 M，所以，即使 P 的默认数量是 1，也有可能会创建很多个 M 出来。

// P 和 M 何时会被创建

// 1、P 何时创建：在确定了 P 的最大数量 n 后，运行时系统会根据这个数量创建 n 个 P。

// 2、M 何时创建：没有足够的 M 来关联 P 并运行其中的可运行的 G。比如所有的 M 此时都阻塞住了，而 P 中还有很多就绪任务，就会去寻找空闲的 M，而没有空闲的，就会去创建新的 M。

// (2) 调度器的设计策略
// 复用线程：避免频繁的创建、销毁线程，而是对线程的复用。

// 1）work stealing 机制

// ​ 当本线程无可运行的 G 时，尝试从其他线程绑定的 P 偷取 G，而不是销毁线程。

// 2）hand off 机制

// ​ 当本线程因为 G 进行系统调用阻塞时，线程释放绑定的 P，把 P 转移给其他空闲的线程执行。

// 利用并行：GOMAXPROCS 设置 P 的数量，最多有 GOMAXPROCS 个线程分布在多个 CPU 上同时运行。GOMAXPROCS 也限制了并发的程度，比如 GOMAXPROCS = 核数/2，则最多利用了一半的 CPU 核进行并行。

// 抢占：在 coroutine 中要等待一个协程主动让出 CPU 才执行下一个协程，在 Go 中，一个 goroutine 最多占用 CPU 10ms，防止其他 goroutine 被饿死，这就是 goroutine 不同于 coroutine 的一个地方。

// 全局 G 队列：在新的调度器中依然有全局 G 队列，但功能已经被弱化了，当 M 执行 work stealing 从其他 P 偷不到 G 时，它可以从全局 G 队列获取 G。

// (3) go func () 调度流程

// 从上图我们可以分析出几个结论：

// ​ 1、我们通过 go func () 来创建一个 goroutine；

// ​ 2、有两个存储 G 的队列，一个是局部调度器 P 的本地队列、一个是全局 G 队列。新创建的 G 会先保存在 P 的本地队列中，如果 P 的本地队列已经满了就会保存在全局的队列中；

// ​ 3、G 只能运行在 M 中，一个 M 必须持有一个 P，M 与 P 是 1：1 的关系。M 会从 P 的本地队列弹出一个可执行状态的 G 来执行，如果 P 的本地队列为空，就会想其他的 MP 组合偷取一个可执行的 G 来执行；

// ​ 4、一个 M 调度 G 执行的过程是一个循环机制；

// ​ 5、当 M 执行某一个 G 时候如果发生了 syscall 或则其余阻塞操作，M 会阻塞，如果当前有一些 G 在执行，runtime 会把这个线程 M 从 P 中摘除 (detach)，然后再创建一个新的操作系统的线程 (如果有空闲的线程可用就复用空闲线程) 来服务于这个 P；

// ​ 6、当 M 系统调用结束时候，这个 G 会尝试获取一个空闲的 P 执行，并放入到这个 P 的本地队列。如果获取不到 P，那么这个线程 M 变成休眠状态， 加入到空闲线程中，然后这个 G 会被放入全局队列中。

// (4) 调度器的生命周期

// 特殊的 M0 和 G0

// M0

// M0 是启动程序后的编号为 0 的主线程，这个 M 对应的实例会在全局变量 runtime.m0 中，不需要在 heap 上分配，M0 负责执行初始化操作和启动第一个 G， 在之后 M0 就和其他的 M 一样了。

// G0

// G0 是每次启动一个 M 都会第一个创建的 gourtine，G0 仅用于负责调度的 G，G0 不指向任何可执行的函数，每个 M 都会有一个自己的 G0。在调度或系统调用时会使用 G0 的栈空间，全局变量的 G0 是 M0 的 G0。

// 我们来跟踪一段代码

// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello world")
// }
// 接下来我们来针对上面的代码对调度器里面的结构做一个分析。

// 也会经历如上图所示的过程：

// 1.runtime 创建最初的线程 m0 和 goroutine g0，并把 2 者关联。
// 2.调度器初始化：初始化 m0、栈、垃圾回收，以及创建和初始化由 GOMAXPROCS 个 P 构成的 P 列表。
// 3.示例代码中的 main 函数是 main.main，runtime 中也有 1 个 main 函数 ——runtime.main，代码经过编译后，runtime.main 会调用 main.main，程序启动时会为 runtime.main 创建 goroutine，称它为 main goroutine 吧，然后把 main goroutine 加入到 P 的本地队列。
// 4.启动 m0，m0 已经绑定了 P，会从 P 的本地队列获取 G，获取到 main goroutine。
// 5.G 拥有栈，M 根据 G 中的栈信息和调度信息设置运行环境
// 6.M 运行 G
// 7.G 退出，再次回到 M 获取可运行的 G，这样重复下去，直到 main.main 退出，runtime.main 执行 Defer 和 Panic 处理，或调用 runtime.exit 退出程序。
// 调度器的生命周期几乎占满了一个 Go 程序的一生，runtime.main 的 goroutine 执行之前都是为调度器做准备工作，runtime.main 的 goroutine 运行，才是调度器的真正开始，直到 runtime.main 结束而结束。

// (5) 可视化 GMP 编程
// 有 2 种方式可以查看一个程序的 GMP 的数据。

// 方式 1：go tool trace

// trace 记录了运行时的信息，能提供可视化的 Web 页面。

// 简单测试代码：main 函数创建 trace，trace 会运行在单独的 goroutine 中，然后 main 打印”Hello World” 退出。

// trace.go

// package main

// import (
//     "os"
//     "fmt"
//     "runtime/trace"
// )

// func main() {

//     //创建trace文件
//     f, err := os.Create("trace.out")
//     if err != nil {
//         panic(err)
//     }

//     defer f.Close()

//     //启动trace goroutine
//     err = trace.Start(f)
//     if err != nil {
//         panic(err)
//     }
//     defer trace.Stop()

//     //main
//     fmt.Println("Hello World")
// }
// 运行程序

// $ go run trace.go
// Hello World
// 会得到一个 trace.out 文件，然后我们可以用一个工具打开，来分析这个文件。

// $ go tool trace trace.out
// 2020/02/23 10:44:11 Parsing trace...
// 2020/02/23 10:44:11 Splitting trace...
// 2020/02/23 10:44:11 Opening browser. Trace viewer is listening on http://127.0.0.1:33479
// 我们可以通过浏览器打开 http://127.0.0.1:33479 网址，点击 view trace 能够看见可视化的调度流程。

// G 信息

// 点击 Goroutines 那一行可视化的数据条，我们会看到一些详细的信息。

// 一共有两个G在程序中，一个是特殊的G0，是每个M必须有的一个初始化的G，这个我们不必讨论。

// 其中 G1 应该就是 main goroutine (执行 main 函数的协程)，在一段时间内处于可运行和运行的状态。

// M 信息

// 点击 Threads 那一行可视化的数据条，我们会看到一些详细的信息。

// 一共有两个 M 在程序中，一个是特殊的 M0，用于初始化使用，这个我们不必讨论。

// P 信息

// G1 中调用了 main.main，创建了 trace goroutine g18。G1 运行在 P1 上，G18 运行在 P0 上。

// 这里有两个 P，我们知道，一个 P 必须绑定一个 M 才能调度 G。

// 我们在来看看上面的 M 信息。

// 我们会发现，确实 G18 在 P0 上被运行的时候，确实在 Threads 行多了一个 M 的数据，点击查看如下：

// 多了一个 M2 应该就是 P0 为了执行 G18 而动态创建的 M2.

// 方式 2：Debug trace

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     for i := 0; i < 5; i++ {
//         time.Sleep(time.Second)
//         fmt.Println("Hello World")
//     }
// }
// 编译

// $ go build trace2.go
// 通过 Debug 方式运行

// $ GODEBUG=schedtrace=1000 ./trace2
// SCHED 0ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=1 idlethreads=1 runqueue=0 [0 0]
// Hello World
// SCHED 1003ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
// Hello World
// SCHED 2014ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
// Hello World
// SCHED 3015ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
// Hello World
// SCHED 4023ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
// Hello World
// SCHED：调试信息输出标志字符串，代表本行是 goroutine 调度器的输出；
// 0ms：即从程序启动到输出这行日志的时间；
// gomaxprocs: P 的数量，本例有 2 个 P, 因为默认的 P 的属性是和 cpu 核心数量默认一致，当然也可以通过 GOMAXPROCS 来设置；
// idleprocs: 处于 idle 状态的 P 的数量；通过 gomaxprocs 和 idleprocs 的差值，我们就可知道执行 go 代码的 P 的数量；
// threads: os threads/M 的数量，包含 scheduler 使用的 m 数量，加上 runtime 自用的类似 sysmon 这样的 thread 的数量；
// spinningthreads: 处于自旋状态的 os thread 数量；
// idlethread: 处于 idle 状态的 os thread 的数量；
// runqueue=0： Scheduler 全局队列中 G 的数量；
// [0 0]: 分别为 2 个 P 的 local queue 中的 G 的数量。
// 1.1.3. 三、Go 调度器调度场景过程全解析
// (1) 场景 1

// P 拥有 G1，M1 获取 P 后开始运行 G1，G1 使用 go func() 创建了 G2，为了局部性 G2 优先加入到 P1 的本地队列。

// (2) 场景 2

// G1 运行完成后 (函数：goexit)，M 上运行的 goroutine 切换为 G0，G0 负责调度时协程的切换（函数：schedule）。从 P 的本地队列取 G2，从 G0 切换到 G2，并开始运行 G2 (函数：execute)。实现了线程 M1 的复用。

// (3) 场景 3

// 假设每个 P 的本地队列只能存 3 个 G。G2 要创建了 6 个 G，前 3 个 G（G3, G4, G5）已经加入 p1 的本地队列，p1 本地队列满了。

// (4) 场景 4

// G2 在创建 G7 的时候，发现 P1 的本地队列已满，需要执行负载均衡 (把 P1 中本地队列中前一半的 G，还有新创建 G 转移到全局队列)

// （实现中并不一定是新的 G，如果 G 是 G2 之后就执行的，会被保存在本地队列，利用某个老的 G 替换新 G 加入全局队列）

// 这些 G 被转移到全局队列时，会被打乱顺序。所以 G3,G4,G7 被转移到全局队列。

// (5) 场景 5

// G2 创建 G8 时，P1 的本地队列未满，所以 G8 会被加入到 P1 的本地队列。

// G8 加入到 P1 点本地队列的原因还是因为 P1 此时在与 M1 绑定，而 G2 此时是 M1 在执行。所以 G2 创建的新的 G 会优先放置到自己的 M 绑定的 P 上。

// (6) 场景 6

// 规定：在创建 G 时，运行的 G 会尝试唤醒其他空闲的 P 和 M 组合去执行。

// 假定 G2 唤醒了 M2，M2 绑定了 P2，并运行 G0，但 P2 本地队列没有 G，M2 此时为自旋线程（没有 G 但为运行状态的线程，不断寻找 G）。

// (7) 场景 7

// M2 尝试从全局队列 (简称 “GQ”) 取一批 G 放到 P2 的本地队列（函数：findrunnable()）。M2 从全局队列取的 G 数量符合下面的公式：

// n = min(len(GQ)/GOMAXPROCS + 1, len(GQ/2))

// 至少从全局队列取 1 个 g，但每次不要从全局队列移动太多的 g 到 p 本地队列，给其他 p 留点。这是从全局队列到 P 本地队列的负载均衡。

// 假定我们场景中一共有 4 个 P（GOMAXPROCS 设置为 4，那么我们允许最多就能用 4 个 P 来供 M 使用）。所以 M2 只从能从全局队列取 1 个 G（即 G3）移动 P2 本地队列，然后完成从 G0 到 G3 的切换，运行 G3。

// (8) 场景 8

// 假设 G2 一直在 M1 上运行，经过 2 轮后，M2 已经把 G7、G4 从全局队列获取到了 P2 的本地队列并完成运行，全局队列和 P2 的本地队列都空了，如场景 8 图的左半部分。

// 全局队列已经没有 G，那 m 就要执行 work stealing (偷取)：从其他有 G 的 P 哪里偷取一半 G 过来，放到自己的 P 本地队列。P2 从 P1 的本地队列尾部取一半的 G，本例中一半则只有 1 个 G8，放到 P2 的本地队列并执行。

// (9) 场景 9

// G1 本地队列 G5、G6 已经被其他 M 偷走并运行完成，当前 M1 和 M2 分别在运行 G2 和 G8，M3 和 M4 没有 goroutine 可以运行，M3 和 M4 处于自旋状态，它们不断寻找 goroutine。

// 为什么要让 m3 和 m4 自旋，自旋本质是在运行，线程在运行却没有执行 G，就变成了浪费 CPU. 为什么不销毁现场，来节约 CPU 资源。因为创建和销毁 CPU 也会浪费时间，我们希望当有新 goroutine 创建时，立刻能有 M 运行它，如果销毁再新建就增加了时延，降低了效率。当然也考虑了过多的自旋线程是浪费 CPU，所以系统中最多有 GOMAXPROCS 个自旋的线程 (当前例子中的 GOMAXPROCS=4，所以一共 4 个 P)，多余的没事做线程会让他们休眠。

// (10) 场景 10

// ​ 假定当前除了 M3 和 M4 为自旋线程，还有 M5 和 M6 为空闲的线程 (没有得到 P 的绑定，注意我们这里最多就只能够存在 4 个 P，所以 P 的数量应该永远是 M>=P, 大部分都是 M 在抢占需要运行的 P)，G8 创建了 G9，G8 进行了阻塞的系统调用，M2 和 P2 立即解绑，P2 会执行以下判断：如果 P2 本地队列有 G、全局队列有 G 或有空闲的 M，P2 都会立马唤醒 1 个 M 和它绑定，否则 P2 则会加入到空闲 P 列表，等待 M 来获取可用的 p。本场景中，P2 本地队列有 G9，可以和其他空闲的线程 M5 绑定。

// (11) 场景 11

// G8 创建了 G9，假如 G8 进行了非阻塞系统调用。

// ​ M2 和 P2 会解绑，但 M2 会记住 P2，然后 G8 和 M2 进入系统调用状态。当 G8 和 M2 退出系统调用时，会尝试获取 P2，如果无法获取，则获取空闲的 P，如果依然没有，G8 会被记为可运行状态，并加入到全局队列，M2 因为没有 P 的绑定而变成休眠状态 (长时间休眠等待 GC 回收销毁)。

// 1.1.4. 四、小结
// 总结，Go 调度器很轻量也很简单，足以撑起 goroutine 的调度工作，并且让 Go 具有了原生（强大）并发的能力。Go 调度本质是把大量的 goroutine 分配到少量线程上去执行，并利用多核并行，实现更强大的并发。

// 转自公众号：刘丹冰Aceld

// 爬虫小案例
// 1. 爬虫小案例
// 1.1.1. 爬虫步骤
// 明确目标（确定在哪个网站搜索）
// 爬（爬下内容）
// 取（筛选想要的）
// 处理数据（按照你的想法去处理）
// package main

// import (
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "regexp"
// )

// //这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
// var (
//     // \d是数字
//     reQQEmail = `(\d+)@qq.com`
// )

// // 爬邮箱
// func GetEmail() {
//     // 1.去网站拿数据
//     resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
//     HandleError(err, "http.Get url")
//     defer resp.Body.Close()
//     // 2.读取页面内容
//     pageBytes, err := ioutil.ReadAll(resp.Body)
//     HandleError(err, "ioutil.ReadAll")
//     // 字节转字符串
//     pageStr := string(pageBytes)
//     //fmt.Println(pageStr)
//     // 3.过滤数据，过滤qq邮箱
//     re := regexp.MustCompile(reQQEmail)
//     // -1代表取全部
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     //fmt.Println(results)

//     // 遍历结果
//     for _, result := range results {
//         fmt.Println("email:", result[0])
//         fmt.Println("qq:", result[1])
//     }
// }

// // 处理异常
// func HandleError(err error, why string) {
//     if err != nil {
//         fmt.Println(why, err)
//     }
// }
// func main() {
//     GetEmail()
// }
// 1.1.2. 正则表达式
// 文档：https://studygolang.com/pkgdoc
// API
// re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
// ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部
// 爬邮箱
// 方法抽取
// 爬超链接
// 爬手机号
// http://www.zhaohaowang.com/ 如果连接失效了自己找一个有手机号的就好了
// 爬身份证号
// http://henan.qq.com/a/20171107/069413.htm 如果连接失效了自己找一个就好了
// 爬图片链接
// package main

// import (
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "regexp"
// )

// var (
//     // w代表大小写字母+数字+下划线
//     reEmail = `\w+@\w+\.\w+`
//     // s?有或者没有s
//     // +代表出1次或多次
//     //\s\S各种字符
//     // +?代表贪婪模式
//     reLinke  = `href="(https?://[\s\S]+?)"`
//     rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
//     reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
//     reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
// )

// // 处理异常
// func HandleError(err error, why string) {
//     if err != nil {
//         fmt.Println(why, err)
//     }
// }

// func GetEmail2(url string) {
//     pageStr := GetPageStr(url)
//     re := regexp.MustCompile(reEmail)
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     for _, result := range results {
//         fmt.Println(result)
//     }
// }

// // 抽取根据url获取内容
// func GetPageStr(url string) (pageStr string) {
//     resp, err := http.Get(url)
//     HandleError(err, "http.Get url")
//     defer resp.Body.Close()
//     // 2.读取页面内容
//     pageBytes, err := ioutil.ReadAll(resp.Body)
//     HandleError(err, "ioutil.ReadAll")
//     // 字节转字符串
//     pageStr = string(pageBytes)
//     return pageStr
// }

// func main() {
//     // 2.抽取的爬邮箱
//     // GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
//     // 3.爬链接
//     //GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
//     // 4.爬手机号
//     //GetPhone("https://www.zhaohaowang.com/")
//     // 5.爬身份证号
//     //GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
//     // 6.爬图片
//     // GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
// }

// func GetIdCard(url string) {
//     pageStr := GetPageStr(url)
//     re := regexp.MustCompile(reIdcard)
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     for _, result := range results {
//         fmt.Println(result)
//     }
// }

// // 爬链接
// func GetLink(url string) {
//     pageStr := GetPageStr(url)
//     re := regexp.MustCompile(reLinke)
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     for _, result := range results {
//         fmt.Println(result[1])
//     }
// }

// //爬手机号
// func GetPhone(url string) {
//     pageStr := GetPageStr(url)
//     re := regexp.MustCompile(rePhone)
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     for _, result := range results {
//         fmt.Println(result)
//     }
// }

// func GetImg(url string) {
//     pageStr := GetPageStr(url)
//     re := regexp.MustCompile(reImg)
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     for _, result := range results {
//         fmt.Println(result[0])
//     }
// }
// 1.1.3. 并发爬取美图
// 下面的两个是即将要爬的网站，如果网址失效自己换一个就好了

// https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/1.html
// package main

// import (
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "regexp"
//     "strconv"
//     "strings"
//     "sync"
//     "time"
// )

// func HandleError(err error, why string) {
//     if err != nil {
//         fmt.Println(why, err)
//     }
// }

// // 下载图片，传入的是图片叫什么
// func DownloadFile(url string, filename string) (ok bool) {
//     resp, err := http.Get(url)
//     HandleError(err, "http.get.url")
//     defer resp.Body.Close()
//     bytes, err := ioutil.ReadAll(resp.Body)
//     HandleError(err, "resp.body")
//     filename = "E:/topgoer.com/src/github.com/student/3.0/img/" + filename
//     // 写出数据
//     err = ioutil.WriteFile(filename, bytes, 0666)
//     if err != nil {
//         return false
//     } else {
//         return true
//     }
// }

// // 并发爬思路：
// // 1.初始化数据管道
// // 2.爬虫写出：26个协程向管道中添加图片链接
// // 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// // 4.下载协程：从管道里读取链接并下载

// var (
//     // 存放图片链接的数据管道
//     chanImageUrls chan string
//     waitGroup     sync.WaitGroup
//     // 用于监控协程
//     chanTask chan string
//     reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
// )

// func main() {
//     // myTest()
//     // DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201909/10242/e5794daf58_4.jpg", "1.jpg")

//     // 1.初始化管道
//     chanImageUrls = make(chan string, 1000000)
//     chanTask = make(chan string, 26)
//     // 2.爬虫协程
//     for i := 1; i < 27; i++ {
//         waitGroup.Add(1)
//         go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
//     }
//     // 3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
//     waitGroup.Add(1)
//     go CheckOK()
//     // 4.下载协程：从管道中读取链接并下载
//     for i := 0; i < 5; i++ {
//         waitGroup.Add(1)
//         go DownloadImg()
//     }
//     waitGroup.Wait()
// }

// // 下载图片
// func DownloadImg() {
//     for url := range chanImageUrls {
//         filename := GetFilenameFromUrl(url)
//         ok := DownloadFile(url, filename)
//         if ok {
//             fmt.Printf("%s 下载成功\n", filename)
//         } else {
//             fmt.Printf("%s 下载失败\n", filename)
//         }
//     }
//     waitGroup.Done()
// }

// // 截取url名字
// func GetFilenameFromUrl(url string) (filename string) {
//     // 返回最后一个/的位置
//     lastIndex := strings.LastIndex(url, "/")
//     // 切出来
//     filename = url[lastIndex+1:]
//     // 时间戳解决重名
//     timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
//     filename = timePrefix + "_" + filename
//     return
// }

// // 任务统计协程
// func CheckOK() {
//     var count int
//     for {
//         url := <-chanTask
//         fmt.Printf("%s 完成了爬取任务\n", url)
//         count++
//         if count == 26 {
//             close(chanImageUrls)
//             break
//         }
//     }
//     waitGroup.Done()
// }

// // 爬图片链接到管道
// // url是传的整页链接
// func getImgUrls(url string) {
//     urls := getImgs(url)
//     // 遍历切片里所有链接，存入数据管道
//     for _, url := range urls {
//         chanImageUrls <- url
//     }
//     // 标识当前协程完成
//     // 每完成一个任务，写一条数据
//     // 用于监控协程知道已经完成了几个任务
//     chanTask <- url
//     waitGroup.Done()
// }

// // 获取当前页图片链接
// func getImgs(url string) (urls []string) {
//     pageStr := GetPageStr(url)
//     re := regexp.MustCompile(reImg)
//     results := re.FindAllStringSubmatch(pageStr, -1)
//     fmt.Printf("共找到%d条结果\n", len(results))
//     for _, result := range results {
//         url := result[0]
//         urls = append(urls, url)
//     }
//     return
// }

// // 抽取根据url获取内容
// func GetPageStr(url string) (pageStr string) {
//     resp, err := http.Get(url)
//     HandleError(err, "http.Get url")
//     defer resp.Body.Close()
//     // 2.读取页面内容
//     pageBytes, err := ioutil.ReadAll(resp.Body)
//     HandleError(err, "ioutil.ReadAll")
//     // 字节转字符串
//     pageStr = string(pageBytes)
//     return pageStr
// }
// 作者：孙建超
