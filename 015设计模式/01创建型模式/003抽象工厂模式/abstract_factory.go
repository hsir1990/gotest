//抽象方法更加灵活，可以自己生产自己的东西，通过cook做自己的东西
//NewSimpleShapeFactory 生成 simpleLunchFactory 方法调用CreateFood方法，在调用rise的结构体来调用 Cook()

package main

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("It is a rise")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("It is a tomato")
}

type LunchFactory interface {
	CreatFood() Lunch
	CreatVegetable() Lunch
}

type simpleLunchFactory struct {
}

//用接口就不用使用指针的方式了，不用在前面加星号了
func NewSimpleShapeFactory() LunchFactory {
	return &simpleLunchFactory{}
}

//因为本身就是值类型，所以返回接口地址要用simpleLunchFactory结构的指针类型来接受
// func NewSimpleShapeFactory() *simpleLunchFactory {
// 	return &simpleLunchFactory{}
// }

func (s *simpleLunchFactory) CreatFood() Lunch {
	return &Rise{} //返回的需要是实例，而不是接口
}

func (s *simpleLunchFactory) CreatVegetable() Lunch {
	return &Tomato{}
}
func main() {

	factory := NewSimpleShapeFactory()
	food := factory.CreatFood()
	food.Cook()

	vegetable := factory.CreatVegetable()
	vegetable.Cook()
	// It is a rise
	// It is a tomato
}

// lunch
// 英 [lʌntʃ]   美 [lʌntʃ]
// n.
// 午餐;午饭
// vi.
// (尤指在餐馆)用午餐

// shape
// 英 [ʃeɪp]   美 [ʃeɪp]
// n.
// 形状;外形;样子;呈…形状的事物;模糊的影子;状况;情况;性质
// v.
// 使成为…形状(或样子);塑造;决定…的形成;影响…的发展;准备(做某动作);摆好姿势
