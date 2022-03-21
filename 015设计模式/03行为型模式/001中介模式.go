中介者模式
中介者模式封装对象之间互交，使依赖变的简单，并且使复杂互交简单化，封装在中介者中。

例子中的中介者使用单例模式生成中介者。

中介者的change使用switch判断类型。


// driver
// 英 [ˈdraɪvə(r)]   美 [ˈdraɪvər]  
// n.
// 驾驶员;司机;驾车者;球杆;驱动程序;驱动因素

// card
// 英 [kɑːd]   美 [kɑːrd]  
// n.
// (尤指显示个人信息的)卡片;信用卡;厚纸片;薄纸板;现金卡;储值卡;贺卡;纸牌游戏;怪人;（赛马大会的）赛事一览表;梳理机
// vt.
// (用钢丝刷)梳理;要求出示身份证(以确认年龄，如购酒)

// Mediator
// 英 [ˈmiːdieɪtə(r)]   美 [ˈmiːdieɪtər]  
// 调停者;中介;中介变量;调解者;中介模式


mediator.go
package mediator

import (
    "fmt"
    "strings"
)

type CDDriver struct {
    Data string
}

func (c *CDDriver) ReadData() {
    c.Data = "music,image"

    fmt.Printf("CDDriver: reading data %s\n", c.Data)
    GetMediatorInstance().changed(c)
}

type CPU struct {
    Video string
    Sound string
}

func (c *CPU) Process(data string) {
    sp := strings.Split(data, ",")
    c.Sound = sp[0]
    c.Video = sp[1]

    fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
    GetMediatorInstance().changed(c)
}

type VideoCard struct {
    Data string
}

func (v *VideoCard) Display(data string) {
    v.Data = data
    fmt.Printf("VideoCard: display %s\n", v.Data)
    GetMediatorInstance().changed(v)
}

type SoundCard struct {
    Data string
}

func (s *SoundCard) Play(data string) {
    s.Data = data
    fmt.Printf("SoundCard: play %s\n", s.Data)
    GetMediatorInstance().changed(s)
}

type Mediator struct {
    CD    *CDDriver
    CPU   *CPU
    Video *VideoCard
    Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
    if mediator == nil {
        mediator = &Mediator{}
    }
    return mediator
}

func (m *Mediator) changed(i interface{}) {
    switch inst := i.(type) {
    case *CDDriver:
        m.CPU.Process(inst.Data)
    case *CPU:
        m.Sound.Play(inst.Sound)
        m.Video.Display(inst.Video)
    }
}
mediator_test.go
package mediator

import "testing"

func TestMediator(t *testing.T) {
    mediator := GetMediatorInstance()
    mediator.CD = &CDDriver{}
    mediator.CPU = &CPU{}
    mediator.Video = &VideoCard{}
    mediator.Sound = &SoundCard{}

    //Tiggle
    mediator.CD.ReadData()

    if mediator.CD.Data != "music,image" {
        t.Fatalf("CD unexpect data %s", mediator.CD.Data)
    }

    if mediator.CPU.Sound != "music" {
        t.Fatalf("CPU unexpect sound data %s", mediator.CPU.Sound)
    }

    if mediator.CPU.Video != "image" {
        t.Fatalf("CPU unexpect video data %s", mediator.CPU.Video)
    }

    if mediator.Video.Data != "image" {
        t.Fatalf("VidoeCard unexpect data %s", mediator.Video.Data)
    }

    if mediator.Sound.Data != "music" {
        t.Fatalf("SoundCard unexpect data %s", mediator.Sound.Data)
    }
}
