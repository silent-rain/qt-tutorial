package main

import (
	"fmt"
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

	// 初始化QFrame的颜色为黑色
	color := gui.NewQColor3(0, 0, 0, 0)

	frm := widgets.NewQFrame(widget, 0)
	frm.SetGeometry2(130, 22, 100, 100)
	frm.SetStyleSheet("QWidget { background-color: " + color.Name() + " }")

	btn := widgets.NewQPushButton2("Dialog", widget)
	btn.Move2(20, 20)
	btn.ConnectClicked(func(checked bool) {
		colorDialog := widgets.NewQColorDialog(widget).GetColor(
			color,
			widget,
			"调色板",
			widgets.QColorDialog__ShowAlphaChannel)

		// 我们要先检查col的值。如果点击的是Cancel按钮，返回的颜色值是无效的。
		//当颜色值有效时，我们通过样式表(style sheet)来改变背景颜色。
		if colorDialog.IsValid() {
			fmt.Println("调色： ", colorDialog.Name())
			frm.SetStyleSheet("QWidget { background-color: " + colorDialog.Name() + " }")
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
