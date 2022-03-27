//迭代器比较简单，取到他的索引取到他的值，通过索引移动
// 有hasNext和Next
package main

import (
	"fmt"
)

type Iterator interface{
	Index() *int  //注意接口使用，方法返回的是指针类型，结构的方法返回的也需要时指针类型才能继承，所以不能这样写Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

type ArrayIterator struct{
	array []interface{}
	index *int
}

func (a *ArrayIterator) Index() *int{
	return a.index
}

func (a *ArrayIterator) Value() interface{}{
	return a.array[*(a.index)]  //*a.index 等于*(a.index)
	// return a.array[*((*a).index)] //这样也对
	// return (*a).array[*((*a).index)] //这样也对
	// return a.array[a.index] //这样会报错
}

func (a *ArrayIterator) HasNext() bool{
	return *a.index+1<=len(a.array)
}

func (a *ArrayIterator) Next(){
	if a.HasNext(){
		*a.index++
	}
}
//测试
// type Itera1 interface{
// 	Index1() 
// }
// type  Itera11 struct{}
// func (i Itera11) Index1() {
// 	fmt.Println("1111")
// }

func main() {
	array := []interface{}{1,3,5,9,5,2}
	a := 0
	// var iterator *Iterator = &ArrayIterator{array:array, index:&a}//报错，因为在这接口不用使用指针再定义一边了// cannot use &ArrayIterator{...} (type *ArrayIterator) as type *Iterator in assignment:
	// var iterator Iterator = ArrayIterator{array:array, index:&a} //这样也报错，因为需要使用地址传递 cannot use ArrayIterator{...} (type ArrayIterator) as type Iterator in assignment:
	// var iterator ArrayIterator = ArrayIterator{array:array, index:&a}//这样也是对的

	var iterator Iterator = &ArrayIterator{array:array, index:&a} //因为Iterator是接口，所以不用定义为指针类型
	for it := iterator; iterator.HasNext();iterator.Next(){
		index,value := it.Index(),it.Value().(int)
		if value != array[*index]{
			fmt.Println("error...")
		}
		fmt.Println(*index, value)
	}
	fmt.Println(iterator)
	// 0 1
	// 1 3
	// 2 5
	// 3 9
	// 4 5
	// 5 2
	// &{[1 3 5 9 5 2] 0xc00000a0b8}

	//测试
	// var it Iterator = &ArrayIterator{}
	// var it Itera1 = Itera11{}
	// fmt.Println(it)
	
	

}

// iterator
// 【计】迭代器，迭代程序