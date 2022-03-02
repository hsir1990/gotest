/*
 * @Author: Hsir
 * @Date: 2022-03-02 18:07:07
 * @LastEditTime: 2022-03-02 18:30:02
 * @LastEditors: Do not edit
 * @Description: In User Settings Edit
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	suffix := makeSuffix(".jpg") 
	fmt.Println("suffix===>",suffix("nihoa.jpg")) 
	fmt.Println("suffix===>",suffix("nihao")) 
	
}

//编写一个makeSuffix(suffix string) 可以接收一个文件后缀名（比如 .jpg），
// 并返回一个闭包，调用闭包，可以传入一个文件名如果有则直接返回，没有则添加.jpg
//string.HasSuffix,该函数可以判断某个字符串是否有指定的后缀
func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if !strings.HasSuffix(name string) string{
			return name+suffix
		}
		return name
	}
}