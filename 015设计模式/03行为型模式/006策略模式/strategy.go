//策略模式就是当一件事情要发生时，我们要执行事先的预案
//我们在写代码时，有生成，测试，开发环境，我们把代码放到某个环境中，这就是一个预案
//通过调用，注入进a b 然后调用其方法
package main

import (
	"fmt"
)

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func (s *strategyA) Execute() {
	fmt.Println("A plan Executed")
}

func NewStrategyA() Strategy {
	return &strategyA{}
}

type strategyB struct {
}

func (s *strategyB) Execute() {
	fmt.Println("B plan Executed")
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}

func main() {
	strategyA := NewStrategyA()
	c := NewContext()
	c.SetStrategy(strategyA)
	c.Execute()

	strategyB := NewStrategyB()
	c.SetStrategy(strategyB)
	c.Execute()

	// 	A plan Executed
	// B plan Executed

}

// strategy
// 英 [ˈstrætədʒi]   美 [ˈstrætədʒi]
// n.
// 策略;计策;行动计划;策划;规划;部署;统筹安排;战略;战略部署
// 复数： strategies

// execute
// 英 [ˈeksɪkjuːt]   美 [ˈeksɪkjuːt]
// vt.
// (尤指依法)处决，处死;实行;执行;实施;成功地完成(技巧或动作);制作，做成（艺术品）;执行（法令）
// 第三人称单数： executes现在分词： executing过去式： executed过去分词： executed

// context
// 英 [ˈkɒntekst]   美 [ˈkɑːntekst]
// n.
// 上下文;语境;(事情发生的)背景，环境，来龙去脉
// 复数： contexts
// 记忆技巧：con 共同 + text 编织 →〔内容〕共同编织在一起的 → 上下文
