原型模式
原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。

原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。

prototype.go
package prototype

//Cloneable 是原型对象需要实现的接口
type Cloneable interface {
    Clone() Cloneable
}

type PrototypeManager struct {
    prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
    return &PrototypeManager{
        prototypes: make(map[string]Cloneable),
    }
}

func (p *PrototypeManager) Get(name string) Cloneable {
    return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
    p.prototypes[name] = prototype
}



prototype_test.go
package prototype

import "testing"

var manager *PrototypeManager

type Type1 struct {
    name string
}

func (t *Type1) Clone() Cloneable {
    tc := *t
    return &tc
}

type Type2 struct {
    name string
}

func (t *Type2) Clone() Cloneable {
    tc := *t
    return &tc
}

func init() {
    manager = NewPrototypeManager()

    t1 := &Type1{
        name: "type1",
    }
    manager.Set("t1", t1)
}


func TestClone(t *testing.T) {
    t1 := manager.Get("t1")

    t2 := t1.Clone()

    if t1 == t2 {
        t.Fatal("error! get clone not working")
    }
}

func TestCloneFromManager(t *testing.T) {
    c := manager.Get("t1").Clone()

    t1 := c.(*Type1)
    if t1.name != "type1" {
        t.Fatal("error")
    }

}

