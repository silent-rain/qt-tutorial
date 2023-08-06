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
	layoutWidget := widgets.NewQWidget(app, core.Qt__Widget)
	layoutWidget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(layoutWidget)

	// 纵向布局
	okVButton := widgets.NewQPushButton2("OK-V", layoutWidget)
	cancelVButton := widgets.NewQPushButton2("Cancel-V", layoutWidget)
	vbox := widgets.NewQVBoxLayout2(layoutWidget)
	vbox.SetObjectName("vbox")
	vbox.SetContentsMargins(20, 20, 0, 0)
	vbox.AddWidget(okVButton, 0, core.Qt__AlignTop)
	vbox.AddWidget(cancelVButton, 0, core.Qt__AlignTop)

	// 横向布局
	okButton := widgets.NewQPushButton2("OK", layoutWidget)
	cancelButton := widgets.NewQPushButton2("Cancel", layoutWidget)
	hbox := widgets.NewQHBoxLayout()
	hbox.SetObjectName("vbox")
	hbox.SetContentsMargins(0, 0, 0, 0)
	hbox.AddStretch(1)
	hbox.AddWidget(okButton, 0, core.Qt__AlignRight)
	hbox.AddWidget(cancelButton, 0, core.Qt__AlignRight)

	// 横向布局中添加纵向布局
	vbox.AddStretch(1)
	vbox.AddLayout(hbox, 0)
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
