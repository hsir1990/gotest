//工厂是生产东西，客户来拿东西，买你对不同的客户生产不同的东西。
// 点不同的饭店，生产不同的东西
//函数返回的是接口，子类返回以接口的类型返回
package main

import (
	"fmt"
)

type Restaurant interface {
	GetFood()
}

type Donglaishun struct {
}

func (d *Donglaishun) GetFood() {
	fmt.Println("东来顺的饭菜准备继续")
}

type Qingfeng struct {
}

func (q *Qingfeng) GetFood() {
	fmt.Println("庆丰的包子准备就绪")
}

func NewRestaurant(s string) Restaurant {
	switch s {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil

}

func main() {
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()
	// 	东来顺的饭菜准备继续
	// 庆丰的包子准备就绪

}

// restaurant
// 英 [ˈrestrɒnt]   美 [ˈrestrɑːnt]
// n.
// 餐馆;餐厅


// creation
// 英 [kriˈeɪʃn]   美 [kriˈeɪʃn]  
// n.
// 创造;创作;创建;作品;(尤指《圣经》所述由上帝)创造天地;世界

// creational
// 创建型;创建模式;创建型模式;创造型;其目的可分为创建型