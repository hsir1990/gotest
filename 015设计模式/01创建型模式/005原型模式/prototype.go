//原型模式，通过clone一个自己出来，通过指针指向，最后还是给自己，等于结构套结构，自己指向自己
package main

import (
	"fmt"
)

type Prototype interface {
	Name() string
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
}

func (c *ConcretePrototype) Name() string {
	return c.name
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: c.name}
}

func main() {
	name := "出去浪"
	proto := ConcretePrototype{name: name}
	newProto := proto.Clone()
	actualName := newProto.Name()
	// assert.Equal(t, name, actualName) // assert库，预期的名字和实际的名字作比较，是否相等，相等就通过了
	fmt.Println(name)
	fmt.Println(actualName)
	// 出去浪
	// 出去浪

}

// prototype
// 英 [ˈprəʊtətaɪp]   美 [ˈproʊtətaɪp]
// n.
// 原型;雏形;最初形态

// actual
// 英 [ˈæktʃuəl]   美 [ˈæktʃuəl]
// adj.
// 真实的;实际的;(强调事情最重要的部分)真正的，…本身
// 记忆技巧：act 动作，行动 + ual 有…性质的，属于…的 → 行动的 → 实际的
