package main

import (
	"fmt"
	"container/ring"
	"container/list"
	"container/heap"
	"time"
)
	
//通过环统计最近10次的平均值
func tenRing(){
	const WINDOW_SIZE = 10
	ring := ring.New(WINDOW_SIZE)
	for i:=0; i<100;i++{
		ring.Value = i
		ring = ring.Next()  //注意这处的写法
	}
	var sum int = 0
	ring.Do(func(i interface{}){ //通过Do（）来遍历ring，内部实际上调用了NExt（），会将value的值赋给i
		sum += i.(int)
		// fmt.Printf("%T  /n",i)
	})
	fmt.Printf("100的平均值是：%.1f  \n",float64(sum)/float64(WINDOW_SIZE))
}

//费波那数列
//1 2 3 4 5 6 7
//0 1 1 2 3 5 8
var a1 int= 0
func FibonacciWithRecursion(n int) int {
	if n ==1 || n ==2{
		return n-1
	}
	a1++
	return FibonacciWithRecursion(n-1)+FibonacciWithRecursion(n-2)
}

//用list模拟栈来使用
func FibonacciWithStack(n int) int{  //go语言中list包含了stack的功能
	Stack := list.New()    //func New() *List    New创建一个链表。
	if n ==1 || n ==2 {
		return n-1
	}
	Stack.PushBack(0)
	Stack.PushBack(1) //PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。  func (l *List) PushBack(v interface{}) *Element
	for i:= 2; i < n; i++ {
		a := Stack.Back() // func (l *List) Back() *Element   Back返回链表最后一个元素或nil。
		Stack.Remove(a)  //func (l *List) Remove(e *Element) interface{}    Remove删除链表中的元素e，并返回e.Value。
		b := Stack.Back()
		Stack.Remove(b)
		Stack.PushBack(a.Value.(int))
		Stack.PushBack(a.Value.(int)+b.Value.(int))
	}
	a := Stack.Back()
	result := Stack.Remove(a)
	return result.(int)


}


//用堆实现超时缓存
//每个节点可以用struct表示，然后放入到sclie里面，再转换成堆
var cache map[int]*Node

const LIFE = 10

func init(){
	cache = make(map[int]*Node)
}

type Node struct{
	Deadline int64
	Key int
}

type TimeoutHeap []*Node  //等于是重定义[]*Node

func (pq TimeoutHeap) Len() int{
	return len(pq)
}

func (pq TimeoutHeap) Less(i, j int) bool{
	return pq[i].Deadline < pq[i].Deadline
}

func (pq TimeoutHeap) Swap(i , j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TimeoutHeap) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *TimeoutHeap) Pop() interface{}{
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func testTimeoutCache(){
	pq := make(TimeoutHeap, 0, 5)
	heap.Init(&pq)  //从无序状态构建堆

	for i := 0; i < 10 ; i++{
		node := &Node{Deadline:time.Now().UnixNano() + LIFE, Key:i}
		cache[i] = node //放入缓存
		heap.Push(&pq, node) //同时放入堆
		time.Sleep(20 * time.Millisecond)
	}

	ticker := time.NewTicker(5* time.Millisecond) //每隔5毫秒检查一下是否有元素到期

	for{
		<- ticker.C
		for{
			currentTimestamp := time.Now().UnixNano() //获取当前时间
			if pq.Len() <=0{
				break
			}
			first := pq[0] //取得小根堆顶元素
			if currentTimestamp < first.Deadline{
				break
			}else{//当前时间比堆顶元素小，说明堆顶已到期，需要从缓存里删除
				delete(cache, first.Key)
				heap.Pop(&pq)
				fmt.Printf("delete  %v \n", *first)
			}
		}
	}
}

func main(){
	tenRing()
	//递归容易让函数重复算好几次，比如下面f(5)被算了1次， f(4)被算了1次， 但是f(3)被重复算了2次，f(2)也被重复算了2次，f（1）也是
	fmt.Println(FibonacciWithRecursion(5))
	fmt.Println("al是递归函数算的次数：",a1)

	fmt.Println("用栈的形式搞定：",FibonacciWithStack(7))
	fmt.Println("11")
	testTimeoutCache()
}


// container
// 英 [kənˈteɪnə(r)]   美 [kənˈteɪnər]  
// n.
// 容器;集装箱;货柜

// compress
// 英 [kəmˈpres , ˈkɒmpres]   美 [kəmˈpres , ˈkɑːmpres]  
// v.
// (被)压紧;精简;浓缩;压缩;压缩(文件等)
// n.
// (止血、减痛等的)敷布，压布