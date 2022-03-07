package utilmain

import (
	"fmt"
)

func Bao(b int) {
	fmt.Println("包之间的引用---", b)
}

type student struct {
	Name string
	age  int
}

//当结构体中的变量首字母是小写的时候，外部的不能直接使用
//使用工厂模式实现夸包创建结构体实例（变量）
func NewStudent(n string, a int) *student {
	return &student{
		Name: n,
		age:  a,
	}
}

//如果age字段小写，我们也提供一个方法
func (s *student) GetAge() int {
	return s.age
}




type person struct {
	Name string
	age int   //其它包不能直接访问..
	sal float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}

// func NewPeople(n string) *person{
// 	return &person{
// 		Name:n,
// 	}
// }

//为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age >0 && age <150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
		//给程序员给一个默认值
	}
}

func (p *person) GetAge() int {
	return p.age
}


func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确..")
		
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
