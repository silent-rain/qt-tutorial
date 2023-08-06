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

	// 表格布局
	grid := widgets.NewQGridLayout(layoutWidget)
	grid.SetContentsMargins(0, 0, 0, 0)

	names := []string{
		"Cls", "Bck", "", "Close",
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", ".", "=", "+",
	}
	// 我们创建一个网格中的位置的列表
	var positions [20][2]int
	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			item := [2]int{i, j}
			positions[k] = item
			k += 1
		}
	}

	// 创建按钮并使用 AddWidget()方法添加到布局中。
	for index := range positions {
		button := widgets.NewQPushButton2(names[index], layoutWidget)
		value := positions[index]
		grid.AddWidget2(button, value[0], value[1], core.Qt__AlignLeft)
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
