package main

import (
	"encoding/json"
	"fmt"
)

type Circle struct {
	radius float64
}

func (c Circle) radius1(rad float64) float64 {
	c.radius = 2
	rad = c.radius
	return rad
}
func (c *Circle) radius2(rad1 float64) float64 {
	// (*c).radius =2
	// rad = (*c).radius
	//也可以写成如下
	c.radius = 2
	rad1 = c.radius
	return rad1
}

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写一个方法改变i的值
func (i *integer) change() {
	*i = *i + 1
}

func main() {

	//go用结构体代替了类，去掉了传统的oop语言的继承，方法重载，构造函数和析构函数，隐藏的this指针等等
	//go仍然有面向对象编程的继承，封装和多态
	//和其他的oop不同的是，golang没有extends关键字，继承是通过匿字段来实现的

	//定义结构体
	type Cat struct {
		name string //即叫属性又叫field，还叫字段
		age  int
	}
	//定义变量
	var cat1 Cat
	cat1.name = "tom"
	cat1.age = 18
	fmt.Println("cat1 = ", cat1)

	//字段可以是基本数据类型，数组，也可以是引用类型
	//引用以后没有赋值，会有默认值，布尔类型是false，数值是0，字符串是""，
	//数组类型的默认值和它的元素类型相关，比如score[3]int 则为[0,0,0]
	//指针，slice和map的零值都是nil，即还没有分配空间

	type People struct {
		Name   string
		Age    int
		Scores [5]float64
		ptr    *int              //指针
		slice  []int             //切片
		map1   map[string]string //map
	}

	var people1 People

	//使用slice，一定要使用make
	people1.slice = make([]int, 2, 3)
	people1.slice[0] = 1
	//使用map,一定要make
	people1.map1 = make(map[string]string, 2)
	people1.map1["no1"] = "tom"

	//不同结构体变量的字段是独立，互不影响，一个结构变量字段的更改，不影响另外一个，结构体是值类型
	type Monster struct {
		Name string
		Age  int
	}
	var monster1 Monster
	monster1.Name = "hsir"
	monster1.Age = 18

	monster2 := monster1
	monster2.Name = "hsir1990"

	fmt.Println("monster1 = ", monster1) //monster1 = {}     monster1并没有改变
	fmt.Println("monster2 = ", monster2) //

	//声明方式一   直接声明
	//var person Person
	//声明方式二  {}
	var monster3 Monster = Monster{"tom", 18}
	fmt.Println("monster3 = ", monster3)

	//方式三  &
	var monster4 *Monster = new(Monster)
	//因为monster4是一个指针，因此标准的给字段赋值方式
	//(*monster4).Name = "smitch" 也可以这样写， monster4.Nam = "smitch" //系统为了简化，就默认可以这样了，对(*monster4).Name进行了处理  （go）
	(*monster4).Name = "smitch"
	monster4.Age = 18

	fmt.Println("*monster4=", *monster4)

	//方式4 -- {}
	var monster5 *Monster = &Monster{}
	(*monster5).Name = "smitch"
	monster5.Age = 18

	fmt.Println("*monster5=", *monster5)

	//第三种和第四种返回的是结构体指针

	var monster6 *Monster = &monster2
	(*monster6).Name = "tom1"
	fmt.Println("*monster2=", monster2.Name)    //变成了tom1
	fmt.Println("*monster6=", (*monster6).Name) //变成了tom1     //不能写成 *monster6.Name,因为 . 的运行符的优先级比 * 高

	//1.  结构体的所有字段在内存中是连续的
	//结构体
	type Point struct {
		x int
		y int
	}

	//结构体
	type Rect struct {
		leftUp, rightDown Point
	}

	//结构体
	type Rect2 struct {
		leftUp, rightDown *Point
	}

	r1 := Rect{Point{1, 2}, Point{3, 4}}

	//r1有四个int, 在内存中是连续分布
	//打印地址
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的，
	//但是他们指向的地址不一定是连续

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	//打印地址
	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p \n",
		&r2.leftUp, &r2.rightDown)

	//他们指向的地址不一定是连续...， 这个要看系统在运行时是如何分配
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p \n",
		r2.leftUp, r2.rightDown)

	//2.  结构体是用户单独定义的类型,和其他类型进行转换时需要有完全相同的字段（名字，个数和类型）
	type A struct {
		Num int
	}
	type B struct {
		Num int
	}
	var a A
	var b B
	a = A(b) // ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
	fmt.Println(a, b)

	//3.  结构体进行type重新定义（相当于取别名）.golang 认为是新的数据类型，但是相互间可以强转
	// type Student2 struct {
	// 	Name "string"
	// }
	// type Stu2 Student2
	// var stu1 Student2
	// var stu2 Stu2
	// stu2 = stu1  //报错可以写成 stu2 = Stu2(stu1)

	type interger int
	var i interger = 10
	var j int = 20
	j = int(i)
	fmt.Println(i, j)
	//4.  struct 的每个字段上，可以写写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
	type Monster2 struct {
		Name string `json:"name"` //`json:"name"`就是stuct tag
		Age  int    `json:"age"`
	}
	monster7 := Monster2{
		"牛魔王", 15,
	}
	//将monster序列化为json格式字串
	//json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster7)
	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))

	//方法中 func (p Person) test()... p表示哪个Persion变量调用，
	//这个p就是它的副本，这点和函数传参非常相似

	//方法的调用和传参机制原理
	//变量方法的调用和普通函数的区别在于，该变量本身也会作为一个参数传递到方法（如果变量是值类型，则进行值拷贝，如果变量是引用类型，则进行地址拷贝）

	//结构方法需要定义到main函数外
	// type Circle struct {
	// 	radius float64
	// }

	// func (c Circle) radius1(rad float64) float64 {
	// 	c.radius = 2
	// 	rad = c.radius
	// 	return rad
	// }
	// func (c *Circle) radius2(rad1 float64) float64 {
	// 	// (*c).radius =2
	// 	// rad = (*c).radius
	// 	//也可以写成如下
	// 	c.radius = 2
	// 	rad1 = c.radius
	// 	return rad1
	// }

	var c1 Circle
	c1.radius = 1
	c11 := c1.radius1(3)
	fmt.Println("c11=", c11, "c1.radius=", c1.radius)
	var c2 Circle
	c2.radius = 1
	c21 := c2.radius2(3)
	fmt.Println("c21=", c21, "c21.radius=", c2.radius)

	//3. go 中的方法作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，都可以有方法，而不仅仅是struct，比如int， float32等都可以有方法
	var i1 integer = 1
	i1.print()
	i1.change()
	fmt.Println("i1=", i1)

	//4 方法的大写，外部才能引用和函数一样
	//5 如果一个类型实现了String（）这个方法，那么fmt.Println默认会调用这个变量的 String（）进行输出
	stu1 := Student3{"tom", 10}
	fmt.Println(&stu1) //fmt.Println会默认调用这个变量

	//对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	p1 := Person{"tom"}
	test01(p1)
	test02(&p1)
	// 以上两个都对

	//对于struct的方法，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样可以

	//对于方法（如struct的方法），

	p1.test03()
	fmt.Println("main() p1.name=", p1.Name) // tom

	(&p1).test03() // 从形式上是传入地址，但是本质仍然是值拷贝

	fmt.Println("main() p1.name=", p1.Name) // tom

	(&p1).test04()
	fmt.Println("main() p1.name=", p1.Name) // mary
	p1.test04()                             // 等价 (&p).test04 , 从形式上是传入值类型，但是本质仍然是地址拷贝

	//总结
	//不管调用形式如何，真正决定时值拷贝还是地址拷贝，看这个方法是和哪个类型绑定
	//如果是和值类型，比如（p Person）,则是值拷贝，如果和指针类型，比如是（p *Person）则是地址拷贝

	//方式1
	//在创建结构体变量时，就直接指定字段的值
	var stud1 = Stu{"小明", 19} // stud1---> 结构体数据空间
	stud2 := Stu{"小明~", 20}

	//在创建结构体变量时，把字段名和字段值写在一起, 这种写法，就不依赖字段的定义顺序.
	var stud3 = Stu{
		Name: "jack",
		Age:  20,
	}
	stud4 := Stu{
		Age:  30,
		Name: "mary",
	}

	fmt.Println(stud1, stud2, stud3, stud4) //{小明 19} {小明~ 20} {jack 20} {mary 30}

	//方式2， 返回结构体的指针类型(!!!)
	var stud5 *Stu = &Stu{"小王", 29} // stud5--> 地址 ---》 结构体数据[xxxx,xxx]
	stud6 := &Stu{"小王~", 39}

	//在创建结构体指针变量时，把字段名和字段值写在一起, 这种写法，就不依赖字段的定义顺序.
	var stud7 = &Stu{
		Name: "小李",
		Age:  49,
	}
	stud8 := &Stu{
		Age:  59,
		Name: "小李~",
	}
	fmt.Println(*stud5, *stud6, *stud7, *stud8) //{小王 29} {小王~ 39} {小李 49} {小李~ 59}

}

type Stu struct {
	Name string
	Age  int
}

type Person struct {
	Name string
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

//接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03() =", p.Name) // jack
}

func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test03() =", p.Name) // mary
}

type Student3 struct {
	Name string
	Age  int
}

//会覆盖string  //不用引用也可以实现 String的重写
func (s *Student3) String() string {
	str := fmt.Sprintf("Name=[%v],Age=[%v]", s.Name, s.Age)
	return str
}
