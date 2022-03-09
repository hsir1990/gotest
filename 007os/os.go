package main

//os一般指操作系统
//I/O输入/输出(Input/Output)
//bufio主要包含的是带缓冲的IO操作
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//os.File 封装所有文件相关操作，File 是一个结构体

	//func Open(name string) (file *File ,err error)
	//Open打开一个文件用于读取，如果操作成功，返回的文件对象的方法可用于读取数据；
	//对应的文件描述符具有O_RDONLY模式，如果出错，错误底层类型是 *PathError
	//func (f *File) Close error
	// Close关闭文件f，使文件不能用于读写，它返回可能出现的错误

	//打开文件
	//概念说明：file的叫法
	//1.file 叫file对象
	//2.file叫file指针
	//3.file叫file文件句柄
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//输出下文件，看看文件是什么，看出file就是一个指针 *File
	fmt.Printf("file=%v \n", file) //file=&{0xc00007e780}

	// //关闭文件
	// err = file.Close()
	// if err != nil{
	// 	fmt.Println("close file err=",err)
	// }

	//或者 当函数退出时，要及时的关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏

	//创建一个 *Reader，是带缓冲的
	/*
		const(
			defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		// if err == io.EOF { // io.EOF表示文件的末尾
		// 	break
		// }
		//输出内容
		fmt.Print(str)
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		} //奇怪写在这个位置才打印

	}
	fmt.Println("文件读取结束。。。。")

	//读取文件的内容并显示在终端（带缓冲区的方式），使用os.Open, file.Clone, bufio.NewReader(),reader.ReadString 函数和方法。

	//2. 读取文件的内容并显示在终端（使用ioutil一次将整个文件读入到内存中），这种方式适用于文件不大的情况。相关函数 ioutil.ReadFile
	//使用ioutil.ReadFile 一次性将文件读取到位
	// file1 := "F:/work/go/src/gotest/007os/test.txt"
	file1 := "./test.txt"
	content1, err1 := ioutil.ReadFile(file1)
	if err1 != nil {
		fmt.Printf("read file err= %v ", err1)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v",content) // []byte
	fmt.Printf("%v", string(content1))
	//我们没有显示的OPen文件，因此也不需要显示的close文件
	//因为，文件的open和close被封装到ReadFile 函数内部

	//写文件操作应用实例
	// 	func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	// OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
	//第二个参数：文件打开模式（可以组合） 第三个参数：权限控制 （linux）

	//1. 创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件 d:/abce.txt
	// filePath := "d:/abce.txt"
	filePath4 := "./abce.txt"
	file4, err4 := os.OpenFile(filePath4, os.O_WRONLY|os.O_CREATE, 0666)
	if err4 != nil {
		fmt.Printf("open file err=%v\n", err4)
		return
	}
	//及时关闭file句柄
	defer file4.Close()
	//准备写入5句 "hello, Gardon"
	str4 := "hello,Gardon\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer4 := bufio.NewWriter(file4)
	for i := 0; i < 5; i++ {
		writer4.WriteString(str4)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer4.Flush()

	//2.打开一个存在的文件中，将原来的内容覆盖成新的内容10句 "你好，尚硅谷!"
	//创建一个新文件，写入内容 5句 “hello ,Gardon”
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath3 := "./abc.txt"
	file3, err3 := os.OpenFile(filePath3, os.O_WRONLY|os.O_TRUNC, 0666)
	if err3 != nil {
		fmt.Printf("open file err=%v\n", err3)
		return
	}
	//及时关闭file句柄
	defer file3.Close()
	//准备写入5句 "你好,尚硅谷!"
	str3 := "你好,尚硅谷!\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer3 := bufio.NewWriter(file3)
	for i := 0; i < 10; i++ {
		writer3.WriteString(str3)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer3.Flush()

	//3. 打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!'
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath5 := "./abc.txt"
	file5, err5 := os.OpenFile(filePath5, os.O_WRONLY|os.O_APPEND, 0666)
	if err5 != nil {
		fmt.Printf("open file err=%v\n", err5)
		return
	}
	//及时关闭file句柄
	defer file5.Close()
	//准备写入5句 "你好,尚硅谷!"
	str5 := "ABC,ENGLISH!\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer5 := bufio.NewWriter(file5)
	for i := 0; i < 10; i++ {
		writer5.WriteString(str5)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer5.Flush()

	//4. 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath6 := "./test1.txt"
	file6, err6 := os.OpenFile(filePath6, os.O_RDWR|os.O_APPEND, 0666)

	if err6 != nil {
		fmt.Printf("open file err=%v\n", err6)
		return
	}
	//及时关闭file句柄
	defer file6.Close()

	//先读取原来文件的内容，并显示在终端.
	reader6 := bufio.NewReader(file6)
	for {

		str6, err7 := reader6.ReadString('\n')

		//显示到终端
		fmt.Print(str6)
		if err7 == io.EOF { //如果读取到文件的末尾
			break
		}
	}

	//准备写入5句 "你好,尚硅谷!"
	str7 := "hello,北京!\r\n" // \r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer7 := bufio.NewWriter(file6)
	for i := 0; i < 5; i++ {
		writer7.WriteString(str7)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer7.Flush()

	fmt.Println("1111")

	// 5. 将一个文件的内容，写入到另外一个文件。注意 这两个文件已经存在了
	//使用ioutil.ReadFile / ioutil.WriteFile 完成写入文件的任务
	//将d:/abc.txt 文件内容导入到  e:/kkk.txt
	//1. 首先将  d:/abc.txt 内容读取到内存
	//2. 将读取到的内容写入 e:/kkk.txt
	file1Path8 := "./ggg.txt"
	file2Path8 := "./kkk.txt"
	data8, err8 := ioutil.ReadFile(file1Path8)
	if err8 != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n", err8)
		return
	}
	err8 = ioutil.WriteFile(file2Path8, data8, 0666)
	if err8 != nil {
		fmt.Printf("write file error=%v\n", err8)
	}

	//判断文件是否存在  os.Stat()函数返回的错误值进行判断
	// 返回错误为nil，证明文件或者文件夹存在
	//返回 os.IsNotExist()判断为true ，说明文件或者文件夹不存在
	//返回其他类型错误，则不确定是否存在
	// func PathExists(path string) (bool, error){
	// 	_, err9 := os.Stat(path)
	// 	if err9 == nil{ //说明文件存在
	// 		return true ,nil
	// 	}
	// 	if err9==os.IsNotExist(err){
	// 		return false ,nil
	// 	}
	// 	return false ,nil
	// }

	//拷贝文件 将一个图片/电影/MP3 拷贝到另一个文件 中
	//io包中 fun Copy(dst Writer, src Reader)(written int64, err error)
	//将d:/flower.jpg 文件拷贝到 e:/abc.jpg

	//调用CopyFile 完成文件拷贝
	srcFile_1 := "./flower.jpeg"
	dstFile_1 := "./abc.jpeg"
	_, err_1 := CopyFile(dstFile_1, srcFile_1)
	if err_1 == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err_1)
	}

	//6. 统计一个文件中含有的英文，数字，空格及其它字符数量
	//思路: 打开一个文件, 创一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	fileName_2 := "./abc.txt"
	file_2, err_2 := os.Open(fileName_2)
	if err_2 != nil {
		fmt.Printf("open file err=%v\n", err_2)
		return
	}
	defer file_2.Close()
	//定义个CharCount 实例
	var count CharCount
	//创建一个Reader
	reader_2 := bufio.NewReader(file_2)

	//开始循环的读取fileName的内容
	for {
		str_2, err_2 := reader_2.ReadString('\n')
		if err_2 == io.EOF { //读到文件末尾就退出
			break
		}
		//遍历 str ，进行统计
		for _, v := range str_2 {

			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

	//7. os.Args 是一个string的切片，用来存储所有的命令行参数
	fmt.Println("命令行的参数有", len(os.Args))
	//命令行的参数有 4

	//遍历os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v \n", i, v)
	} // ./os.exe  aa bb c:/a/s/a.log

	// args[0] = F:\work\go\path\src\gotest\007os\os.exe
	// args[1] = aa
	// args[2] = bb
	// args[3] = c:/a/s/a.log

	//flage 包用来解析命令行参数
	// 说明：前面的方式是比较原生的方法，对解析参数不是特别方便，特别是带有指定参数形式的命令行

	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" ,就是 -u 指定参数
	//"" , 默认值
	//"用户名,默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	//这里有一个非常重要的操作,转换， 必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v",
		user, pwd, host, port)
	//./os.go -u root -pwd root -h 192.168.0.1 -port 3306
	//user=root  pwd=root host=192.168.0.1 port=3306

}

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	// 	func Open(name string) (file *File, err error)
	// Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)

}
