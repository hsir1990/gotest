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
func (s *student) getAge() int {
	return s.age
}
