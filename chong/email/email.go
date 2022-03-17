package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

// 爬邮箱
func GetEmail() {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	//fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)

	// 遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	GetEmail()
}

// email: 1184822807@qq.com
// qq: 1184822807
// email: 1184822807@qq.com
// qq: 1184822807
// email: 598088118@qq.com
// qq: 598088118
// email: 598088118@qq.com
// qq: 598088118
// email: 835428013@qq.com
// qq: 835428013
// email: 835428013@qq.com
// qq: 835428013
// email: 934555653@qq.com
// qq: 934555653
// email: 934555653@qq.com
// qq: 934555653
// email: 695527214@qq.com
// qq: 695527214
// email: 695527214@qq.com
// qq: 695527214
// email: 410708076@qq.com
// qq: 410708076
// email: 410708076@qq.com
// qq: 410708076
// email: 55867960@qq.com
// qq: 55867960
// email: 55867960@qq.com
// qq: 55867960
// email: 2044989269@qq,com
// qq: 2044989269
// email: 2044989269@qq,com
// qq: 2044989269
// email: 823104483@qq.com
// qq: 823104483
// email: 823104483@qq.com
// qq: 823104483
