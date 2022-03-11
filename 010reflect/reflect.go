package main

// reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取其动态类型信息，
// 该函数返回一个Type类型值。调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	//n3 := rVal.Float()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}

//专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3. 获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {

	//使用反射机制，编写函数的适配器，桥链接
	//1 反射可以在运行时动态获取变量的各种信息，比如变量的类型（type），类别（kind）
	//2 如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段，方法）
	//3 通过反射，可以修改变量的值，可以调用关联的方法
	//4使用反射，需要import（"reflect"）
	//

	//反射常见应用场景有以下两种
	//1）不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种
	//需要对函数或方法反射。 例如一下这种桥接模式
	//func bridge(funcPrt interface{}, args ...interface())
	//第一个参数 funcPrt 以接口的形式传入函数指针，函数参数args以可变参数的形式传入，
	//bridge函数中可以用反射来动态执行funcPtr函数
	// 2）对结构体序列化时，如果结构体有指定Tag，也会使用到反射生成对应的字符串
	// type Monster struct{
	// 	Name string `json:"name"`
	// }

	//反射函数
	//1 reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型
	//2 reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型reflect.Value 是一个结构体类型。
	//通过reflect.Value,可以获取到关于该变量的很多信息
	//3 变量，interface{} 和 reflect.Value是可以互相转换的，这点在实际开发中，会经常使用到。

	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1. 先定义一个int
	var num int = 100
	reflectTest01(num)

	//2. 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)

	// 	rType= int
	// n2= 102
	// rVal=100 rVal type=reflect.Value
	// num2= 100
	// rType= main.Student
	// kind =struct kind=struct
	// iv={tom 20} iv type=main.Student
	// stu.Name=tom

	//1 Type和Kinb的
	// Type是类型，Kind是类别， Type和Kind可能是相同，也可能是不同的
	// 比如： var num int = 10 num 的Type是int， Kind 也是int
	//  var stu Student stu的Type是pkg1.Student ， Kind是struct

	//3 通过反射可以在让变量在interface{}和Reflect.Value 之间相互转换，这点在前面画过
	// 示意图并在快速入门案例中讲解过，这里我看下如何在代码中体现
	// 变量 《======》 interface{} 《========》reflect.Value
	//4 使用反射的方式来获取变量的值（并返回对应的类型），要求数据类型匹配，比如x 是 int，
	// 那么就应该使用 reflect.Value(x).Int(),而不能使用其它的，否则报panic
	// func (v Value) Int() int64
	// 返回v持有的有符号整数（表示为int64），如果V的Kind不是int，int8，int16，int32，int64会panic
	// func testInt(b interface{}){
	// 	val := reflect.ValueOf(b)
	// 	fmt.Printf("v = %v \n", val.Int())
	// 	fmt.Printf("v = %v \n", val.Float()) // error
	// }
	//5 通过反射的来修改变量，注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样才能改变传入的变量的值，
	//同时需要使用到reflect.Value.Elem() 方法

	// //通过反射，修改,
	// // num int 的值
	// // 修改 student的值

	// func reflect01(b interface{}) {
	// 	//2. 获取到 reflect.Value
	// 	rVal := reflect.ValueOf(b)
	// 	// 看看 rVal的Kind是
	// 	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	// 	//3. rVal
	// 	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	// 	rVal.Elem().SetInt(20)
	// }
	// func main() {

	// 	var num int = 10
	// 	reflect01(&num)
	// 	fmt.Println("num=", num) // 20

	// 	//你可以这样理解rVal.Elem()
	// 	// num := 9
	// 	// ptr *int = &num
	// 	// num2 := *ptr  //=== 类似 rVal.Elem()
	// }

	// 6reflect.Value.Elem() 应该如何理解？

	// reflect.Value.Elem() //用于获取指针

	// var num int = 10
	// fn := reflect.ValueOf(&num)
	// fn.Elem().SetInt(200)
	// fmt.Printf("%v \n", num)
	
	// // fn.Elem()用于获取指针变量,类似下面
	// var num =10
	// var b *int = &num
	// *b = 3

	//给你一个变量，  var v float64 = 12 ，请使用反射来得到它的reflect.Value,然后获取对应的Type，
	// Kind 和值，并将reflect.Value 转换成interface{}，再将interface{} 转换成float64

	var str string = "tom"      //ok
	fs := reflect.ValueOf(&str) //ok fs -> string
	fs.Elem().SetString("jack") //ok
	fmt.Printf("%v\n", str)     // jack

	// 使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
	//
	//创建了一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	//将Monster实例传递给TestStruct函数
	TestStruct(a)
	// 2使用反射的方法来获取结构体的tag标签，遍历字段的值，修改字段值，调用结构体方法（要求：
	// 通过传递地址的方式完成，在前面案例上修改即可）
	// 3定义了两个函数test1和test2，定义一个适配器数用作统一处理接口
	// 4 使用反射操作任意结构体类型
	// 5 使用反射创建并操作结构体

}

//定义了一个Monster结构体
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法， 接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//方法，显示s的值
func (s Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}
func TestStruct(a interface{}) {
	//获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	//如果传入的不是struct，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()

	fmt.Printf("struct has %d fields\n", num) //4
	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	//方法的排序默认是按照 函数名的排序（ASCII码）
	val.Method(1).Call(nil) //获取到第二个方法。调用它

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/

}
