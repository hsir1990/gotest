//单例加锁的目的是定格，竞争条件，竞争以后抢这些资源，抢到以后就上锁，为啥只锁写不写读呢？
//增加年龄，填写了200次，有问题不会得到200
package main

import (
	"fmt"
	"sync"
)

var (
	p    *Pet
	once sync.Once
)

func init() {
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
	mux  sync.Mutex
}

func (p *Pet) SetName(n string) {
	p.mux.Lock()
	defer p.mux.Unlock()
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
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			IncrementAge()
		}()

		go func() {
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
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
