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

		qp.SetRenderHint(gui.QPainter__Antialiasing, true)
		qp.SetFont(gui.NewQFont2("Times", 100, 0, true))

		rect := core.NewQRect4(20, 20, 300, 200)
		qp.DrawRect3(rect)

		// 居中绘制文本
		metrics := qp.FontMetrics()

		stringHeight := metrics.Ascent() + metrics.Descent() // 不算 line gap
		stringWidth := metrics.AverageCharWidth()            // 字符串的宽度
		x := rect.X() + (rect.Width()-stringWidth)/2
		y := rect.Y() + (rect.Height()-stringHeight)/2 + metrics.Ascent()
		qp.DrawText3(x, y, "jEh")

		// 绘制字符串的包围矩形
		y = rect.Y() + (rect.Height()-stringHeight)/2
		//qp.SetPen(core.Qt__lightGray)
		qp.DrawRect2(x, y, stringWidth, stringHeight)

		// Menlo 后就可以看到确实是完全居中的，说明 QFontMetrics 得到的字体信息没问题，
		// 只是有的字体为了美观漂亮作了一些调整，对于这些字体如果要完全的居中效果的话，
		// 只好在使用上面的计算方式后再手动的微调一下就好了。
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
