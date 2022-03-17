fmt
1. fmt
fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。

1.1.1. 向外输出
标准库fmt提供了以下几种输出相关函数。

Print
Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。

func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
举个简单的例子：

func main() {
    fmt.Print("在终端打印该信息。")
    name := "枯藤"
    fmt.Printf("我是：%s\n", name)
    fmt.Println("在终端打印单独一行显示")
}
执行上面的代码输出：

    在终端打印该信息。我是：枯藤
    在终端打印单独一行显示
Fprint
Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。

func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
举个例子：

// 向标准输出写入内容
fmt.Fprintln(os.Stdout, "向标准输出写入内容")
fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
    fmt.Println("打开文件出错，err:", err)
    return
}
name := "枯藤"
// 向打开的文件句柄中写入内容
fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
注意，只要满足io.Writer接口的类型都支持写入。

Sprint
Sprint系列函数会把传入的数据生成并返回一个字符串。

func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
简单的示例代码如下：

s1 := fmt.Sprint("枯藤")
name := "枯藤"
age := 18
s2 := fmt.Sprintf("name:%s,age:%d", name, age)
s3 := fmt.Sprintln("枯藤")
fmt.Println(s1, s2, s3)
Errorf
Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

func Errorf(format string, a ...interface{}) error
通常使用这种方式来自定义错误类型，例如：

err := fmt.Errorf("这是一个错误")
1.1.2. 格式化占位符
*printf系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

通用占位符
占位符	说明
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号
示例代码如下：

fmt.Printf("%v\n", 100)
fmt.Printf("%v\n", false)
o := struct{ name string }{"枯藤"}
fmt.Printf("%v\n", o)
fmt.Printf("%#v\n", o)
fmt.Printf("%T\n", o)
fmt.Printf("100%%\n")
输出结果如下：

100
false
{枯藤}
struct { name string }{name:"枯藤"}
struct { name string }
100%
布尔型
占位符	说明
%t	true或false
整型
占位符	说明
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于”U+%04X”
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
示例代码如下：

n := 65
fmt.Printf("%b\n", n)
fmt.Printf("%c\n", n)
fmt.Printf("%d\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n)
fmt.Printf("%X\n", n)
输出结果如下：

    1000001
    A
    65
    101
    41
    41
浮点数与复数
占位符	说明
%b	无小数部分、二进制指数的科学计数法，如-123456p-78
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
示例代码如下：

f := 12.34
fmt.Printf("%b\n", f)
fmt.Printf("%e\n", f)
fmt.Printf("%E\n", f)
fmt.Printf("%f\n", f)
fmt.Printf("%g\n", f)
fmt.Printf("%G\n", f)
输出结果如下：

    6946802425218990p-49
    1.234000e+01
    1.234000E+01
    12.340000
    12.34
    12.34
字符串和[]byte
占位符	说明
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f
%X	每个字节用两字符十六进制数表示（使用A-F）
示例代码如下：

    s := "枯藤"
    fmt.Printf("%s\n", s)
    fmt.Printf("%q\n", s)
    fmt.Printf("%x\n", s)
    fmt.Printf("%X\n", s)
输出结果如下：

    枯藤
    "枯藤"
    e69eafe897a4
    E69EAFE897A4
指针
占位符	说明
%p	表示为十六进制，并加上前导的0x
示例代码如下：

a := 18
fmt.Printf("%p\n", &a)
fmt.Printf("%#p\n", &a)
输出结果如下：

    0xc000054058
    c000054058
宽度标识符
宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下

占位符	说明
%f	默认宽度，默认精度
%9f	宽度9，默认精度
%.2f	默认宽度，精度2
%9.2f	宽度9，精度2
%9.f	宽度9，精度0
示例代码如下：

n := 88.88
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
输出结果如下：

    88.880000
    88.880000
    88.88
        88.88
           89
其他falg
占位符	说明
’+’	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
’ ‘	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
’-’	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
’#’	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
‘0’	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
举个例子：

s := "枯藤"
fmt.Printf("%s\n", s)
fmt.Printf("%5s\n", s)
fmt.Printf("%-5s\n", s)
fmt.Printf("%5.7s\n", s)
fmt.Printf("%-5.7s\n", s)
fmt.Printf("%5.2s\n", s)
fmt.Printf("%05s\n", s)
输出结果如下：

    枯藤
       枯藤
    枯藤
       枯藤
    枯藤
       枯藤
    000枯藤
1.1.3. 获取输入
Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

fmt.Scan
函数定签名如下：

func Scan(a ...interface{}) (n int, err error)
Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
具体代码示例如下：

func main() {
    var (
        name    string
        age     int
        married bool
    )
    fmt.Scan(&name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
将上面的代码编译后在终端执行，在终端依次输入枯藤、18和false使用空格分隔。

    $ ./scan_demo 
    枯藤 18 false
    扫描结果 name:枯藤 age:18 married:false
fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。

fmt.Scanf
函数签名如下：

func Scanf(format string, a ...interface{}) (n int, err error)
Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
本函数返回成功扫描的数据个数和遇到的任何错误。
代码示例如下：

func main() {
    var (
        name    string
        age     int
        married bool
    )
    fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
将上面的代码编译后在终端执行，在终端按照指定的格式依次输入枯藤、18和false。

    $ ./scan_demo 
    1:枯藤 2:18 3:false
    扫描结果 name:枯藤 age:18 married:false
fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。

例如，我们还是按照上个示例中以空格分隔的方式输入，fmt.Scanf就不能正确扫描到输入的数据。

    $ ./scan_demo 
    枯藤 18 false
    扫描结果 name: age:0 married:false
fmt.Scanln
函数签名如下：

func Scanln(a ...interface{}) (n int, err error)
Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
本函数返回成功扫描的数据个数和遇到的任何错误。
具体代码示例如下：

    func main() {
        var (
            name    string
            age     int
            married bool
        )
        fmt.Scanln(&name, &age, &married)
        fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
    }
将上面的代码编译后在终端执行，在终端依次输入枯藤、18和false使用空格分隔。

    $ ./scan_demo 
    枯藤 18 false
    扫描结果 name:枯藤 age:18 married:false
fmt.Scanln遇到回车就结束扫描了，这个比较常用。

bufio.NewReader
有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。示例代码如下：

func bufioDemo() {
    reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
    fmt.Print("请输入内容：")
    text, _ := reader.ReadString('\n') // 读到换行
    text = strings.TrimSpace(text)
    fmt.Printf("%#v\n", text)
}
Fscan系列
这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。

func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
Sscan系列
这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。

func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
赏