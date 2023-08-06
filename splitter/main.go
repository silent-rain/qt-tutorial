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

	hbox := widgets.NewQHBoxLayout2(widget)

	topleft := widgets.NewQFrame(widget, core.Qt__Widget)
	topleft.SetFrameShape(widgets.QFrame__StyledPanel)
	col := gui.NewQColor3(0, 0, 0, 0)
	topleft.SetStyleSheet("QWidget { background-color: " + col.Name() + " }")

	topright := widgets.NewQFrame(widget, core.Qt__Widget)
	topright.SetFrameShape(widgets.QFrame__StyledPanel)
	col = gui.NewQColor3(100, 100, 100, 100)
	topright.SetStyleSheet("QWidget { background-color: " + col.Name() + " }")

	bottom := widgets.NewQFrame(widget, core.Qt__Widget)
	bottom.SetFrameShape(widgets.QFrame__StyledPanel)
	col = gui.NewQColor3(200, 200, 200, 200)
	bottom.SetStyleSheet("QWidget { background-color: " + col.Name() + " }")

	splitter1 := widgets.NewQSplitter2(core.Qt__Horizontal, widget)
	splitter1.AddWidget(topleft)
	splitter1.AddWidget(topright)
	splitter1.SetMinimumHeight(50)

	splitter2 := widgets.NewQSplitter2(core.Qt__Vertical, widget)
	splitter2.AddWidget(splitter1)
	splitter2.AddWidget(bottom)

	hbox.AddWidget(splitter2, 0, core.Qt__AlignLeft)
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
