//单例加锁的目的是定格，竞争条件，竞争以后抢这些资源，抢到以后就上锁，为啥只锁写不写读呢？
//增加年龄，填写了200次，有问题不会得到200
package main

import (
	"fmt"
	"sync"
)

var (
	p    *Pet
	once sync.Once //Once是只执行一次动作的对象。例子可见下面
)

func init() {
	//Do方法当且仅当第一次被调用时才执行括号里面的函数
	once.Do(func() { //保证程序初始级别，就运行的函数
		p = &Pet{}
	})
}

func GetInstance() *Pet {
	return p
}

type Pet struct {
	name string
	age  int
	mux  sync.Mutex //定义加互斥锁
}

func (p *Pet) SetName(n string) {
	p.mux.Lock()         //加互斥锁
	defer p.mux.Unlock() //直到解锁才解开
	p.name = n
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *Pet) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *Pet) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}

func IncrementAge() {
	p := GetInstance()
	p.IncrementAge()
}

func IncrementAge2() {
	p := GetInstance()
	p.IncrementAge()
}
func main() {
	wg := sync.WaitGroup{} //用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。
	wg.Add(200)            //加200个线程数量

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done() //执行减100次
			IncrementAge()
		}()

		go func() {
			defer wg.Done() //执行减100次
			IncrementAge2()
		}()
	}
	wg.Wait() //一直堵塞着，直到所有线程结束才行
	p := GetInstance()
	age := p.GetAge()
	fmt.Println(age) //200
}

// singleton
// 英 [ˈsɪŋɡltən]   美 [ˈsɪŋɡltən]
// n.
// (所提及的)单项物，单个的人;单身男子(或女子);(非孪生的)单生儿，单生幼畜
// 复数： singletons

// sync
// 英 [sɪŋk]   美 [sɪŋk]
// n.
// synchronization 的缩略词
// v.
// synchronize 的缩略词

// Mutex
// 信号量;互斥量;互斥体;互斥锁;互斥

// Increment
// 英 [ˈɪŋkrəmənt]   美 [ˈɪŋkrəmənt]
// 增量;递增;自增;增加;增量方式

// instance
// 英 [ˈɪnstəns]   美 [ˈɪnstəns]
// n.
// 例子;事例;实例
// vt.
// 举…为例
// 第三人称单数： instances复数： instances现在分词： instancing过去式： instanced过去分词： instanced

// //sync.Once的例子
// var once sync.Once
// onceBody := func() {
//     fmt.Println("Only once")
// }
// done := make(chan bool)
// for i := 0; i < 10; i++ {
//     go func() {
//         once.Do(onceBody)
//         done <- true
//     }()
// }
// for i := 0; i < 10; i++ {
//     <-done
// }
