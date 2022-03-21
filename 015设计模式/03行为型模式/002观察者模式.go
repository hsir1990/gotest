观察者模式
观察者模式用于触发联动。

一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。

// subject
// 英 [ˈsʌbdʒɪkt , səbˈdʒekt]   美 [ˈsʌbdʒɪkt , səbˈdʒekt]  
// n.
// 主题;题目;话题;题材;问题;学科;科目;课程;表现对象;绘画(或拍摄)题材;接受试验者;主语;国民，臣民
// adj.
// 可能受…影响的;易遭受…的;取决于;视…而定;受…支配;服从于;受异族统治的
// vt.
// 使臣服;使顺从;(尤指)压服


// attach
// 英 [əˈtætʃ]   美 [əˈtætʃ]  
// v.
// 贴上;重视;把…固定，把…附(在…上);认为有重要性(或意义、价值、分量等);(有时不受欢迎或未受邀请而)参加，和…在一起，缠着;（使）与…有联系


// observer
// 英 [əbˈzɜːvə(r)]   美 [əbˈzɜːrvər]  
// 观察者;观察员;观察家;观测者;目击者;旁听者;评论员


// notify
// 英 [ˈnəʊtɪfaɪ]   美 [ˈnoʊtɪfaɪ]  
// vt.
// 通知;(正式)通报


obserser.go
package observer

import "fmt"

type Subject struct {
    observers []Observer
    context   string
}

func NewSubject() *Subject {
    return &Subject{
        observers: make([]Observer, 0),
    }
}

func (s *Subject) Attach(o Observer) {
    s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
    for _, o := range s.observers {
        o.Update(s)
    }
}

func (s *Subject) UpdateContext(context string) {
    s.context = context
    s.notify()
}

type Observer interface {
    Update(*Subject)
}

type Reader struct {
    name string
}

func NewReader(name string) *Reader {
    return &Reader{
        name: name,
    }
}

func (r *Reader) Update(s *Subject) {
    fmt.Printf("%s receive %s\n", r.name, s.context)
}
obserser_test.go
package observer

func ExampleObserver() {
    subject := NewSubject()
    reader1 := NewReader("reader1")
    reader2 := NewReader("reader2")
    reader3 := NewReader("reader3")
    subject.Attach(reader1)
    subject.Attach(reader2)
    subject.Attach(reader3)

    subject.UpdateContext("observer mode")
    // Output:
    // reader1 receive observer mode
    // reader2 receive observer mode
    // reader3 receive observer mode
}
文