//中介模式创建一个中介，中介结构体种包含两个客户，然后客户会把中介设置到自己的属性里，
//一个客户问bill在哪，然后调用中介通讯的方法去联系bill，bill回答以后，接着调用中介的方法去调用，自己去响应回复
//通过买卖房屋的案例
package main

import (
	"fmt"
)

type Mediator interface {
	Communicate(who string)
}

type WildStallion interface {
	SetMediator(mediator Mediator)
}

type Bill struct {
	mediator Mediator
}

func (b *Bill) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *Bill) Respond() {
	fmt.Println("Bill: What")
	b.mediator.Communicate("bill")
}

type Ted struct {
	mediator Mediator
}

func (t *Ted) Talk() {
	fmt.Println("Ted:Bill")
	t.mediator.Communicate("Ted")
}

func (t *Ted) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

func (t *Ted) Respond() {
	// fmt.Println("Ted: Strange things are afoot at the Circle K.")
	fmt.Println("Ted: how much")
}

type ConcreteMediator struct {
	Bill
	Ted
}

func NewMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

func (m *ConcreteMediator) Communicate(who string) {
	if who == "Ted" {
		m.Bill.Respond()
	} else {
		m.Ted.Respond()
	}
}

func main() {
	mediator := NewMediator()
	mediator.Ted.Talk()
	// 	Ted:Bill
	// Bill: What
	// Ted: how much

}

// mediator
// 英 [ˈmiːdieɪtə(r)]   美 [ˈmiːdieɪtər]
// n.
// 调停者;斡旋者;解决纷争的人(或机构)

// communicate
// 英 [kəˈmjuːnɪkeɪt]   美 [kəˈmjuːnɪkeɪt]
// v.
// (与某人)交流(信息或消息、意见等);沟通;传达，传递(想法、感情、思想等);传染;相通

// wild
// 英 [waɪld]   美 [waɪld]
// adj.
// 自然生长的;野的;野生的;天然的;荒凉的;荒芜的;缺乏管教的;无法无天的;放荡的;感情炽烈的;盲目的;很棒的;热衷于…;狂暴的
// n.
// 自然环境;野生状态;偏远地区;人烟稀少的地区
// adv.
// 狂暴地;猛烈地;胡乱地

// stallion
// 英 [ˈstæliən]   美 [ˈstæliən]
// n.
// 牡马;(尤指)种马

// bill
// 英 [bɪl]   美 [bɪl]
// n.
// 账单;(餐馆的)账单;(提交议会讨论的)议案，法案;（剧院等的）节目单;海报;鸟嘴;有…形喙的
// vt.
// 给(某人)开账单，发账单(要求付款);把(某人或事物)宣传为…;宣布…将做某事

// respond
// 英 [rɪˈspɒnd]   美 [rɪˈspɑːnd]
// v.
// 回应;响应;作出反应;(口头或书面)回答;反应灵敏;作出正确反应;有改进
// vi.
// 响应;作出反应;反应灵敏;作出正确反应
// n.
// (支承拱或闭合连拱廊列柱的)附墙柱，壁联;应唱圣歌

// ted  特德
// 英 [tɛd]   美 [tɛd]
// n.
// 同 Teddy boy， 阿飞（指 20 世纪 50 年代穿潮流服饰，爱好摇滚乐的年轻男子）
// vt.
// 翻晒

// strange
// 英 [streɪndʒ]   美 [streɪndʒ]
// adj.
// 奇怪的;奇特的;异常的;陌生的;不熟悉的

// concrete
// 英 [ˈkɒŋkriːt]   美 [ˈkɑːŋkriːt]
// n.
// 混凝土
// adj.
// 混凝土制的;确实的，具体的(而非想象或猜测的);有形的;实在的
// vt.
// 用混凝土覆盖
