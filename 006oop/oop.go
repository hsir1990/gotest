package main

import (
	"fmt"
	"text/util"
	// utils "text/util"
)

//当结构体中的变量首字母是小写的时候，外部的不能直接使用
//使用工厂模式实现夸包创建结构体实例（变量）
func main() {
	//工厂模式  go的结构体没有构造函数，通常可以使用工厂模式来解决问题
	var stu = utilmain.NewStudent("tom-",18)
	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"age=",utilmain.getAge())
	
	fmt.Println("dayin")
}
