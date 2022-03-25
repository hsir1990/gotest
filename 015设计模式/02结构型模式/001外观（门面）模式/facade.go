//门面模式也就是外观模式，解决的问题是把你所有的内部封装都封装到门面里面，就像门面出外卖一样，
//汽车不会把发动机，油箱之类的暴露给你
//调用门面结构体，门面结构体中包含发动机和外壳等，然后通过结构体去调用里面的部件，在调用里面的方法
package main

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Println("CarModel - SetModel")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Println("CarEngine - SetEngine")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetBody() {
	fmt.Println("CarBody - SetBody")
}

type CarAccessories struct {
}

func NewCarAccessories() *CarAccessories {
	return &CarAccessories{} //这个地需要加地址符号，因为这个地方不属于结构体
}

func (c *CarAccessories) SetAccessories() {
	fmt.Println("CarAccessories - SetAccessories")
}

//汽车门面
type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		accessories: NewCarAccessories(),
		body:        NewCarBody(),
		engine:      NewCarEngine(),
		model:       NewCarModel(),
	}
}

func (c *CarFacade) CreateCompleteCar() {
	fmt.Println("**********  Creating a Car  ***********")
	c.model.SetModel()
	c.body.SetBody()
	c.engine.SetEngine()
	c.accessories.SetAccessories()
	fmt.Println("**********  Car creation is completed.  ***********")
}

func main() {
	facade := NewCarFacade()
	facade.CreateCompleteCar()

	// **********  Creating a Car  ***********
	// CarModel - SetModel
	// CarBody - SetBody
	// CarEngine - SetEngine
	// CarAccessories - SetAccessories
	// **********  Car creation is completed.  ***********

}

// facade
// 英 [fəˈsɑːd]   美 [fəˈsɑːd]
// n.
// 外观;(建筑物的)正面，立面;(虚假的)表面，外表

// accessories
// 英 [əkˈsɛsəriz]   美 [ækˈsɛsəriz]
// n.
// 附件;配件;附属物;(衣服的)配饰;从犯;同谋;帮凶
// accessory的复数

// complete
// 英 [kəmˈpliːt]   美 [kəmˈpliːt]
// vt.
// 完成;结束;填写(表格);使完整;使完美
// adj.
// 完整的;(用以强调)完全的，彻底的;全部的;整个的;包括，含有(额外部分或特征);完成
