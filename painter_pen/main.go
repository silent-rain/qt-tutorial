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

		//我们创建一个QPen对象。颜色是黑色的。
		//宽度设置为2像素,这样我们可以看到笔风格之间的差异。Qt.SolidLine是预定义的钢笔样式。
		//pen := gui.NewQPen3(gui.NewQColor2(core.Qt__black))
		pen := gui.NewQPen4(gui.NewQBrush2(core.Qt__SolidPattern), 2, core.Qt__SolidLine, core.Qt__FlatCap, core.Qt__MiterJoin)
		qp.SetPen(pen)
		qp.DrawLine3(20, 40, 250, 40)

		pen.SetStyle(core.Qt__DashDotDotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 80, 250, 80)

		pen.SetStyle(core.Qt__DashDotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 120, 250, 120)

		pen.SetStyle(core.Qt__DotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 160, 250, 160)

		pen.SetStyle(core.Qt__DashDotDotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 200, 250, 200)

		//这里我们定义了一个画笔风格。
		//我们设置了Qt.CustomDashLine并调用了setDashPattern()方法，
		//它的参数(一个数字列表)定义了一种风格，必须有偶数个数字；
		//其中奇数表示绘制实线，偶数表示留空。数值越大，直线或空白就越大。
		//这里我们定义了1像素的实线，4像素的空白，5像素实线，4像素空白。。。
		pen.SetStyle(core.Qt__CustomDashLine)
		qp.SetPen(pen)
		// pen.SetDashPattern([]float64{1, 4, 5, 4})
		qp.DrawLine3(20, 240, 250, 240)

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
