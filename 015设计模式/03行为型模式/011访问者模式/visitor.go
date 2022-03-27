//访问者中 new一个struct是为了给一个地址空间
//模式是有一个目标，当访问者来的时候，目标会做出一些反应，实现和封装的一个分离
package main

import (
	"fmt"
)

type IVisitor interface{
	Visit()
}

type WeiBoVisitor struct{

}

func (w *WeiBoVisitor) Visit(){
	fmt.Println("Visit WeiBo")
}

type IQIYIVisitor struct{

}

func (i IQIYIVisitor) Visit(){
	fmt.Println("Visitor IQiYi")
}

type IElement interface{
	Accept(visitor IVisitor)
}

type Element struct{

}

func (e Element) Accept(v IVisitor){
	v.Visit()
}


func main() {
	e := new(Element)
	// e := Element{}//这样也可以
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
	// Visit WeiBo
	// Visitor IQiYi
	
}

// element
// 英 [ˈelɪmənt]   美 [ˈelɪmənt]  
// n.
// 要素;基本部分;典型部分;少量;有点;有些;(大团体或社会中的)一组，一群，一伙;元素（如金、氧、碳）;（尤指恶劣的）天气;（学科的）基本原理，基础，纲要;（尤指动物的）自然环境，适宜的环境;电热元件



// accept
// 英 [əkˈsept]   美 [əkˈsept]  
// v.
// (认为合适或足够好而)接受;接受(建议、邀请等);同意;收受;认可;承认，承担（责任等）;相信（某事属实）;容忍，忍受（困境等）;欢迎;接纳，接受