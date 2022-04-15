package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//写一个函数，完成登录
func login(userId int, userPwd string) (err error) {

	//下一个就要开始定协议..
	// fmt.Printf(" userId = %d userPwd=%s\n", userId, userPwd)

	// return nil

	//1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务//通过套接字发送消息
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. 创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将loginMes 序列化
	data, err := json.Marshal(loginMes) //对象序列化以后，data是切片
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5. 把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6. 将 mes进行序列化化
	data, err = json.Marshal(mes) //现在是一个byte切片
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 7. 到这个时候 data就是我们要发送的消息
	// 7.1 先把 data的长度发送给服务器
	// 先获取到 data的长度->转成一个表示长度的byte切片  //因为发送消息的时候要先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))                   //获取长度，转成unit32类型
	var buf [4]byte                              //定义一个切片
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //将长度转到切片里面
	// 发送长度
	n, err := conn.Write(buf[:4]) //Write发送的是切片，本身是数组，所以要切一下
	if n != 4 || err != nil {     //发送的长度不是4或者nil不为空会报错，不发送了
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//fmt.Printf("客户端，发送消息的长度=%d 内容=%s", len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data) //然后发送消息本身   //分别写入了数据长度len和数据data，接收时也会在一个函数里面分别接收读取
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20..")
	// 这里还需要处理服务器端返回的消息.  //读取数据包封装成一个readPkg函数
	mes, err = readPkg(conn) // mes 就是

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
