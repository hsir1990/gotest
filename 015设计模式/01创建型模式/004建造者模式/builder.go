//建造者模式没听太懂  保持行为分离
//concretebuilder实现了build所有可以注入进去，也就是传递给他，其实注入和传递给他是一样的意思
//把接口和具体的现实分开了
//构造了最好的哪个build
//从行为到具体的实现是分离的

//先建了一个ConcreteBuilder的对象，然后把对象注入给director,然后调用Construct方法，调用 d.builder.Builder()方法等于是调用的接口的Builder方法
// 这样就把接口和具体的实现方法分开了，之后build就等于true了，然后获取结果
package main

import (
	"fmt"
)

type Builder interface {
	Builder()
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) Director {
	return Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.Builder()
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{built: false}
	//return ConcreteBuilder{false}
}

func (c *ConcreteBuilder) Builder() {
	c.built = true
}

type Product struct {
	Built bool
}

func (c *ConcreteBuilder) GetResult() Product {
	return Product{c.built}
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(&builder) //注入的是接口，所有要改成地址传入，不能用结构体
	director.Construct()

	product := builder.GetResult()
	fmt.Println(product.Built)
	//true
}

// concrete
// 英 [ˈkɒŋkriːt]   美 [ˈkɑːŋkriːt]
// n.
// 混凝土
// adj.
// 混凝土制的;确实的，具体的(而非想象或猜测的);有形的;实在的
// vt.
// 用混凝土覆盖
