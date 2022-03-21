装饰模式
装饰模式使用对象组合的方式动态改变或增加对象行为。

Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。

使用匿名组合，在装饰器中不必显式定义转调原对象方法。


// concrete
// 英 [ˈkɒŋkriːt]   美 [ˈkɑːŋkriːt]  
// n.
// 混凝土
// adj.
// 混凝土制的;确实的，具体的(而非想象或猜测的);有形的;实在的
// vt.
// 用混凝土覆盖


// decorator
// 英 [ˈdekəreɪtə(r)]   美 [ˈdekəreɪtər]  
// 修饰器;修饰模式;装饰模式;装饰;装饰器


// component
// 英 [kəmˈpəʊnənt]   美 [kəmˈpoʊnənt]  
// n.
// 组成部分;成分;部件
// adj.
// 成分的;组成的;合成的;构成的

decorator.go
package decorator

type Component interface {
    Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
    return 0
}

type MulDecorator struct {
    Component
    num int
}

func WarpMulDecorator(c Component, num int) Component {
    return &MulDecorator{
        Component: c,
        num:       num,
    }
}

func (d *MulDecorator) Calc() int {
    return d.Component.Calc() * d.num
}

type AddDecorator struct {
    Component
    num int
}

func WarpAddDecorator(c Component, num int) Component {
    return &AddDecorator{
        Component: c,
        num:       num,
    }
}

func (d *AddDecorator) Calc() int {
    return d.Component.Calc() + d.num
}
decorator_test.go
package decorator

import "fmt"

func ExampleDecorator() {
    var c Component = &ConcreteComponent{}
    c = WarpAddDecorator(c, 10)
    c = WarpMulDecorator(c, 8)
    res := c.Calc()

    fmt.Printf("res %d\n", res)
    // Output:
    // res 80
}