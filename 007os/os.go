package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)
func main(){
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
	file , err := os.Open("F:/work/go/src/gotest/007os/test.txt")
	if err != nil{
		fmt.Println("open file err=",err)
	}
	//输出下文件，看看文件是什么，看出file就是一个指针 *File
	fmt.Printf("file=%v \n",file) //file=&{0xc00007e780}

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
	for{
		str,err := reader.ReadString('\n') //读到一个换行就结束
		// if err == io.EOF { // io.EOF表示文件的末尾
		// 	break
		// }
		//输出内容
		fmt.Print(str)
		if err == io.EOF{ //io.EOF表示文件的末尾
			break
		}  //奇怪写在这个位置才打印

		
	}
	fmt.Println("文件读取结束。。。。")


	//读取文件的内容并显示在终端（带缓冲区的方式），使用os.Open, file.Clone, bufio.NewReader(),reader.ReadString 函数和方法。


	//2. 读取文件的内容并显示在终端（使用ioutil一次将整个文件读入到内存中），这种方式适用于文件不大的情况。相关函数 ioutil.ReadFile
	//使用ioutil.ReadFile 一次性将文件读取到位
	// file1 := "F:/work/go/src/gotest/007os/test.txt"
	file1 := "./test.txt"
	content1, err1 := ioutil.ReadFile(file1);
	if err1 != nil{
		fmt.Printf("read file err= %v ",err1)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v",content) // []byte
	fmt.Printf("%v",string(content1))
	//我们没有显示的OPen文件，因此也不需要显示的close文件
	//因为，文件的open和close被封装到ReadFile 函数内部

	//写文件操作应用实例
// 	func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
// OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
//第二个参数：文件打开模式（可以组合） 第三个参数：权限控制 （linux）

}