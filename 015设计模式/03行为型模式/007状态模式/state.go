//状态模式，比如开关，有开有关，比如应用程序上，限流，设置限流的时候有一个开关
//新建一个机器，机器里面有状态结构，newOn和newOff实现了这个接口，然后在机器内部来回轮转切换
package main

import (
	"fmt"
)

type State interface {
	On(m *Machine)
	Off(m *Machine)
}

type Machine struct {
	current State
}

func NewMachine() *Machine {
	return &Machine{NewOFF()}
}

func (m *Machine) setCurrent(s State) {
	m.current = s
}

func (m *Machine) On() {
	m.current.On(m)
}

func (m *Machine) Off() {
	m.current.Off(m)
}

type ON struct {
}

func NewON() State {
	return &ON{}
}

func (o *ON) On(m *Machine) {
	fmt.Println("已经开启")
}

func (o *ON) Off(m *Machine) {
	fmt.Println("从On的状态到Off")
	m.setCurrent(NewOFF())
}

type OFF struct {
}

func NewOFF() State {
	return &OFF{}
}

func (o *OFF) On(m *Machine) {
	fmt.Println("从OFF的状态到ON")
	m.setCurrent(NewON())
}

func (o *OFF) Off(m *Machine) {
	fmt.Println("已经关闭")
}

func main() {
	machine := NewMachine()
	machine.Off()
	machine.On()
	machine.On()
	machine.Off()
	// 	已经关闭
	// 从OFF的状态到ON
	// 已经开启
	// 从On的状态到Off

}

// machine
// 英 [məˈʃiːn]   美 [məˈʃiːn]
// n.
// 机器;机械装置;(不提全称时的简略说法)机器;(组织的)核心机构;机械化的人
// v.
// (用机器)制造，加工成型
// 第三人称单数： machines复数： machines现在分词： machining过去式： machined过去分词： machined
// 记忆技巧：mach 机器 + ine 抽象名词 → 机器
