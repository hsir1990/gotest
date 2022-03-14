//golang的主要设计目标之一就是面向大规模后端服务程序,网络通讯这块是服务端程序必不可少的也是至关重要的一部份

//网络编程有两种
//1) TCP socket 编程,是网络编程的主流.之所以叫TCP socket 编程,是因为底层是基于Tcp/ip协议的,比如  QQ 聊天
//2) b/s结构的http编程,我们使用浏览器去访问服务器时,使用的就是http协议,而http底层依旧是用tcp socket 实现的

// 协议(tcp/ip)
//TCP/IP(Transmission Control Protocol/Internet Protocol)的简写,中文译名为传输控制协议/因特网互联协议, 又叫网络通讯协议,
//这个协议是Internet最基本的协议,Internet国际互联网络的基础,简单地说,就是由网络层的IP协议和传输层的TCP协议组成的.

//具体可看18,19章

//加版本打标签，为了给测试，get文件

// $ git commit -m "Emphasize our friendliness" say.go
// $ git tag v1.0.1
// $ git push --tags origin v1