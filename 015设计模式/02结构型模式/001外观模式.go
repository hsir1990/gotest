外观模式
API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。

facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。
// facade
// 英 [fəˈsɑːd]   美 [fəˈsɑːd]  
// n.
// 外观;(建筑物的)正面，立面;(虚假的)表面，外表

facade.go
package facade

import "fmt"

func NewAPI() API {
    return &apiImpl{
        a: NewAModuleAPI(),
        b: NewBModuleAPI(),
    }
}

// API是facade包的facade接口
//API is facade interface of facade package
type API interface {
    Test() string
}

//facade implement
type apiImpl struct {
    a AModuleAPI
    b BModuleAPI
}

func (a *apiImpl) Test() string {
    aRet := a.a.TestA()
    bRet := a.b.TestB()
    return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
    return &aModuleImpl{}
}

//AModuleAPI ...
type AModuleAPI interface {
    TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
    return "A module running"
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
    return &bModuleImpl{}
}

//BModuleAPI ...
type BModuleAPI interface {
    TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
    return "B module running"
}
facade_test.go
package facade

import "testing"

var expect = "A module running\nB module running"

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
    api := NewAPI()
    ret := api.Test()
    if ret != expect {
        t.Fatalf("expect %s, return %s", expect, ret)
    }
}
