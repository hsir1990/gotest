package main

import (
	"fmt"
	 "gotest/util"
	// utils "gotest/util"
)
type A struct {
	Name string
	age int
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
	type Goods struct{
		Name string
		Age int
	}
	type Book struct{
		Goods  //这里就是嵌套匿名结构体Goods
		Writer string
	}

	//结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段，方法，都可以使用
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


	//当我们直接通过b访问字段或者方法时，其执行流程如下  比如b.Name
	//编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段
	//如果没有就去看B中嵌入的匿名结构体：A有没有声明Name字段，如果有就调用，如果没有继续查找，。。。如果找不到就报错
	//当结构体和匿名结构体都有相同字段或者方法时，编译器采用就近访问原则访问，如果访问匿名结构体的字段和方法，可以通过匿名结构体名来区分如下

	var b B
	b.Name = "jack" // ok
	b.A.Name = "scott"
	b.age = 100  //ok
	b.SayOk()  // B SayOk  jack
	b.A.SayOk() //  A SayOk scott
	b.hello() //  A hello ? "jack" 还是 "scott"

	//结构体嵌入两个（或多个）匿名结构体，如果两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)
	//在访问时，就必须明确指定匿名结构体名字，否则编译报错



}
