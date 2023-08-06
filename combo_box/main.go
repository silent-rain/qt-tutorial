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
	//widget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	widget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(widget)

	// 标签
	lbl := widgets.NewQLabel2("Ubuntu", widget, 0)
	lbl.Move2(50, 150)

	// 下拉列表
	combo := widgets.NewQComboBox(widget)
	combo.Move2(50, 50)
	combo.AddItem("Ubuntu", core.NewQVariant15("Ubuntu"))
	combo.AddItem("Mandriva", core.NewQVariant15("Mandriva"))
	combo.AddItem("Fedora", core.NewQVariant15("Fedora"))
	combo.AddItem("Arch", core.NewQVariant15("Arch"))
	combo.AddItem("Gentoo", core.NewQVariant15("Gentoo"))

	// 当选中某个条目时会调用方法。
	combo.ConnectActivated2(func(text string) {
		// 在方法中我们将QLabel控件的内容设置为选中的条目，然后调整它的尺寸。
		lbl.SetText(text)
		// 滚动长度
		lbl.AdjustSize()
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
