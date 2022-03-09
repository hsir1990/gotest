package cal

import (
	_ "fmt"
	"testing" //引入go 的testing框架包
)

//编写要给测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T) {

	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")

}

//编写要给测试用例，去测试addUpper是否正确
func TestGetSub(t *testing.T) {

	//调用
	res := getSub(10, 3)
	if res != 7 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("getSub(10, 3) 执行错误，期望值=%v 实际值=%v\n", 7, res)
	}

	//如果正确，输出日志
	t.Logf("getSub(10, 3) 执行正确!!!!...")

}

// func TestHello(t *testing.T) {

// 	fmt.Println("TestHello被调用..")

// }

//1) 测试用例文件必须以 _test.go结尾。比如cal_test.go, cal 不是固定的
//2）测试用例函数比如以Test开头，一般来说就是 Test+被测试的函数名，比如 TestAddUpper
//3) TestAddUpper(t *testing.T) 的形式类型必须时  *testing.T
//4) 一个测试用例文件中，可以有多个测试用例函数，比如 TestAddUpper,  TestSub
//5) 运行测试用例指令
//go test [如果运行正确，无日志，错误时，会输出日志]
//go test -v [运行正确或是错误，都输出日志]
//6）当出现错误时，可以使用t.FatalF 来格式化输出错误信息，并推出程序
//8）t.Logf 方法可以输出相应的日志
//9）PASS表示测试用例运行成功，FAIL表示测试用例运行失败
// 10）测试单个文件，一定要带上被测试的原文件
//go test -v cal_test.go cal.go
// 11)测试单个方法
// go test -v -test.run TestAddUpper


