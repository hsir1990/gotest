桥模式
桥接模式分离抽象部分和实现部分。使得两部分独立扩展。

桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。

策略模式使抽象部分和实现部分分离，可以独立变化。


// abstract
// 英 [ˈæbstrækt , æbˈstrækt]   美 [ˈæbstrækt , æbˈstrækt]  
// n.
// 摘要;抽象派艺术作品;(文献等的)概要
// adj.
// 抽象的(与个别情况相对);纯理论的;抽象的(与具体经验相对);抽象(派)的
// vt.
// 把…抽象出;提取;抽取;分离;写摘要


bridge.go
package bridge

import "fmt"

type AbstractMessage interface {
    SendMessage(text, to string)
}

type MessageImplementer interface {
    Send(text, to string)
}

type MessageSMS struct{}

func ViaSMS() MessageImplementer {
    return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
    fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
    return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
    fmt.Printf("send %s to %s via Email", text, to)
}

type CommonMessage struct {
    method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
    return &CommonMessage{
        method: method,
    }
}

func (m *CommonMessage) SendMessage(text, to string) {
    m.method.Send(text, to)
}

type UrgencyMessage struct {
    method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
    return &UrgencyMessage{
        method: method,
    }
}

func (m *UrgencyMessage) SendMessage(text, to string) {
    m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
bridge_test.go
package bridge

func ExampleCommonSMS() {
    m := NewCommonMessage(ViaSMS())
    m.SendMessage("have a drink?", "bob")
    // Output:
    // send have a drink? to bob via SMS
}

func ExampleCommonEmail() {
    m := NewCommonMessage(ViaEmail())
    m.SendMessage("have a drink?", "bob")
    // Output:
    // send have a drink? to bob via Email
}

func ExampleUrgencySMS() {
    m := NewUrgencyMessage(ViaSMS())
    m.SendMessage("have a drink?", "bob")
    // Output:
    // send [Urgency] have a drink? to bob via SMS
}

func ExampleUrgencyEmail() {
    m := NewUrgencyMessage(ViaEmail())
    m.SendMessage("have a drink?", "bob")
    // Output:
    // send [Urgency] have a drink? to bob via Email
}
