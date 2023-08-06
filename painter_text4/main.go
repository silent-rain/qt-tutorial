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

	widget.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		qp := gui.NewQPainter2(widget)
		//qp.Begin(widget)

		qp.SetRenderHint(gui.QPainter__Antialiasing, true)

		text := "QPainter provides highly optimized functions to do " +
			"most of the drawing GUI programs require. It can draw " +
			"everything from simple lines to complex shapes " +
			"like pies and chords. " +
			"看看是否也支持中文呢，如果不支持那就悲剧了！"
		rect := core.NewQRect4(20, 20, 300, 200)
		qp.DrawRect3(rect) // 画出矩形，可以看到超出此矩形的部分文本不可见

		//core.Qt__TextWordWrap | core.Qt__AlignRight
		qp.DrawText4(rect, int(core.Qt__TextWordWrap), text, rect)
		qp.End()

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
