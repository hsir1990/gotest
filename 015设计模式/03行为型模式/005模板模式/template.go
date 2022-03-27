//模板模式，要入职一个公司，给你一个表格，表格里面填写地址姓名等各种信息，为的是解决这类信息
// 工作经历等各种信息，都已经打印好了，你直接往里面填就行
// 多态使用，让coder通过worklterface去传入参数
package main

import (
	"fmt"
)

type WorkInterface interface{
	GetUp()
	Work()
	Sleep()
}

type Worker struct{
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker{
	return &Worker{w}
}

func (w *Worker) Daily(){
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct{

}

func (c *Coder) GetUp(){
	fmt.Println("coder GetUp")
}

func (c *Coder) Work(){
	fmt.Println("coder Work")
}

func (c *Coder) Sleep(){
	fmt.Println("coder Sleep")
}

func main() {
	Worker := NewWorker(&Coder{})  //因为引用的是接口，使用的这个c *Coder，所以使用是引用类型
	// Worker := NewWorker(Coder{})  //这样会报错
	Worker.Daily()
// 	coder GetUp
// coder Work
// coder Sleep

}

// daily
// 英 [ˈdeɪli]   美 [ˈdeɪli]  
// adj.
// 每日的;日常的;每个工作日的;按日的
// adv.
// 每日;每天
// n.
// (除星期日外每日发行的)日报;(不寄宿的)仆人