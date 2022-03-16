package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//给Monster绑定方法Store, 可以将一个Monster变量(对象),序列化后保存到文件中

func (this *Monster) Store() bool {

	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	//保存到文件
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

//给Monster绑定方法ReStore, 可以将一个序列化的Monster,从文件中读取，
//并反序列化为Monster对象,检查反序列化，名字正确
func (this *Monster) ReStore() bool {

	//1. 先从文件中，读取序列化的字符串
	filePath := "d:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}

	//2.使用读取到data []byte ,对反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
		return false
	}
	return true
}

//TDD（Test Driven Development）
//Go语言中如何做单元测试和基准测试
// 类型	格式	作用
// 测试函数	函数名前缀为Test	测试程序的一些逻辑行为是否正确
// 基准函数	函数名前缀为Benchmark	测试函数的性能
// 示例函数	函数名前缀为Example	为文档提供示例文档

//go test遍历*_test.go文件，生成临时的main函数和文件
// go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
// -test.run pattern: 只跑哪些单元测试用例

// -test.bench patten: 只跑那些性能测试用例

// -test.benchmem : 是否在性能测试的时候输出内存情况

// -test.benchtime t : 性能测试运行的时间，默认是1s

// -test.cpuprofile cpu.out : 是否输出cpu性能分析文件

// -test.memprofile mem.out : 是否输出内存性能分析文件
//具体的可以看下面链接，或者test 包
//https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95.html

// 目录结构：

//     test
//       |
//        —— calc.go
//       |
//        —— calc_test.go

// func DeepEqual
// func DeepEqual(a1, a2 interface{}) bool
// 用来判断两个值是否深度一致：除了类型相同；在可以时（主要是基本类型）会使用==；但还会比较array、slice的成员，map的键值对，结构体字段进行深入比对。map的键值对，对键只使用==，但值会继续往深层比对。DeepEqual函数可以正确处理循环的类型。函数类型只有都会nil时才相等；空切片不等于nil切片；还会考虑array、slice的长度、map键值对数。
// var b1 bool = reflect.DeepEqual()

// // split/split.go

// package split

// import "strings"

// // split package with a single split function.

// // Split slices s into all substrings separated by sep and
// // returns a slice of the substrings between those separators.
// func Split(s, sep string) (result []string) {
//     i := strings.Index(s, sep)

//     for i > -1 {
//         result = append(result, s[:i])
//         s = s[i+1:]
//         i = strings.Index(s, sep)
//     }
//     result = append(result, s)
//     return
// }
// 在当前目录下，我们创建一个split_test.go的测试文件，并定义一个测试函数如下：

// // split/split_test.go

// package split

// import (
//     "reflect"
//     "testing"
// )

// func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//     got := Split("a:b:c", ":")         // 程序输出的结果
//     want := []string{"a", "b", "c"}    // 期望的结果
//     if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
//         t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
//     }
// }

// // go test -v -run="More"
// -v 显示时间等细节
// -run是正则匹配函数名字

//  //子测试 t.Run()
// func TestSplit(t *testing.T) {
//     type test struct { // 定义test结构体
//         input string
//         sep   string
//         want  []string
//     }
//     tests := map[string]test{ // 测试用例使用map存储
//         "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//         "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//         "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//         "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//     }
//     for name, tc := range tests {
//         t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
//             got := Split(tc.input, tc.sep)
//             if !reflect.DeepEqual(got, tc.want) {
//                 t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//             }
//         })
//     }
// }

// 测试覆盖率
// 我们可以使用go test -cover来查看测试覆盖率

// Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件
// go test -cover -coverprofile=c.out

// // 基准测试
// // 格式
// // func BenchmarkName(b *testing.B){
// //     // ...
// // }

//b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性

// func BenchmarkSplit(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         Split("枯藤老树昏鸦", "老")
//     }
// }

// 基准测试并不会默认执行，需要增加-bench参数，

// go test -bench=Split

// split $ go test -bench=Split
// goos: darwin
// goarch: amd64
// pkg: github.com/pprof/studygo/code_demo/test_demo/split
// BenchmarkSplit-8        10000000               203 ns/op
// PASS
// ok      github.com/pprof/studygo/code_demo/test_demo/split       2.255s
// 其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

// 我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。

//     split $ go test -bench=Split -benchmem
//     goos: darwin
//     goarch: amd64
//     pkg: github.com/pprof/studygo/code_demo/test_demo/split
//     BenchmarkSplit-8        10000000               215 ns/op             112 B/op          3 allocs/op
//     PASS
//     ok      github.com/pprof/studygo/code_demo/test_demo/split       2.394s
// 其中，112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配

// func Split(s, sep string) (result []string) {
//     i := strings.Index(s, sep)

//     for i > -1 {
//         result = append(result, s[:i])
//         s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
//         i = strings.Index(s, sep)
//     }
//     result = append(result, s)
//     return
// }

// // 我们将我们的Split函数优化如下：

// func Split(s, sep string) (result []string) {
//     result = make([]string, 0, strings.Count(s, sep)+1)  //自己先定义内存分配，可以减少系统自己分配内存
//     i := strings.Index(s, sep)
//     for i > -1 {
//         result = append(result, s[:i])
//         s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
//         i = strings.Index(s, sep)
//     }
//     result = append(result, s)
//     return
// }
// 这一次我们提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。我们来看一下这个改进会带来多大的性能提升：

//     split $ go test -bench=Split -benchmem
//     goos: darwin
//     goarch: amd64
//     pkg: github.com/pprof/studygo/code_demo/test_demo/split
//     BenchmarkSplit-8        10000000               127 ns/op              48 B/op          1 allocs/op
//     PASS
//     ok      github.com/pprof/studygo/code_demo/test_demo/split       1.423s
// 这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。

// 性能比较函数
func benchmark(b *testing.B, size int) { /* ... */ }
func Benchmark10(b *testing.B)         { benchmark(b, 10) }
func Benchmark100(b *testing.B)        { benchmark(b, 100) }

// 例如我们编写了一个计算斐波那契数列的函数如下：

// // fib.go

// // Fib 是一个计算第n个斐波那契数的函数
// func Fib(n int) int {
//     if n < 2 {
//         return n
//     }
//     return Fib(n-1) + Fib(n-2)
// }
// 我们编写的性能比较函数如下：

// // fib_test.go

// func benchmarkFib(b *testing.B, n int) {
//     for i := 0; i < b.N; i++ {
//         Fib(n)
//     }
// }

// func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
// func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
// func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
// func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
// func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
// func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
// 运行基准测试：

//     split $ go test -bench=.
//     goos: darwin
//     goarch: amd64
//     pkg: github.com/pprof/studygo/code_demo/test_demo/fib
//     BenchmarkFib1-8         1000000000               2.03 ns/op
//     BenchmarkFib2-8         300000000                5.39 ns/op
//     BenchmarkFib3-8         200000000                9.71 ns/op
//     BenchmarkFib10-8         5000000               325 ns/op
//     BenchmarkFib20-8           30000             42460 ns/op
//     BenchmarkFib40-8               2         638524980 ns/op
//     PASS
//     ok      github.com/pprof/studygo/code_demo/test_demo/fib 12.944s

// 这里需要注意的是，默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。

// 最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。例如：

//     split $ go test -bench=Fib40 -benchtime=20s
//     goos: darwin
//     goarch: amd64
//     pkg: github.com/pprof/studygo/code_demo/test_demo/fib
//     BenchmarkFib40-8              50         663205114 ns/op
//     PASS
//     ok      github.com/pprof/studygo/code_demo/test_demo/fib 33.849s
// 这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了。

// // 错误示范1
// func BenchmarkFibWrong(b *testing.B) {
//     for n := 0; n < b.N; n++ {
//         Fib(n)
//     }
// }

// // 错误示范2
// func BenchmarkFibWrong2(b *testing.B) {
//     Fib(b.N)
// }

// 重置时间
// b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如：

// func BenchmarkSplit(b *testing.B) {
//     time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
//     b.ResetTimer()              // 重置计时器
//     for i := 0; i < b.N; i++ {
//         Split("枯藤老树昏鸦", "老")
//     }
// }

// 并行测试
// func (b B) RunParallel(body func(PB))会以并行的方式执行给定的基准测试。

// RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用。

// func BenchmarkSplitParallel(b *testing.B) {
//     // b.SetParallelism(1) // 设置使用的CPU数
//     b.RunParallel(func(pb *testing.PB) {
//         for pb.Next() {
//             Split("枯藤老树昏鸦", "老")
//         }
//     })
// }
// 执行一下基准测试：

// split $ go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: github.com/pprof/studygo/code_demo/test_demo/split
// BenchmarkSplit-8                10000000               131 ns/op
// BenchmarkSplitParallel-8        50000000                36.1 ns/op
// PASS
// ok      github.com/pprof/studygo/code_demo/test_demo/split       3.308s
// 还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。

// // TestMain
// func TestMain(m *testing.M) {
//     fmt.Println("write setup code here...") // 测试之前的做一些设置
//     // 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
//     retCode := m.Run()                         // 执行测试
//     fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
//     os.Exit(retCode)                           // 退出测试
// }

// //子测试的Setup与Teardown
// // 测试集的Setup与Teardown
// func setupTestCase(t *testing.T) func(t *testing.T) {
//     t.Log("如有需要在此执行:测试之前的setup")
//     return func(t *testing.T) {
//         t.Log("如有需要在此执行:测试之后的teardown")
//     }
// }

// // 子测试的Setup与Teardown
// func setupSubTest(t *testing.T) func(t *testing.T) {
//     t.Log("如有需要在此执行:子测试之前的setup")
//     return func(t *testing.T) {
//         t.Log("如有需要在此执行:子测试之后的teardown")
//     }
// }

// 使用方式如下：

// func TestSplit(t *testing.T) {
//     type test struct { // 定义test结构体
//         input string
//         sep   string
//         want  []string
//     }
//     tests := map[string]test{ // 测试用例使用map存储
//         "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//         "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//         "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//         "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"", "枯藤", "树昏鸦"}},
//     }
//     teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
//     defer teardownTestCase(t)            // 测试之后执行testdoen操作

//     for name, tc := range tests {
//         t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
//             teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
//             defer teardownSubTest(t)           // 测试之后执行testdoen操作
//             got := Split(tc.input, tc.sep)
//             if !reflect.DeepEqual(got, tc.want) {
//                 t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//             }
//         })
//     }
// }
// 测试结果如下：

//     split $ go test -v
//     === RUN   TestSplit
//     === RUN   TestSplit/simple
//     === RUN   TestSplit/wrong_sep
//     === RUN   TestSplit/more_sep
//     === RUN   TestSplit/leading_sep
//     --- PASS: TestSplit (0.00s)
//         split_test.go:71: 如有需要在此执行:测试之前的setup
//         --- PASS: TestSplit/simple (0.00s)
//             split_test.go:79: 如有需要在此执行:子测试之前的setup
//             split_test.go:81: 如有需要在此执行:子测试之后的teardown
//         --- PASS: TestSplit/wrong_sep (0.00s)
//             split_test.go:79: 如有需要在此执行:子测试之前的setup
//             split_test.go:81: 如有需要在此执行:子测试之后的teardown
//         --- PASS: TestSplit/more_sep (0.00s)
//             split_test.go:79: 如有需要在此执行:子测试之前的setup
//             split_test.go:81: 如有需要在此执行:子测试之后的teardown
//         --- PASS: TestSplit/leading_sep (0.00s)
//             split_test.go:79: 如有需要在此执行:子测试之前的setup
//             split_test.go:81: 如有需要在此执行:子测试之后的teardown
//         split_test.go:73: 如有需要在此执行:测试之后的teardown
//     === RUN   ExampleSplit
//     --- PASS: ExampleSplit (0.00s)
//     PASS
//     ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s

// 示例函数的格式
// 被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没有返回值

// 示例函数示例
// 下面的代码是我们为Split函数编写的一个示例函数：

// func ExampleSplit() {
//     fmt.Println(split.Split("a:b:c", ":"))
//     fmt.Println(split.Split("枯藤老树昏鸦", "老"))
//     // Output:
//     // [a b c]
//     // [ 枯藤 树昏鸦]
// }
// 为你的代码编写示例代码有如下三个用处：

//     示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。

//     示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。

//         split $ go test -run Example
//         PASS
//         ok      github.com/pprof/studygo/code_demo/test_demo/split       0.006s
//     示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。下图为stri

// 如何编写压力测试
// 压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：

// 压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母

//     func BenchmarkXXX(b *testing.B) { ... }
// go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数

// 在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行 文件名也必须以_test.go结尾

// 下面我们新建一个压力测试文件webbench_test.go，代码如下所示：

// package gotest

// import (
//     "testing"
// )

// func Benchmark_Division(b *testing.B) {
//     for i := 0; i < b.N; i++ { //use b.N for looping
//         Division(4, 5)
//     }
// }

// func Benchmark_TimeConsumingFunction(b *testing.B) {
//     b.StopTimer() //调用该函数停止压力测试的时间计数

//     //做一些初始化的工作,例如读取文件数据,数据库连接之类的,
//     //这样这些时间不影响我们测试函数本身的性能

//     b.StartTimer() //重新开始时间
//     for i := 0; i < b.N; i++ {
//         Division(4, 5)
//     }
// }
// 我们执行命令go test webbench_test.go -test.bench=".*"，可以看到如下结果：

//     Benchmark_Division-4                            500000000          7.76 ns/op         456 B/op          14 allocs/op
//     Benchmark_TimeConsumingFunction-4            500000000          7.80 ns/op         224 B/op           4 allocs/op
//     PASS
//     ok      gotest    9.364s
// 上面的结果显示我们没有执行任何TestXXX的单元测试函数，显示的结果只执行了压力测试函数，第一条显示了Benchmark_Division执行了500000000次，每次的执行平均时间是7.76纳秒，第二条显示了Benchmark_TimeConsumingFunction执行了500000000，每次的平均执行时间是7.80纳秒。最后一条显示总共的执行时间。
