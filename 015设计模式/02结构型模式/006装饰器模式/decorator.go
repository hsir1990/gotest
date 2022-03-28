//装饰者模式解决假如每个方法里面都要加一个日志,这时就可以用装饰者
//装饰要真正调用的方法,把装饰好的方法,放到核心的方法中执行,就是装饰者模式
//wraplogger 是装饰者,装饰log,返回一个函数
//函数中可以包含你的核心方法
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

type piFunc func(int) float64

func WrapLogger(fun piFunc, logger *log.Logger) piFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v,n=%v,result=%v", time.Since(t), n, result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}
	result := 0.0
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}

func main() {
	fmt.Println(Pi(5000))
	fmt.Println(Pi(10000))
	fmt.Println(Pi(50000))
	// 3.141392653591791
	// 3.1414926535900336
	// 3.1415726535897823

	f := WrapLogger(Pi, log.New(os.Stdout, "test ", 1))
	f(100000)
	// test 2022/03/28 took=173.4777ms,n=100000,result=3.141582653589719

}

// decorator
// 英 [ˈdekəreɪtə(r)]   美 [ˈdekəreɪtər]
// n.
// 室内装修设计师;(房屋的)油漆匠，裱糊匠
// 复数： decorators
