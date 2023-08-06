package gosignal

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
)

// GoSignal GO 自定义信号
type GoSignal struct {
	*core.QObject
	innerFunc interface{}
}

// NewGoSignal 新建 GO 自定义信号对象
func NewGoSignal(parent core.QObject_ITF) *GoSignal {
	return &GoSignal{
		QObject: core.NewQObject(parent),
	}
}

// ConnectGoSignal 连接信号
func (g *GoSignal) ConnectGoSignal(f interface{}) {
	// 不存在信号则注册
	if !qt.ExistsSignal(g.Pointer(), "goSignal") {
		g.innerFunc = f
		qt.ConnectSignal(g.Pointer(), "goSignal", unsafe.Pointer(&f))
	}
}

// Emit 发送无参数信号
func (g *GoSignal) Emit() (err error) {
	if signal := qt.LendSignal(g.Pointer(), "goSignal"); signal == nil {
		return fmt.Errorf("not ExistsSignal -> goSignal")
	}
	funcValue := reflect.ValueOf(g.innerFunc)
	params := make([]reflect.Value, 0)
	go funcValue.Call(params)
	// go g.innerFunc.(func())()
	return nil
}

// Emit2 发送携带参数的自定义信号，对应响应函数也必须有参数接收
func (g *GoSignal) Emit2(args interface{}) error {
	if signal := qt.LendSignal(g.Pointer(), "goSignal"); signal == nil {
		return fmt.Errorf("not ExistsSignal -> goSignal")
	}
	funcValue := reflect.ValueOf(g.innerFunc)
	params := []reflect.Value{reflect.ValueOf(args)}
	go funcValue.Call(params)
	// go g.innerFunc.(func(args interface{}))(args)
	return nil
}
