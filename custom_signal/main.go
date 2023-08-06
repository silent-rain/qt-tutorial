package main

import (
	"fmt"
	"os"

	"gosignal"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// 控制窗口显示在屏幕中心的方法
func center(app *widgets.QMainWindow) {
	// 获得窗口
	qr := app.FrameGeometry()

	//  获得屏幕中心点
	cp := widgets.NewQDesktopWidget().AvailableGeometry2(app).Center()
	// 显示到屏幕中心
	qr.MoveCenter(cp)
	app.Move(qr.TopLeft())
}

// InitMainWindow 初始化主窗口
func InitMainWindow() *widgets.QMainWindow {
	// 创建窗口
	app := widgets.NewQMainWindow(nil, 0)

	// 设置窗口的标题
	app.SetWindowTitle("Qt 教程")

	// 设置窗口的位置和大小
	app.SetGeometry2(300, 300, 300, 220)

	// 设置窗口的图标，引用当前目录下的web.png图片
	app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	return app
}

// InitComponents 初始化组件
func InitComponents(app *widgets.QMainWindow) {
	// 布局窗口组件载体
	widget := widgets.NewQWidget(app, core.Qt__Widget)
	widget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(widget)

	lbl := widgets.NewQLabel(widget, core.Qt__Widget)
	lbl.SetText("默认标签")
	lbl.Move2(10, 20)

	btn1 := widgets.NewQPushButton2("Modify Label", widget)
	btn1.SetObjectName("Button1")
	btn1.SetCheckable(true)
	btn1.Move2(30, 50)

	modifyLabelSignal := gosignal.NewGoSignal(nil)
	modifyLabelSignal.ConnectGoSignal(func(str string) {
		lbl.SetText(str)
	})

	btn1.ConnectClicked(func(checked bool) {
		fmt.Println("btn1 checked: ", checked)
		text := "默认标签"
		if checked {
			text = "信号标签"
		}
		modifyLabelSignal.Emit2(text)
	})

	btn2 := widgets.NewQPushButton2("Close App", widget)
	btn2.SetObjectName("Button2")
	btn2.Move2(150, 50)

	closeAppSignal := gosignal.NewGoSignal(nil)
	closeAppSignal.ConnectGoSignal(func() {
		app.Close()
	})

	btn2.ConnectClicked(func(checked bool) {
		closeAppSignal.Emit()
	})
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化主窗口
	app := InitMainWindow()

	// 初始化组件
	InitComponents(app)

	// 控制窗口显示在屏幕中心的方法
	center(app)

	// 显示主窗口
	app.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
