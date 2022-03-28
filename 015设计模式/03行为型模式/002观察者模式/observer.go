//观察者模式,观察者行为有通知,订阅者有添加观察者,移除观察者,还有通知事件
//比如我们很多人订阅了公众号,然后大V发通知,使用方法循环给用户订阅者
package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data int
}

type Observer interface {
	NotifyCallback(envet Event)
}

type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

type eventObserver struct {
	ID   int
	Time time.Time
}

type eventSubject struct {
	Observers sync.Map
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Recieved:%d after %v\n", event.Data, time.Since(e.Time))
}
func (e *eventSubject) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}
func (e *eventSubject) RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}

func (e *eventSubject) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}

func main() {
	// 	for x := range Fib(20) {
	// 		fmt.Println(x)
	// 	}
	// // 	0
	// // 1
	// // 1
	// // 2
	// // 3
	// // 5
	// // 8
	// // 13

	n := eventSubject{Observers: sync.Map{}}

	obs1 := eventObserver{ID: 1, Time: time.Now()}
	obs2 := eventObserver{ID: 2, Time: time.Now()}

	n.AddListener(obs1)
	n.AddListener(obs2)

	for x := range Fib(10) {
		n.Notify(Event{Data: x})
	}

	// 	Recieved:0 after 0s
	// Recieved:0 after 0s
	// Recieved:1 after 0s
	// Recieved:1 after 0s
	// Recieved:1 after 0s
	// Recieved:1 after 0s
	// Recieved:2 after 0s
	// Recieved:2 after 0s
	// Recieved:3 after 0s
	// Recieved:3 after 0s
	// Recieved:5 after 0s
	// Recieved:5 after 0s
	// Recieved:8 after 0s
	// Recieved:8 after 0s

}

// observer
// 英 [əbˈzɜːvə(r)]   美 [əbˈzɜːrvər]
// n.
// 观察者;观察员;观察家;观测者;目击者;旁听者;评论员

// notify
// 英 [ˈnəʊtɪfaɪ]   美 [ˈnoʊtɪfaɪ]
// vt.
// 通知;(正式)通报
// 第三人称单数： notifies现在分词： notifying过去式： notified过去分词： notified
// 记忆技巧：not 知道 + ify 使… → 使…知道 → 通知

// subject
// 英 [ˈsʌbdʒɪkt , səbˈdʒekt]   美 [ˈsʌbdʒɪkt , səbˈdʒekt]
// n.
// 主题;题目;话题;题材;问题;学科;科目;课程;表现对象;绘画(或拍摄)题材;接受试验者;主语;国民，臣民
// adj.
// 可能受…影响的;易遭受…的;取决于;视…而定;受…支配;服从于;受异族统治的
// vt.
// 使臣服;使顺从;(尤指)压服
// 第三人称单数： subjects复数： subjects现在分词： subjecting过去式： subjected过去分词： subjected
// 派生词： subjection n.
// 记忆技巧：sub 下 + ject 投掷，扔 → 扔下去〔让大家讨论〕→ 主题

// event
// 英 [ɪˈvent]   美 [ɪˈvent]
// n.
// 事件;发生的事情;(尤指)重要事情，大事;公开活动;社交场合;(体育运动的)比赛项目
// 复数： events
// 记忆技巧：e 出 + vent 来 → 出来的事 → 事件
