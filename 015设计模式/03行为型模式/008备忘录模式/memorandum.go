// 备忘录模式，就是你有事情记录在小本本上，然后呢，你忘记了，再然后把小本本拿出来，就能知道什么事情什么情况了，如果听说过就打1，没听说过就打0
//比如你刚买进股票时7元，然后翻倍在翻倍，28元了，你记录在你的小本本上，然后过了一会，跌了一半，变成了14元，这时候，你忘记了，这个最高的时候，是多少钱来，这时候你就把你的备忘录拿出来，看了一看就是28元
package main

import (
	"fmt"
)

type Memento struct {
	state int
}

func NewMemento(value int) *Memento {
	return &Memento{value}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value: value}
}

func (n *Number) Double() {
	n.value = 2 * n.value
}

func (n *Number) Half() {
	n.value /= 2
}

func (n *Number) Value() int {
	return n.value
}

func (n *Number) CreateMemento() *Memento {
	return NewMemento(n.value)
}

func (n *Number) ReinstateMemeto(memento *Memento) {
	n.value = memento.state
}

func main() {
	n := NewNumber(7)
	n.Double()
	n.Double()
	memento := n.CreateMemento() //生成备忘录记录
	n.Half()
	fmt.Println(n.value) //14
	n.ReinstateMemeto(memento)
	fmt.Println(n.value)
	//28

}

// memorandum
// 英 [ˌmeməˈrændəm]   美 [ˌmeməˈrændəm]
// n.
// 备忘录;协议备忘录;建议书;报告

// memento
// 英 [məˈmentəʊ]   美 [məˈmentoʊ]
// n.
// 纪念品
