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

	// 这是初始黑颜色的值。
	color := gui.NewQColor3(0, 0, 0, 0)

	frm := widgets.NewQFrame(widget, 0)
	frm.SetGeometry2(150, 20, 100, 100)
	frm.SetStyleSheet("QWidget { background-color: " + color.Name() + " }")

	// 我们创建一个QPushButton并通过其SetCheckable()方法来得到一个ToggleButton。
	redb := widgets.NewQPushButton2("Red", widget)
	redb.SetCheckable(true)
	redb.Move2(10, 10)

	greenb := widgets.NewQPushButton2("Green", widget)
	greenb.SetCheckable(true)
	greenb.Move2(10, 60)

	blueb := widgets.NewQPushButton2("Blue", widget)
	blueb.SetCheckable(true)
	blueb.Move2(10, 110)

	// 将clicked信号连接到用户自定义的方法。我们通过clicked信号操作一个布尔值。
	redb.ConnectClicked(setColor(redb, color, frm))
	greenb.ConnectClicked(setColor(greenb, color, frm))
	blueb.ConnectClicked(setColor(blueb, color, frm))
}

func setColor(btn *widgets.QPushButton, color *gui.QColor, frm *widgets.QFrame) func(checked bool) {
	return func(checked bool) {
		val := 0
		if checked {
			val = 255
		}

		if btn.Text() == "Red" {
			color.SetRed(val)
		} else if btn.Text() == "Green" {
			color.SetGreen(val)
		} else {
			color.SetBlue(val)
		}
		frm.SetStyleSheet("QWidget { background-color: " + color.Name() + " }")
	}
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
