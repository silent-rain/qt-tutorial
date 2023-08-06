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

	lbl := widgets.NewQLabel2("Knowledge only matters", widget, 0)
	lbl.Move2(130, 20)

	btn := widgets.NewQPushButton2("Dialog", widget)
	// btn.SetSizePolicy(widgets.QSizePolicy__Fixed)
	btn.Move2(20, 20)

	vbox := widgets.NewQVBoxLayout2(widget)
	vbox.AddWidget(lbl, 0, core.Qt__AlignLeft)
	vbox.AddWidget(btn, 0, core.Qt__AlignLeft)

	btn.ConnectClicked(func(checked bool) {
		// 这一行代码弹出字体选择对话框，GetFont2()方法返回字体名称和ok参数，
		//如果用户点击了ok他就是True,否则就是false
		ok := false
		font := widgets.NewQFontDialog(widget).GetFont2(&ok, widget)
		// 如果我们点击了ok，标签的字体就会被改变
		if ok {
			lbl.SetFont(font)
		}
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
