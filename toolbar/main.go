package main

import (
	"os"

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

	// 状态栏
	//timeout: 显示时间, -1: 一直显示
	app.StatusBar().ShowMessage("Ready", -1)
	// app.StatusBar().ShowMessage("Ready", 5)

	return app
}

// InitComponents 初始化组件
func InitComponents(app *widgets.QMainWindow) {
	// 子按钮
	// exitAction := widgets.NewQAction3(gui.NewQIcon5("images/app.ico"), "&Exit", app)
	// // 快捷键，自定义
	// exitAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
	// // 事件触发
	// exitAction.ConnectTriggered(func(checked bool) {
	// 	app.Close()
	// })

	// var actions []*widgets.QAction
	// actions = append(actions, exitAction)

	// 创建工具栏
	// toolbar := widgets.NewQToolBar("Exit", app)
	toolbar := app.AddToolBar3("Tools")

	// 子按钮
	exitAction := toolbar.AddAction2(gui.NewQIcon5("images/app.ico"), "&Exit")
	// 快捷键，自定义
	exitAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
	// 事件触发
	exitAction.ConnectTriggered(func(checked bool) {
		app.Close()
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
