//解释器模式，我们写代码的时候，有编译的问题，你是按什么编译的，什么语法，什么规则，有没有按go的语法来编写，编译器要知道的，要不会报错。有编译器的好处是写的不对，会直接拦截报错
//执行一个表达式，表达式中创建一个地址空间，然后循环表达式，最终把需要解释的都放到这个栈里，返回放到一个树里，之后把产量赋值给这个值，
//再从树里执行这个方法，执行
package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(variables map[string]Expression) int
}

type Integer struct {
	integer int
}

func (i *Integer) Interpret(variables map[string]Expression) int {
	return i.integer
}

type Plus struct {
	leftOperand  Expression
	rightOperand Expression
}

func (p *Plus) Interpret(variables map[string]Expression) int {
	return p.leftOperand.Interpret(variables) + p.rightOperand.Interpret(variables)
}

func (e *Evaluator) Interpret(context map[string]Expression) int {
	return e.syntaxTree.Interpret(context)
}

type Variable struct {
	name string
}

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (v Variable) Interpret(variables map[string]Expression) int {
	value, found := variables[v.name]
	if !found {
		return 0
	}
	return value.Interpret(variables)
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

type Evaluator struct {
	syntaxTree Expression
}

func NewEvaluator(expression string) *Evaluator {
	expressionStack := new(Stack)
	for _, token := range strings.Split(expression, " ") {
		switch token {
		case "+":
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left, right}
			expressionStack.Push(subExpression)
		default:
			expressionStack.Push(&Variable{token})
		}
	}
	synctaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree: synctaxTree}
}

func main() {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Integer{6}
	variables["x"] = &Integer{10}
	variables["z"] = &Integer{41}
	result := sentence.Interpret(variables)
	fmt.Println(result) //51
	// assert.Equal(t,51,result)

}

// interpreter
// 英 [ɪnˈtɜːprətə(r)]   美 [ɪnˈtɜːrprətər]
// n.
// 口译译员;口译工作者;演绎(音乐、戏剧中人物等)的人;解释程序
// 复数： interpreters

// expression
// 英 [ɪkˈspreʃn]   美 [ɪkˈspreʃn]
// n.
// 表示;表达;表露;表情;神色;词语;措辞;表达方式;感情，表情;式
// 复数： expressions
// 记忆技巧：express 表达 + ion 表名词 → 表达；表情

// interpret
// 英 [ɪnˈtɜːprət]   美 [ɪnˈtɜːrprət]
// v.
// 解释;诠释;口译;说明;把…理解为;领会;演绎
// 第三人称单数： interprets现在分词： interpreting过去式： interpreted过去分词： interpreted
// 派生词： interpretable adj.

// variables
// 英 [ˈveərɪəblz]   美 [ˈvɛriəbəlz]
// n.
// 可变情况;变量;可变因素
// variable的复数

// Operand
// 英 [ˈɒpərænd]   美 [ˈɑːpərænd]
// 操作数;运算元;运算对象;作数;操作符

// context
// 英 [ˈkɒntekst]   美 [ˈkɑːntekst]
// n.
// 上下文;语境;(事情发生的)背景，环境，来龙去脉
// 复数： contexts
// 记忆技巧：con 共同 + text 编织 →〔内容〕共同编织在一起的 → 上下文

// evaluator
// 英 [ɪˈvæljʊeɪtə]   美 [ɪˈvæljuˌeɪtər]
// n.
// 评审因子；评估员；鉴别器；求值程序

// syntax
// 英 [ˈsɪntæks]   美 [ˈsɪntæks]
// n.
// 句法;句法规则;语构

// stack
// 英 [stæk]   美 [stæk]
// n.
// 堆栈;(通常指码放整齐的)一叠，一摞，一堆;大量;许多;一大堆;(尤指工厂的)大烟囱;书库
// v.
// (使)放成整齐的一叠(或一摞、一堆);使成叠(或成摞、成堆)地放在…;使码放在…;(令飞机)分层盘旋等待着陆
// 第三人称单数： stacks复数： stacks现在分词： stacking过去式： stacked过去分词： stacked

// sentence
// 英 [ˈsentəns]   美 [ˈsentəns]
// n.
// 句子;判决;宣判;判刑
// vt.
// 判决;宣判;判刑
// 第三人称单数： sentences复数： sentences现在分词： sentencing过去式： sentenced过去分词： sentenced
