//代理模式，比如租房的时候，找代理，他能帮着找到物美价廉的房子
//agent找到房子了，执行函数,房子地址北京，价格8000
//自己找代理，然后给他分配任务
//他在帮你找房子
//适合的场景再使用设计模式，不建议生搬硬套
package main

import (
	"fmt"
	"strconv"
)

type ITask interface {
	RentHouse(desc string, price int)
}

type Task struct {
}

func (t *Task) RentHouse(desc string, price int) {
	fmt.Sprintln(fmt.Printf("租房子的地址%s,价钱%s,中介费%s.", desc, strconv.Itoa(price), strconv.Itoa(price)))
}

type AgentTask struct {
	task *Task
}

func NewAgentTask() *AgentTask {
	return &AgentTask{task: &Task{}}
}

func (a *AgentTask) RentHouse(desc string, price int) {
	a.task.RentHouse(desc, price)
}

func main() {
	agent := NewAgentTask()
	agent.RentHouse("北京", 8000)
	//租房子的地址北京,价钱8000,中介费8000.

}

// proxy
// 英 [ˈprɒksi]   美 [ˈprɑːksi]
// n.
// 代理;代理人;代表;代理权;代表权;受托人;(测算用的)代替物;指标
// 复数： proxies

// task
// 英 [tɑːsk]   美 [tæsk]
// n.
// 任务;(尤指艰巨或令人厌烦的)工作;(尤指语言教学中旨在帮助达到某一学习目的的)活动
// vt.
// 交给某人(任务);派给某人(工作)
// 第三人称单数： tasks复数： tasks现在分词： tasking过去式： tasked过去分词： tasked

// rent
// 英 [rent]   美 [rent]
// n.
// 租金;破裂处;裂口;撕裂
// v.
// 租用，租借(房屋、土地、机器等);出租;将…租给;(短期)租用，租借;以…出租
// rend的过去分词和过去式
// 第三人称单数： rents复数： rents现在分词： renting过去式： rented过去分词： rented
