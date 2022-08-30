package main 

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main(){

	wg.Add(2)
	ch := make(chan int)
	go sender(ch)
	go recver(ch)
	wg.Wait()
}

func sender(ch chan int){
	defer wg.Done()
	for i:=0;i<10;i++ {
		time.Sleep(time.Second * 2 )
		fmt.Println("sender:", i)
		ch <- i
	}
}
func recver(ch chan int){
	defer wg.Done()
	// i:=0
	// for a := range ch{
	// 	fmt.Println("recver:",a)
	// 	i++
	// 	if i>10{
	// 		break
	// 	}
	// }


//这样也能执行
	// var recv int
	// for {
	// 	recv = <-ch  
	// 	fmt.Println(recv)
	// }

	//这样也行

	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v, ok)
			// if v == "tuner" {
			// 	break
			// }
		}
	}

}




// sender: 0
// reciver: 0
// sender: 1
// reciver: 1
// sender: 2
// reciver: 2
// sender: 3
// reciver: 3
// sender: 4
// reciver: 4
// sender: 5
// reciver: 5
// sender: 6
// reciver: 6
// sender: 7
// reciver: 7
// sender: 8
// reciver: 8
// sender: 9
// reciver: 9

