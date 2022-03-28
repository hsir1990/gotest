//适配器模式，比如咱们去国外，国外是110伏的，咱们用不了，需要找一个转换头，也就是适配器
//通过继承方法，去实现适配
package main

import "fmt"

type Target interface {
	Execute()
}

type Adaptee struct {
}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("充电...")
}

type Adapter struct {
	*Adaptee //匿名结构体//继承了之前的结构体的属性和方法
}

func (a *Adapter) Execute() {
	a.SpecificExecute()
}

func main() {
	adapter := Adapter{}
	adapter.Execute()
	// 充电...

}

// adaptor
// 英 [əˈdæptə(r)]   美 [əˈdæptər]
// n.
// (电器设备的)转接器，适配器;(供多个设备连接电源的)多头插头，多功能插头

// specific
// 英 [spəˈsɪfɪk]   美 [spəˈsɪfɪk]
// adj.
// 具体的;特定的;特有的;明确的;独特的
// n.
// 特效药;特性;细节;显著的性质;特性
// 复数： specifics比较级： more specific最高级： most specific
// 记忆技巧：speci 种类 + fic 具有某种性质的 → 分出类别 → 明确的
