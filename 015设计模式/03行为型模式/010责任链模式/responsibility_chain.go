//责任链，比如你们遇到一个生产级的大bug，这时候就需要有人背锅，首先是你们的大老板，然后是你们的leader，最后是你，一层一层的锅背下来
//最后一个为空则省略
package main

import (
	"fmt"
	"strconv"
)

type Handler interface {
	Handler(handlerID int) string
}

type handler struct {
	name      string
	next      Handler
	handlerID int
}

func NewHandler(name string, next Handler, handlerID int) *handler {
	return &handler{
		name:      name,
		next:      next,
		handlerID: handlerID,
	}
}

func (h *handler) Handler(handlerID int) string {
	if h.handlerID == handlerID {
		return h.name + " handled " + strconv.Itoa(handlerID)
	}
	if h.next == nil {
		return "11"
	}
	return h.next.Handler(handlerID)
}

func main() {
	wang := NewHandler("lao wang", nil, 1)
	zhang := NewHandler("lao zhang", wang, 2)

	r := wang.Handler(1)
	fmt.Println(r) // lao wang handled 1
	r = zhang.Handler(1)
	fmt.Println(r) //lao wang handled 1
	r = zhang.Handler(2)
	fmt.Println(r) // lao zhang handled 2

}

// chain
// 英 [tʃeɪn]   美 [tʃeɪn]
// n.
// 链子;链条;锁链;一系列，一连串(人或事);连锁商店;约束;连环式
// vt.
// 用锁链拴住(或束缚、固定)

// responsibility
// 英 [rɪˌspɒnsəˈbɪləti]   美 [rɪˌspɑːnsəˈbɪləti]
// n.
// 责任;负责;事故责任;职责;义务;任务

// handler
// 英 [ˈhændlə(r)]   美 [ˈhændlər]
// n.
// 驯兽员;(尤指)驯犬员;搬运工;操作者;组织者;顾问
