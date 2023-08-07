package main

import (
	"os"

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

	textEdit := widgets.NewQTextEdit(widget)
	textEdit.Resize2(300, 200)

	// 退出操作
	exitAction := widgets.NewQAction3(gui.NewQIcon5("images/app.ico"), "&Exit", nil)
	// 快捷键，自定义
	exitAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
	// 提示语
	exitAction.SetStatusTip("Exit application")
	// 事件触发
	exitAction.ConnectTriggered(func(checked bool) {
		app.Close()
	})

	actions := widgets.NewQActionGroup(app)
	actions.AddAction(exitAction)

	//创建一个菜单栏
	// menubar := widgets.NewQMenuBar(app)
	menubar := app.MenuBar()
	// 添加菜单
	fileMenu := menubar.AddMenu2("&File")
	// 添加按钮
	fileMenu.AddActions(actions.Actions())

	// 创建工具栏
	// toolbar := widgets.NewQToolBar("Exit", app)
	toolbar := app.AddToolBar3("Exit")
	toolbar.AddActions(actions.Actions())
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
