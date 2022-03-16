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

	//3.  结构体进行type重新定义  类型定义  （相当于半个取别名）.golang 认为是新的数据类型，但是相互间可以强转
	// type Student2 struct {
	// 	Name "string"
	// }
	// type Stu2 Student2
	// var stu1 Student2
	// var stu2 Stu2
	// stu2 = stu1  //报错可以写成 stu2 = Stu2(stu1)

	//类型别名
	// type MyInt = int
	// var b MyInt
	// fmt.Printf("type of b:%T\n", b) //type of b:int
	//类型定义
	// type NewInt int
	// var a NewInt
	// fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt

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

//临时数据，匿名结构体
// var user struct{Name string; Age int}
//     user.Name = "pprof.cn"
//     user.Age = 18
//     fmt.Printf("%#v\n", user)

// //创建指针类型结构体
// var p2 = new(person)
// fmt.Printf("%T\n", p2)     //*main.person
// p2.name = "测试"
// p2.age = 18
// p2.city = "北京"
// fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}

// 取结构体的地址实例化
// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。

// p3 := &person{}
// fmt.Printf("%T\n", p3)     //*main.person
// fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
// p3.name = "博客"
// p3.age = 30
// p3.city = "成都"
// fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

// p3.name = "博客"其实在底层是(*p3).name = "博客"，这是Go语言帮我们实现的语法糖。

// 使用键值对初始化
// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

// p5 := person{
//     name: "pprof.cn",
//     city: "北京",
//     age:  18,
// }
// fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}
// 也可以对结构体指针进行键值对初始化，例如：

// p6 := &person{
//     name: "pprof.cn",
//     city: "北京",
//     age:  18,
// }
// fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"pprof.cn", city:"北京", age:18}

// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

// p7 := &person{
//     city: "北京",
// }
// fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

// 结构体内存布局
// type test struct {
//     a int8
//     b int8
//     c int8
//     d int8
// }
// n := test{
//     1, 2, 3, 4,
// }
// fmt.Printf("n.a %p\n", &n.a)
// fmt.Printf("n.b %p\n", &n.b)
// fmt.Printf("n.c %p\n", &n.c)
// fmt.Printf("n.d %p\n", &n.d)
// 输出：

//     n.a 0xc0000a0060
//     n.b 0xc0000a0061
//     n.c 0xc0000a0062
//     n.d 0xc0000a0063

// 面试题1
// type student struct {
//     name string
//     age  int
// }

// func main() {
//     m := make(map[string]*student)
//     stus := []student{
//         {name: "pprof.cn", age: 18},
//         {name: "测试", age: 23},
//         {name: "博客", age: 28},
//     }

//     for _, stu := range stus {
//         m[stu.name] = &stu//循环过程中，stu变量只声明了一次，所以stu地址即&stu是不变的，值是变化的。所以&stu始终不变
//     }
//     for k, v := range m {
//         fmt.Println(k, "=>", v.name)
//     }
// }
// pprof.cn => 博客
// 测试 => 博客
// 博客 => 博客
// for range 每次产生的 key 和 value 其实是对应的 stus 里面值的拷贝，不是对应的 stus 里面的值的引用，所以出现了这种问题。
// stu 是 stus 在for循环中申请的一个局部变量，每次循环都会拷贝 stus 中对应的值 stu。迭代遍历之后，stu 每次会被重新赋值，而在 m 这个 map 中记录的 value 只不过是 stu 的内存地址。
//可能是因为每次定义数据，用的是同一个地址，然后地址相同

// 重新申请一个变量，即可解决
//     for _, stu := range stus {
//         s:=stu
//         m[stu.name] = &s
//     }

// 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self。
//.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// 结构体的匿名字段
// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

// //Person 结构体Person类型
// type Person struct {
//     string
//     int
// }

// func main() {
//     p1 := Person{
//         "pprof.cn",
//         18,
//     }
//     fmt.Printf("%#v\n", p1)        //main.Person{string:"pprof.cn", int:18}
//     fmt.Println(p1.string, p1.int) //pprof.cn 18
// }
// 匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

// 嵌套结构体
// 一个结构体中可以嵌套包含另一个结构体或结构体指针。

// //Address 地址结构体
// type Address struct {
//     Province string
//     City     string
// }

// //User 用户结构体
// type User struct {
//     Name    string
//     Gender  string
//     Address Address
// }

// func main() {
//     user1 := User{
//         Name:   "pprof",
//         Gender: "女",
//         Address: Address{
//             Province: "黑龙江",
//             City:     "哈尔滨",
//         },
//     }
//     fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
// }
// 1.3.16. 嵌套匿名结构体
// //Address 地址结构体
// type Address struct {
//     Province string
//     City     string
// }

// //User 用户结构体
// type User struct {
//     Name    string
//     Gender  string
//     Address //匿名结构体
// }

// func main() {
//     var user2 User
//     user2.Name = "pprof"
//     user2.Gender = "女"
//     user2.Address.Province = "黑龙江"    //通过匿名结构体.字段名访问
//     user2.City = "哈尔滨"                //直接访问匿名结构体的字段名
//     fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
// }
// 当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

// 结构体的“继承”
// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

// //Animal 动物
// type Animal struct {
//     name string
// }

// func (a *Animal) move() {
//     fmt.Printf("%s会动！\n", a.name)
// }

// //Dog 狗
// type Dog struct {
//     Feet    int8
//     *Animal //通过嵌套匿名结构体实现继承
// }

// func (d *Dog) wang() {
//     fmt.Printf("%s会汪汪汪~\n", d.name)
// }

// func main() {
//     d1 := &Dog{
//         Feet: 4,
//         Animal: &Animal{ //注意嵌套的是结构体指针
//             name: "乐乐",
//         },
//     }
//     d1.wang() //乐乐会汪汪汪~
//     d1.move() //乐乐会动！
// }

// func (recevier type) methodName(参数列表)(返回值列表){}
// 当接受者不是一个指针时，该方法操作对应接受者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。

//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用

// 	普通函数与方法的区别
// 1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。

// 2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。

// package main

// //普通函数与方法的区别（在接收者分别为值类型和指针类型的时候）

// import (
//     "fmt"
// )

// //1.普通函数
// //接收值类型参数的函数
// func valueIntTest(a int) int {
//     return a + 10
// }

// //接收指针类型参数的函数
// func pointerIntTest(a *int) int {
//     return *a + 10
// }

// func structTestValue() {
//     a := 2
//     fmt.Println("valueIntTest:", valueIntTest(a))
//     //函数的参数为值类型，则不能直接将指针作为参数传递
//     //fmt.Println("valueIntTest:", valueIntTest(&a))
//     //compile error: cannot use &a (type *int) as type int in function argument

//     b := 5
//     fmt.Println("pointerIntTest:", pointerIntTest(&b))
//     //同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
//     //fmt.Println("pointerIntTest:", pointerIntTest(b))
//     //compile error:cannot use b (type int) as type *int in function argument
// }

// //2.方法
// type PersonD struct {
//     id   int
//     name string
// }

// //接收者为值类型
// func (p PersonD) valueShowName() {
//     fmt.Println(p.name)
// }

// //接收者为指针类型
// func (p *PersonD) pointShowName() {
//     fmt.Println(p.name)
// }

// func structTestFunc() {
//     //值类型调用方法
//     personValue := PersonD{101, "hello world"}
//     personValue.valueShowName()
//     personValue.pointShowName()

//     //指针类型调用方法
//     personPointer := &PersonD{102, "hello golang"}
//     personPointer.valueShowName()
//     personPointer.pointShowName()

//     //与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
// }

// func main() {
//     structTestValue()
//     structTestFunc()
// }
// 输出结果：

//     valueIntTest: 12
//     pointerIntTest: 15
//     hello world
//     hello world
//     hello golang
//     hello golang

// 匿名字段
// Golang匿名字段 ：可以像字段成员那样访问匿名字段方法，编译器负责查找。

// package main

// import "fmt"

// type User struct {
//     id   int
//     name string
// }

// type Manager struct {
//     User
// }

// func (self *User) ToString() string { // receiver = &(Manager.User)
//     return fmt.Sprintf("User: %p, %v", self, self)
// }

// func main() {
//     m := Manager{User{1, "Tom"}}
//     fmt.Printf("Manager: %p\n", &m)
//     fmt.Println(m.ToString())
// }
// 输出结果:

//     Manager: 0xc42000a060
//     User: 0xc42000a060, &{1 Tom}
// 通过匿名字段，可获得和继承类似的复用能力。依据编译器查找次序，只需在外层定义同名方法，就可以实现 "override"。

// package main

// import "fmt"

// type User struct {
//     id   int
//     name string
// }

// type Manager struct {
//     User
//     title string
// }

// func (self *User) ToString() string {
//     return fmt.Sprintf("User: %p, %v", self, self)
// }

// func (self *Manager) ToString() string {
//     return fmt.Sprintf("Manager: %p, %v", self, self)
// }

// func main() {
//     m := Manager{User{1, "Tom"}, "Administrator"}

//     fmt.Println(m.ToString())

//     fmt.Println(m.User.ToString())
// }
// 输出结果:

//     Manager: 0xc420074180, &{{1 Tom} Administrator}
//     User: 0xc420074180, &{1 Tom}

// ///////下面的需要测试一下 ，下面的理论不对
// // • 类型 T 方法集包含全部 receiver T 方法。 ----不对
// // • 类型 *T 方法集包含全部 receiver T + *T 方法。
// // • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。----不对
// // • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
// // • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
// type Stu struct {
// 	name string
// 	age  int
// }

// func (s Stu) test() {
// 	fmt.Println("test--- s", s.name)
// }
// func (s *Stu) test1() {
// 	fmt.Println("test1--- s", (*s).name)
// }

// var st Stu = Stu{
// 	name: "tom",
// 	age:  18,
// }
// st.test1() //这样没有问题，可以执行

// 这条规则说的是当我们嵌入一个类型的指针，嵌入类型的接受者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。

// package main

// import (
//     "fmt"
// )

// type S struct {
//     T
// }

// type T struct {
//     int
// }

// func (t T) testT() {
//     fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法")
// }
// func (t *T) testP() {
//     fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
// }

// func main() {
//     s1 := S{T{1}}
//     s2 := &s1
//     fmt.Printf("s1 is : %v\n", s1)
//     s1.testT()
//     s1.testP()
//     fmt.Printf("s2 is : %v\n", s2)
//     s2.testT()
//     s2.testP()
// }
// 输出结果：

//     s1 is : {{1}}
//     如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法
//     如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
//     s2 is : &{{1}}
//     如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法
//     如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法

// 表达式
// Golang 表达式 ：根据调用者不同，方法分为两种表现形式:

//     instance.method(args...) ---> <type>.func(instance, args...)
// 前者称为 method value，后者 method expression。

// 两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。

// package main

// import "fmt"

// type User struct {
//     id   int
//     name string
// }

// func (self *User) Test() {
//     fmt.Printf("%p, %v\n", self, self)
// }

// func main() {
//     u := User{1, "Tom"}
//     u.Test()

//     mValue := u.Test
//     mValue() // 隐式传递 receiver

//     mExpression := (*User).Test
//     mExpression(&u) // 显式传递 receiver
// }
// 输出结果:

//     0xc42000a060, &{1 Tom}
//     0xc42000a060, &{1 Tom}
//     0xc42000a060, &{1 Tom}
// 需要注意，method value 会复制 receiver。

// package main

// import "fmt"

// type User struct {
//     id   int
//     name string
// }

// func (self User) Test() {
//     fmt.Println(self)
// }

// func main() {
//     u := User{1, "Tom"}
//     mValue := u.Test // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

//     u.id, u.name = 2, "Jack"
//     u.Test()

//     mValue()
// }
// 输出结果

//     {2 Jack}
//     {1 Tom}
// 在汇编层面，method value 和闭包的实现方式相同，实际返回 FuncVal 类型对象。

//     FuncVal { method_address, receiver_copy }
// 可依据方法集转换 method expression，注意 receiver 类型的差异。

// package main

// import "fmt"

// type User struct {
//     id   int
//     name string
// }

// func (self *User) TestPointer() {
//     fmt.Printf("TestPointer: %p, %v\n", self, self)
// }

// func (self User) TestValue() {
//     fmt.Printf("TestValue: %p, %v\n", &self, self)
// }

// func main() {
//     u := User{1, "Tom"}
//     fmt.Printf("User: %p, %v\n", &u, u)

//     mv := User.TestValue
//     mv(u)

//     mp := (*User).TestPointer
//     mp(&u)

//     mp2 := (*User).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
//     mp2(&u)
// }
// 输出:

//     User: 0xc42000a060, {1 Tom}
//     TestValue: 0xc42000a0a0, {1 Tom}
//     TestPointer: 0xc42000a060, &{1 Tom}
//     TestValue: 0xc42000a100, {1 Tom}
// 将方法 "还原" 成函数，就容易理解下面的代码了。--------------我还是不理解这

// package main

// type Data struct{}

// func (Data) TestValue() {}

// func (*Data) TestPointer() {}

// func main() {
//     var p *Data = nil
//     p.TestPointer()

//     (*Data)(nil).TestPointer() // method value
//     (*Data).TestPointer(nil)   // method expression

//     // p.TestValue()            // invalid memory address or nil pointer dereference

//     // (Data)(nil).TestValue()  // cannot convert nil to type Data
//     // Data.TestValue(nil)      // cannot use nil as type Data in function argument
// }

// 自定义error：
// package main

// import (
//     "fmt"
//     "os"
//     "time"
// )

// type PathError struct {
//     path       string
//     op         string
//     createTime string
//     message    string
// }

// func (p *PathError) Error() string {
//     return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
//         p.op, p.createTime, p.message)
// }

// func Open(filename string) error {

//     file, err := os.Open(filename)
//     if err != nil {
//         return &PathError{
//             path:       filename, //get path error, path=/Users/pprof/Desktop/go/src/test.txt
//             op:         "read",
//             message:    err.Error(), //message=open /Users/pprof/Desktop/go/src/test.txt: no such file or directory
//             createTime: fmt.Sprintf("%v", time.Now()),  // createTime=2018-04-05 11:25:17.331915 +0800 CST m=+0.000441790
//         }
//     }

//     defer file.Close()
//     return nil
// }

// func main() {
//     err := Open("/Users/5lmh/Desktop/go/src/test.txt")
//     switch v := err.(type) {
//     case *PathError:
//         fmt.Println("get path error,", v)
//     default:

//     }

// }
// 输出结果：

//     get path error, path=/Users/pprof/Desktop/go/src/test.txt
//     op=read
//     createTime=2018-04-05 11:25:17.331915 +0800 CST m=+0.000441790
//     message=open /Users/pprof/Desktop/go/src/test.txt: no such file or directory

//go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

// 接口
// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

// package main

// import "fmt"

// //    go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

// //人
// type Person struct {
//     name string
//     sex  string
//     age  int
// }

// type Student struct {
//     Person
//     id   int
//     addr string
// }

// func main() {
//     // 初始化
//     s1 := Student{Person{"5lmh", "man", 20}, 1, "bj"}
//     fmt.Println(s1)

//     s2 := Student{Person: Person{"5lmh", "man", 20}}
//     fmt.Println(s2)

//     s3 := Student{Person: Person{name: "5lmh"}}
//     fmt.Println(s3)
// }
// 输出结果：

//     {{5lmh man 20} 1 bj}
//     {{5lmh man 20} 0 }
//     {{5lmh  0} 0 }
// 同名字段的情况

// package main

// import "fmt"

// //人
// type Person struct {
//     name string
//     sex  string
//     age  int
// }

// type Student struct {
//     Person
//     id   int
//     addr string
//     //同名字段
//     name string
// }

// func main() {
//     var s Student
//     // 给自己字段赋值了
//     s.name = "5lmh"
//     fmt.Println(s)

//     // 若给父类同名字段赋值，如下
//     s.Person.name = "枯藤"
//     fmt.Println(s)
// }
// 输出结果：

//     {{  0} 0  5lmh}
//     {{枯藤  0} 0  5lmh}
// 所有的内置类型和自定义类型都是可以作为匿名字段去使用

// package main

// import "fmt"

// //人
// type Person struct {
//     name string
//     sex  string
//     age  int
// }

// // 自定义类型
// type mystr string

// // 学生
// type Student struct {
//     Person
//     int
//     mystr
// }

// func main() {
//     s1 := Student{Person{"5lmh", "man", 18}, 1, "bj"}
//     fmt.Println(s1)
// }
// 输出结果：

//     {{5lmh man 18} 1 bj}
// 指针类型匿名字段

// package main

// import "fmt"

// //人
// type Person struct {
//     name string
//     sex  string
//     age  int
// }

// // 学生
// type Student struct {
//     *Person
//     id   int
//     addr string
// }

// func main() {
//     s1 := Student{&Person{"5lmh", "man", 18}, 1, "bj"}
//     fmt.Println(s1)
//     fmt.Println(s1.name)
//     fmt.Println(s1.Person.name)
// }
// 输出结果：

//     {0xc00005c360 1 bj}
//     zs
//     zs

//在Go语言中接口（interface）是一种类型，一种抽象的类型
// interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

// 为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。

// 接口是一个或多个方法签名的集合。
//     任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
//     就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
//     这称为Structural Typing。
//     所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
//     当然，该类型还可以有其他方法。

//     接口只有方法声明，没有实现，没有数据字段。
//     接口可以匿名嵌入其他接口，或嵌入到结构中。
//     对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
//     只有当接口存储的类型和对象都为nil时，接口才等于nil。
//     接口调用不会做receiver的自动转换。
//     接口同样支持匿名字段方法。
//     接口也可实现类似OOP中的多态。
//     空接口可以作为任何类型数据的容器。
//     一个类型可实现多个接口。
//     接口命名习惯以 er 结尾。

// 	type 接口类型名 interface{
//         方法名1( 参数列表1 ) 返回值列表1
//         方法名2( 参数列表2 ) 返回值列表2
//         …
//     }

// //    1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
// type writer interface{
//     Write([]byte) error
// }

// 值接收者和指针接收者实现接口的区别
// 使用值接收者实现接口和使用指针接收者实现接口有什么区别呢？接下来我们通过一个例子看一下其中的区别。

// 我们有一个Mover接口和一个dog结构体。

// type Mover interface {
//     move()
// }

// type dog struct {}
// 1.1.7. 值接收者实现接口
// func (d dog) move() {
//     fmt.Println("狗会动")
// }
// 此时实现接口的是dog类型：

// func main() {
//     var x Mover
//     var wangcai = dog{} // 旺财是dog类型
//     x = wangcai         // x可以接收dog类型
//     var fugui = &dog{}  // 富贵是*dog类型
//     x = fugui           // x可以接收*dog类型
//     x.move()
// }
// 从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

// 1.1.8. 指针接收者实现接口
// 同样的代码我们再来测试一下使用指针接收者有什么区别：

// func (d *dog) move() {
//     fmt.Println("狗会动")
// }
// func main() {
//     var x Mover
//     var wangcai = dog{} // 旺财是dog类型
//     x = wangcai         // x不可以接收dog类型
//     var fugui = &dog{}  // 富贵是*dog类型
//     x = fugui           // x可以接收*dog类型
// }
// 此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。

// type People interface {
//     Speak(string) string
// }

// type Student struct{}

// func (stu *Stduent) Speak(think string) (talk string) {
//     if think == "sb" {
//         talk = "你是个大帅比"
//     } else {
//         talk = "您好"
//     }
//     return
// }

// func main() {
//     var peo People = Student{}
//     think := "bitch"
//     fmt.Println(peo.Speak(think))
// }
//报错
// .\test.go:33:12: undefined: Stduent
// .\test.go:89:6: cannot use Student{} (type Student) as type People in assignment:
//         Student does not implement People (missing Speak method)

// 类型与接口的关系
// 1.2.1. 一个类型实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： Mover接口。

// // Sayer 接口
// type Sayer interface {
//     say()
// }

// // Mover 接口
// type Mover interface {
//     move()
// }
// dog既可以实现Sayer接口，也可以实现Mover接口。

// type dog struct {
//     name string
// }

// // 实现Sayer接口
// func (d dog) say() {
//     fmt.Printf("%s会叫汪汪汪\n", d.name)
// }

// // 实现Mover接口
// func (d dog) move() {
//     fmt.Printf("%s会动\n", d.name)
// }

// func main() {
//     var x Sayer
//     var y Mover

//     var a = dog{name: "旺财"}
//     x = a
//     y = a
//     x.say()
//     y.move()
// }
// 1.2.2. 多个类型实现同一接口
// Go语言中不同的类型还可以实现同一接口 首先我们定义一个Mover接口，它要求必须由一个move方法。

// // Mover 接口
// type Mover interface {
//     move()
// }
// 例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：

// type dog struct {
//     name string
// }

// type car struct {
//     brand string
// }

// // dog类型实现Mover接口
// func (d dog) move() {
//     fmt.Printf("%s会跑\n", d.name)
// }

// // car类型实现Mover接口
// func (c car) move() {
//     fmt.Printf("%s速度70迈\n", c.brand)
// }
// 这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了，不再需要关注它们具体是什么，只需要调用它们的move方法就可以了。

// func main() {
//     var x Mover
//     var a = dog{name: "旺财"}
//     var b = car{brand: "保时捷"}
//     x = a
//     x.move()
//     x = b
//     x.move()
// }
// 上面的代码执行结果如下：

//     旺财会跑
//     保时捷速度70迈
// 并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

// // WashingMachine 洗衣机
// type WashingMachine interface {
//     wash()
//     dry()
// }

// // 甩干器
// type dryer struct{}

// // 实现WashingMachine接口的dry()方法
// func (d dryer) dry() {
//     fmt.Println("甩一甩")
// }

// // 海尔洗衣机
// type haier struct {
//     dryer //嵌入甩干器
// }

// // 实现WashingMachine接口的wash()方法
// func (h haier) wash() {
//     fmt.Println("洗刷刷")
// }
// 1.2.3. 接口嵌套
// 接口与接口间可以通过嵌套创造出新的接口。

// // Sayer 接口
// type Sayer interface {
//     say()
// }

// // Mover 接口
// type Mover interface {
//     move()
// }

// // 接口嵌套
// type animal interface {
//     Sayer
//     Mover
// }
// 嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：

// type cat struct {
//     name string
// }

// func (c cat) say() {
//     fmt.Println("喵喵喵")
// }

// func (c cat) move() {
//     fmt.Println("猫会动")
// }

// func main() {
//     var x animal
//     x = cat{name: "花花"}
//     x.move()
//     x.say()
// }
// 1.3. 空接口
// 1.3.1. 空接口的定义
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

// 空接口类型的变量可以存储任意类型的变量。

// func main() {
//     // 定义一个空接口x
//     var x interface{}
//     s := "pprof.cn"
//     x = s
//     fmt.Printf("type:%T value:%v\n", x, x)
//     i := 100
//     x = i
//     fmt.Printf("type:%T value:%v\n", x, x)
//     b := true
//     x = b
//     fmt.Printf("type:%T value:%v\n", x, x)
// }
// 1.3.2. 空接口的应用
// 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数。

// // 空接口作为函数参数
// func show(a interface{}) {
//     fmt.Printf("type:%T value:%v\n", a, a)
// }
// 空接口作为map的值
// 使用空接口实现可以保存任意值的字典。

// // 空接口作为map值
//     var studentInfo = make(map[string]interface{})
//     studentInfo["name"] = "李白"
//     studentInfo["age"] = 18
//     studentInfo["married"] = false
//     fmt.Println(studentInfo)
// 1.3.3. 类型断言
// 空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

// 接口值
// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

// 我们来看一个具体的例子：

// var w io.Writer
// w = os.Stdout
// w = new(bytes.Buffer)
// w = nil
// 请看下图分解：

// 分解

// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

//     x.(T)
// 其中：

//     x：表示类型为interface{}的变量
//     T：表示断言x可能是的类型。
// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

// 举个例子：

// func main() {
//     var x interface{}
//     x = "pprof.cn"
//     v, ok := x.(string)
//     if ok {
//         fmt.Println(v)
//     } else {
//         fmt.Println("类型断言失败")
//     }
// }
// 上面的示例中如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现：

// func justifyType(x interface{}) {
//     switch v := x.(type) {
//     case string:
//         fmt.Printf("x is a string，value is %v\n", v)
//     case int:
//         fmt.Printf("x is a int is %v\n", v)
//     case bool:
//         fmt.Printf("x is a bool is %v\n", v)
//     default:
//         fmt.Println("unsupport type！")
//     }
// }
// 因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

// 关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
