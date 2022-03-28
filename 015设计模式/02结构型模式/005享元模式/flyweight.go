//享元就是共享单元的意思，意图就是对象的复用，节省内存
//创建一个享元工厂，工厂获取对象，如果池子里有就返回，没有就创建一个
package main

import "fmt"

type FlyWeight struct {
	Name string
}

func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name: name}
}

type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}

func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}

func (f *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	weight, ok := f.pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		f.pool[name] = weight
	}
	return weight
}

func main() {
	factory := NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong beauty")
	xiang := factory.GetFlyWeight("xiao beauty")

	fmt.Println(hong)
	fmt.Println(xiang)
	// &{hong beauty}
	// &{xiao beauty}

}

// flyweight
// 英 [ˈflaɪweɪt]   美 [ˈflaɪweɪt]
// n.
// 特轻量级拳击手，次最轻量级拳击手，最轻量级摔跤手(体重48至51公斤之间)
