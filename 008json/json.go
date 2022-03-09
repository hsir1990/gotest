package main

import (
	"encoding/json"
	"fmt"
)

//对于结构体的序列化，如果我们希望序列化后的key的名字，又我们自己重新定制，那可可以给stuct指定一个tag标签
//定义一个结构体
type Monster struct {
	Name     string `json:"monster_name"` //反射机制
	Age      int    `json:"monster_age"`
	Birthday string //....
	Sal      float64
	Skill    string
}

//对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

//演示将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	//定义一个slice
	var slice []map[string]interface{}
	//反序列化，不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}
func main() {
	// JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。  2001年就开始推广了
	//在网络上传输时先将数据（结构体，map，切片等）序列化成json字符串，到接收方得到json字符串时，在反序列化恢复成原来的数据类型（结构体，map，切片等）。

	//在js（质疑中---）语言中，一切都是对象，因此，任何数据类型都可以通过JSON来表示，例如字符串，数字，对象，数组，map，结构体等

	//序列化
	//演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	//将monster 序列化
	data, err := json.Marshal(&monster) //序列化出来时切片的形式
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
	//monster序列化后={"monster_name":"牛魔王","monster_age":500,"Birthday":"2011-11-11","Sal":8000,"Skill":"牛魔拳"}

	testFloat64() //演示对基本数据类型的序列化

	//反序列化
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"  //需要转义 "
	var monster1 Monster
	err1 := json.Unmarshal([]byte(str), &monster1)
	if err1 != nil {
		fmt.Printf("unmarshal err=%v\n", err1)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster1, monster1.Name)

	unmarshalSlice()

	

}
