//组合模式就是把许多单个的东西放到一起，然后来完成一个任务
//定义一个组合部件的切片，然后利用添加方法，上面写的时接口行为，而我们要添加的是结构的类型的对象，也就是叶子对象，加进去以后，我们又把所有的切片组合放到了第一个切片组里面，
//然后遍历所有的组合，调用组合的traverse方法，这个方法也会遍历所有的叶子的中的traverse方法
package main

import "fmt"

type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func NewLeaf(value int) *Leaf {
	return &Leaf{value}
}

func (l *Leaf) Traverse() {
	fmt.Println(l.value)
}

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Traverse() {
	for idx, _ := range c.children {
		c.children[idx].Traverse()
	}
}

func main() {
	containers := make([]Composite, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			containers[i].Add(NewLeaf(i*3 + j))
		}
	}

	for i := 1; i < 4; i++ {
		containers[0].Add(&(containers[i]))
	}

	for i := 0; i < 4; i++ {
		containers[i].Traverse()
		fmt.Printf("Finished \n")
	}

	// 	0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
	// Finished
	// 3
	// 4
	// 5
	// Finished
	// 6
	// 7
	// 8
	// Finished
	// 9
	// 10
	// 11
	// Finished
	fmt.Println(containers)
	// [{[0xc00000a0b8 0xc00000a0d0 0xc00000a0d8 0xc00002a0d8 0xc00002a0f0 0xc00002a108]} {[0xc00000a0e0 0xc00000a0e8 0xc00000a0f0]} {[0xc00000a0f8 0xc00000a100 0xc00000a108]} {[0xc00000a110 0xc00000a118 0xc00000a120]}]

}

// composite
// 英 [ˈkɒmpəzɪt]   美 [kəmˈpɑːzət]
// adj.
// 混合成的;复合的;合成的;混成的
// n.
// 复合材料;合成物;混合物
// 复数： composites
// 记忆技巧：com 共同 + posit 放，放置 + e → 放到一起 → 合成物

// traverse
// 英 [trəˈvɜːs , ˈtrævɜːs]   美 [trəˈvɜːrs , ˈtrævɜːrs]
// vt.
// 横过;横越;穿过;横渡
// n.
// (在陡坡上的)侧向移动，横过，横越;可横越的地方
// v.
// 穿过;横过;横越;横渡
// adj.
// 横断的
// 第三人称单数： traverses现在分词： traversing过去式： traversed过去分词： traversed

// struct
// 结构;结构体;结构体类型

// leaf
// 英 [liːf]   美 [liːf]
// n.
// 叶;叶片;叶子;有…状叶的;有…片叶的;(纸)页，张;(尤指书的)页;薄金属片;活动桌板
// v.
// 翻…的页，匆匆翻阅
// 第三人称单数： leaves复数： leaves现在分词： leafing过去式： leafed过去分词： leafed

// component
// 英 [kəmˈpəʊnənt]   美 [kəmˈpoʊnənt]
// n.
// 组成部分;成分;部件
// adj.
// 成分的;组成的;合成的;构成的
// 复数： components
// 派生词： component adj.
// 记忆技巧：com 共同 + pon 放，放置 + ent 关于…的 → 共同放到一起 → 组成的

// containers
// 英 [kənˈteɪnəz]   美 [kənˈteɪnərz]
// n.
// 容器;集装箱;货柜
// container的复数

// container
// 英 [kənˈteɪnə(r)]   美 [kənˈteɪnər]
// n.
// 容器;集装箱;货柜
// 复数： containers
// 记忆技巧：contain 包含，容纳 + er 表物 → 容器
