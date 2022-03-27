//命令模式，比如你的leader让你开发一个需求
//一个人接到命令，然后后面的人接着迭代执行，最后一个命令为空，就不用执行了 execute
package main

import (
	"fmt"
)

type Person struct {
	name string
	cmd  Command
}
type Command struct {
	person *Person
	method func()
}

func NewCommand(p *Person, method func()) Command {
	return Command{
		person: p,
		method: method,
	}
}

func (c *Command) Execute() {
	c.method()
}

func NewPerson(name string, cmd Command) Person {
	return Person{
		name: name,
		cmd:  cmd,
	}
}

func (p *Person) Buy() {
	fmt.Println(fmt.Sprintf("%s is buying", p.name))
	p.cmd.Execute()
}

func (p *Person) Cook() {
	fmt.Println(fmt.Sprintf("%s is cooking", p.name))
	p.cmd.Execute()
}

func (p *Person) Wash() {
	fmt.Println(fmt.Sprintf("%s is washing", p.name))
	p.cmd.Execute()
}

func (p *Person) Talk() {
	fmt.Println(fmt.Sprintf("%s is talking", p.name))
	p.cmd.Execute()
}

func (p *Person) Listen() {
	fmt.Println(fmt.Sprintf("%s is listening", p.name))
}

func main() {
	laowang := NewPerson("wang", NewCommand(nil, nil))
	laozhang := NewPerson("zhang", NewCommand(&laowang, laowang.Listen))
	laofeng := NewPerson("feng", NewCommand(&laozhang, laozhang.Buy))
	laoding := NewPerson("feng", NewCommand(&laofeng, laofeng.Cook))
	laoli := NewPerson("feng", NewCommand(&laoding, laoding.Wash))

	laoli.Talk()
	// feng is Talking
	// feng is washing
	// feng is cooking
	// zhang is buying
	// wang is Listening

}

// command
// 英 [kəˈmɑːnd]   美 [kəˈmænd]
// n.
// (给人或动物的)命令;指令;命令;控制;管辖;指挥;兵团，军区，指挥部，司令部;知识
// v.
// 命令;指挥，统率(陆军、海军等);应得;博得;值得;居高临下地掌控;控制

// execute
// 英 [ˈeksɪkjuːt]   美 [ˈeksɪkjuːt]
// vt.
// (尤指依法)处决，处死;实行;执行;实施;成功地完成(技巧或动作);制作，做成（艺术品）;执行（法令）


// behavioral
// 英 [bɪ'heɪvjərəl]   美 [bɪ'heɪvjərəl]  
// adj.
// 关于行为的