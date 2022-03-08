package main

import (
	"fmt"
	"gotest/util"
	// utils "gotest/util"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

type D struct {
	Name string
	age  int
}
type E struct {
	Name  string
	Score float64
}
type C struct {
	D
	E
	//Name string
}

type F struct {
	d D //有名结构体  //组合关系
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}

type G struct {
	Monster
	int
	n int
}

//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

//声明/定义一个接口
type Usb2 interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
	Test()
}

type Phone struct {
}

//让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
}

//让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作~~~。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

//计算机
type Computer struct {
}

//编写一个方法Working 方法，接收一个Usb接口类型变量
//只要是实现了 Usb接口 （所谓实现Usb接口，就是指实现了 Usb接口声明  所有 方法）
func (c Computer) Working(usb Usb) { //use变量会根据传入的实参，来判断是Phone，还是Camera   //既可以接受Phone又可以接受Camera ，就体现了多态

	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

// type BInterface interface {
// 	test01()
// }

// type CInterface interface {
// 	test02()
// }

// type AInterface interface {
// 	BInterface
// 	CInterface
// 	test03()
// }

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}
type Monster1 struct {
}

func (m Monster1) Hello() {
	fmt.Println("Monster1 Hello()~~")
}

func (m Monster1) Say() {
	fmt.Println("Monster1 Say()~~")
}

type DInterface interface {
	test01()
}

type EInterface interface {
	test02()
}

type FInterface interface {
	DInterface
	EInterface
	test03()
}

//如果需要实现FInterface,就需要将DInterface EInterface的方法都实现
type Stu2 struct {
}

func (stu Stu2) test01() {

}
func (stu Stu2) test02() {

}
func (stu Stu2) test03() {

}

type T interface {
}


//定义Student类型
type Student struct {

}

//当结构体中的变量首字母是小写的时候，外部的不能直接使用
//使用工厂模式实现夸包创建结构体实例（变量）
func main() {
	//工厂模式  go的结构体没有构造函数，通常可以使用工厂模式来解决问题
	var stu = utilmain.NewStudent("tom-", 18)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "age=", stu.GetAge())

	fmt.Println("dayin")

	//go对封装做了简化，不和java一样必须写成GetXxx和SetXxx
	p := utilmain.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())

	//继承
	//嵌套匿名结构体
	// type Goods1 struct {
	// 	Name string
	// 	Age  int
	// }
	// type Book struct {
	// 	Goods1  //这里就是嵌套匿名结构体Goods1
	// 	Writer string
	// }

	//3 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段，方法，都可以使用
	// var b B
	// b.A.Name = "tom"
	// b.A.age = 19
	// b.A.SayOk()
	// b.A.hello()

	// //上面的写法可以简化，匿名结构体的简化

	// b.Name = "smith"
	// b.age = 20
	// b.SayOk()
	// b.hello()

	//4 当我们直接通过b访问字段或者方法时，其执行流程如下  比如b.Name
	//编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段
	//如果没有就去看B中嵌入的匿名结构体：A有没有声明Name字段，如果有就调用，如果没有继续查找，。。。如果找不到就报错
	//当结构体和匿名结构体都有相同字段或者方法时，编译器采用就近访问原则访问，如果访问匿名结构体的字段和方法，可以通过匿名结构体名来区分如下

	var b B
	b.Name = "jack" // ok
	b.A.Name = "scott"
	b.age = 100 //ok
	b.SayOk()   // B SayOk  jack
	b.A.SayOk() //  A SayOk scott
	b.hello()   //  A hello ? "jack" 还是 "scott"  //注意：  A hello scott  //就近访问原则

	//结构体嵌入两个（或多个）匿名结构体，如果两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)
	//在访问时，就必须明确指定匿名结构体名字，否则编译报错

	var c C
	//如果c 没有Name字段，而D 和 E有Name, 这时就必须通过指定匿名结构体名字来区分
	//所以 c.Name 就会包编译错误， 这个规则对方法也是一样的！
	// c.Name = "tom" //报错
	c.D.Name = "tom" // error
	fmt.Println("c")

	//多重继承
	// 如果一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承
	//不过为了代码的简介，建议大家尽量不要使用多重继承

	//5.如果一个struct嵌套了一个有名结构体，这种模式是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
	//如果F 中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
	//比如 f.d.Name
	var f F
	f.d.Name = "jack"

	//6 嵌套匿名结构体后，也可以创结构体变量（实例）时，直接指定各个匿名结构体字段的值
	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}

	//演示访问Goods的Name
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)

	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}

	fmt.Println("tv", tv)   //tv {{电视机001 5000.99} {海尔 山东}}
	fmt.Println("tv2", tv2) //tv2 {{电视机002 5000.99} {夏普 北京}}

	tv3 := TV2{&Goods{"电视机003", 7000.99}, &Brand{"创维", "河南"}}

	tv4 := TV2{
		&Goods{
			Name:  "电视机004",
			Price: 9000.99,
		},
		&Brand{
			Name:    "长虹",
			Address: "四川",
		},
	}

	fmt.Println("tv3", *tv3.Goods, *tv3.Brand) //tv3 {电视机003 7000.99} {创维 河南}
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand) //tv4 {电视机004 9000.99} {长虹 四川}

	//如果一个结构体有int类型的匿名字段，就不能有第二个。
	//如果需要多个int字段，就必须给int字段指定名字
	//演示一下匿名字段时基本数据类型的使用
	var g G
	g.Name = "狐狸精"
	g.Age = 300
	g.int = 20
	g.n = 40
	fmt.Println("g=", g) // g= {{狐狸精 300} 20 40}

	//接口 interface  作用为了统一规范，多态特性主要是通过接口来体现
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
	// 手机开始工作。。。
	// 手机停止工作。。。
	// 相机开始工作~~~。。。
	// 相机停止工作。。。

	//interface定义了一组方法，但是这些不需要实现，并且interface不能包含任何变量
	//到某个自定义类型（比如结构体Phone）要使用的时候，在根据具体情况把这些方法写出来（实现）

	//注意
	//接口里的所有方法都是没有方法体的，即接口的方法都是没有实现的方法。接口体现了程序设计的 多态 和 高内聚低耦合 的思想
	//golang 中的接口，不需要显式的实现。
	//只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口，
	//因此，golang中没有implement

	//注意
	//1. 接口本身不能创建实例，但是 可以指向一个实现了该接口的自定义类型的变量（实例）
	var stu1 Stu //结构体变量，实现了 Say() 实现了 AInterface
	var a AInterface = stu1
	a.Say() //Stu Say()

	//2 接口中接口中所有的方法都没有方法体，即都是没有实现的方法。
	//3 在golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口。
	//4 一个自定类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给接口类型
	//5 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
	var i integer = 10
	var b1 AInterface = i
	b1.Say() // integer Say i = 10

	//6 一个自定义类型可以实现多个接口
	//Monster实现了AInterface 和 BInterface
	var monster Monster1
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()   //Monster1 Say()~~
	b2.Hello() // Monster1 Hello()~~

	//7  golang 接口中不能有任何变量

	// type AInterface interface{
	// 	Name string //报错
	// 	Test()
	// }

	//8  一个接口（比如A接口）可以继承多个别的接口（比如B，C接口），这时如果要实现A接口，
	//也必须将B，C接口的方法全部实现
	var stu2 Stu2
	var f1 FInterface = stu2
	f1.test01()

	//9 interface 类型默认是一个指针（引用类型），如果没有对interface初始化就使用，那么会输出nil
	//空接口interface{} 没有任何方法，所以所有类型都实现了空接口，即我们可以把任何一个变量赋给空接口。

	var t T = stu2 //ok
	fmt.Println(t) //{}

	var t2 interface{} = stu2
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t) //8.8 8.8

	//sort.Sort(data Interface)
	//Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。
	//本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）

	//交换可以这样用
	// hs[i],hs[j] = hs[j],hs[i]

	//实现接口 VS 继承
	//1）当A结构体继承了B结构体，那么A结构就自动的继承了B结构体的字段和方法，并且可以直接使用
	//2）当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此我们
	//可以认为：实现接口是对继承机制的补充

	//创建一个LittleMonkey 实例
	//这种写法可以
	// monkey := LittleMonkey{
	// 	Monkey {
	// 		Name : "悟空",
	// 	},
	// }

	//这种写法不对
	// monkey := LittleMonkey{

	// 		Name : "悟空",

	// }
	//这种写法也可以
	var monkey LittleMonkey
	monkey.Name = "悟空"
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
	// 	悟空  生来会爬树..
	// 悟空  通过学习，会飞翔...
	// 悟空  通过学习，会游泳..

	//继承的价值主要在于：解决代码的复用性和可维护性
	//接口的价值主要在于：设计，设计好各种规范（方法），让其它自定义类型去实现这些方法

	//接口比继承更加灵活，继承是满足is -a 的关系，而接口只需满足  like -a的关系

	//接口在一定程度上实现了代码解耦

	//多态  两种形式
	//1.多态参数，上面说的Usb即可以接收手机变量，又可以接收相机变量，就体现了Usb接口的  多态
	//2.多态数组
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb1
	usbArr[0] = Phone1{"vivo"}
	usbArr[1] = Phone1{"小米"}
	usbArr[2] = Camera1{"尼康"}

	fmt.Println(usbArr) //[{vivo} {小米} {尼康}]

	//类型断言
	//如何将一个接口变量，赋给自定义类型的变量==》引出类型变量

	var x13 interface{}
	var b13 float32 = 1.1
	x13 = b13 //空接口，可以接收任意类型

	//x1 => float32

	y13 := x13.(float32)
	fmt.Printf("y的类型是 %T 值是=%v", y13, y13) //y的类型是 float32 值是=1.1

	// 进行类型断言时，如果不匹配，就会报panic，因此进行类型断言时，要确保原来的空接口指向的就是断言的类型
	//如果断言时加上检测机制，可以防止报错

	var x15 interface{}
	var b15 float32 = 2.1
	x15 = b15

	//类型断言（带检测的）
	if y15, ok15 := x15.(float32); ok15 {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y15, y15)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")

	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArrType [3]UsbType
	usbArrType[0] = PhoneType{"vivo"}
	usbArrType[1] = PhoneType{"小米"}
	usbArrType[2] = CameraType{"尼康"}
	// 手机开始工作。。。
	// 手机 在打电话..
	// 手机停止工作。。。
	
	// 手机开始工作。。。
	// 手机 在打电话..
	// 手机停止工作。。。
	
	// 相机开始工作。。。
	// 相机停止工作。。。
	


	//遍历usbArrType
	//Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb 接口声明的方法外，还需要调用PhoneType 特有方法 call. =》类型断言
	var computerT ComputerType
	for _, v := range usbArrType {
		computerT.Working(v)
		fmt.Println()
	}
	//fmt.Println(usbArr)


	//写一个函数，循环判断传入参数的类型
	var n_1 float32 = 1.1
	var n_2 float64 = 2.3
	var n_3 int32 = 30
	var n_ame string = "tom"
	address := "北京"
	n_4 := 300

	stu_1 := Student{}
	stu_2 := &Student{}

	TypeJudge(n_1, n_2, n_3, n_ame, address, n_4, stu_1, stu_2)
	// 第0个参数是 float32 类型，值是1.1
	// 第1个参数是 float64 类型，值是2.3
	// 第2个参数是 整数 类型，值是30
	// 第3个参数是 string 类型，值是tom
	// 第4个参数是 string 类型，值是北京
	// 第5个参数是 整数 类型，值是300
	// 第6个参数是 Student 类型，值是{}
	// 第7个参数是 *Student 类型，值是&{}

}

//写一个函数，循环判断传入参数的类型
func TypeJudge(items... interface{}){
	for index, val := range items{
		switch val.(type){
			case bool :
				fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, val)
			case float32 :
				fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, val)
			case float64 :
				fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, val)
			case int, int32, int64 :
				fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, val)
			case string :
				fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, val)
			case Student :
				fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, val)
			case *Student :
				fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, val)
			default :
				fmt.Printf("第%v个参数是  类型 不确定，值是%v\n", index, val)
		}
			

	}
}




//声明/定义一个接口
type UsbType interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type PhoneType struct {
	name string
}

//让PhoneType 实现 UsbType接口的方法
func (p PhoneType) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p PhoneType) Stop() {
	fmt.Println("手机停止工作。。。")
}

func (p PhoneType) Call() {
	fmt.Println("手机 在打电话..")
}

type CameraType struct {
	name string
}

//让CameraType 实现   UsbType接口的方法
func (c CameraType) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c CameraType) Stop() {
	fmt.Println("相机停止工作。。。")
}

type ComputerType struct {
}

func (computer ComputerType) Working(usb UsbType) {
	usb.Start()
	//如果usb是指向PhoneType结构体变量，则还需要调用Call方法
	//类型断言..[注意体会!!!]
	if phone, ok := usb.(PhoneType); ok {
		phone.Call()
	}
	usb.Stop()
}

//Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树..")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔...")
}

//让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳..")
}

//声明/定义一个接口
type Usb1 interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone1 struct {
	name string
}

//让Phone1 实现 Usb1接口的方法
func (p Phone1) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone1) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera1 struct {
	name string
}

//让Camera1 实现   Usb1接口的方法
func (c Camera1) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera1) Stop() {
	fmt.Println("相机停止工作。。。")
}
