//两个村子有一条河，怎么过呢，我们搭一个桥
//我们画圆圈，有红的，绿的，黄的，怎么画呢？我们就要把圆和画的动作分开，分开以后又怎么把他们联系起来
//用简单的例子理解设计模式
//看最后是哪个注入的

//调用Draw，会调用Circle的实例，然后使用Shape实例，在调用draw接口类型，然后看哪个然后调用它
package main

import (
	"fmt"
	"time"
)

type Draw interface {
	DrawCircle(radius, x, y int)
}

type RedCircle struct {
}

func (r *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}

type YellowCircle struct {
}

func (c *YellowCircle) DrawCircle(radius, x, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}

type Shape struct {
	draw Draw
}

func (s *Shape) Shape(d Draw) {
	s.draw = d
	time.Now().Unix()
}

type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func (c *Circle) Constructor(x int, y int, radius int, draw Draw) {
	c.x = x
	c.y = y
	c.radius = radius
	c.shape.Shape(draw)
}

func (c *Circle) Cook() {
	c.shape.draw.DrawCircle(c.radius, c.x, c.y)
}

func main() {
	redCircle := Circle{}                             //先创建圆
	redCircle.Constructor(100, 100, 10, &RedCircle{}) //给圆的构造函数，加参数，通过shape加载注入draw

	yellowCircle := Circle{}
	yellowCircle.Constructor(200, 200, 10, &YellowCircle{})

	redCircle.Cook() //通过shape获取draw调用DrawCircle方法
	yellowCircle.Cook()

	// 	radius、x、y: 10 100 100
	// radius、x、y: 10 200 200
}

// draw
// 英 [drɔː]   美 [drɔː]
// v.
// 画;(用铅笔、钢笔或粉笔)描绘;描画;拖(动);拉(动);牵引;拖(车);吸引;（向某个方向）移动，行进;拔出;产生，引起，激起（反应或回应）;使说出;获取;进行，作（比较或对比）;抽（签、牌）;以平局结束;提取;抽出;抽（烟）
// n.
// 抽奖;平局;抽签;抽彩;和局;不分胜负;由抽签决定对手的比赛;有吸引力的人（或事物）;吸烟

// shape
// 英 [ʃeɪp]   美 [ʃeɪp]
// n.
// 形状;外形;样子;呈…形状的事物;模糊的影子;状况;情况;性质
// v.
// 使成为…形状(或样子);塑造;决定…的形成;影响…的发展;准备(做某动作);摆好姿势

// cook
// 英 [kʊk]   美 [kʊk]
// v.
// 烹调;烹饪;煮(或烘烤、煎炸等);密谋;秘密策划
// n.
// 厨师;做饭的人;炊事员
