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

	// 进度条
	pbar := widgets.NewQProgressBar(widget)
	pbar.SetGeometry2(30, 40, 200, 25)
	pbar.SetValue(0)

	btn := widgets.NewQPushButton2("Start", widget)
	btn.Move2(40, 80)

	// 计数
	step := 0

	// 定时器,我们使用定时器timer来激活QProgressBar
	// timer := core.NewQBasicTimer()
	timer := core.NewQTimer(widget)
	// 每个QObject及其子类都有一个timerEvent()事件处理器。
	//我们要重新实现这个事件处理器来响应定时器事件。
	timer.ConnectEvent(func(e *core.QEvent) bool {
		if step >= 100 {
			timer.Stop()
			btn.SetText("Finished")
			return true
		} else {
			step = step + 1
			pbar.SetValue(step)
		}
		return false
	})

	//方法中启动/停止定时器。
	btn.ConnectClicked(func(checked bool) {
		if timer.IsActive() {
			timer.Stop()
			btn.SetText("Start")
		} else {
			//我们调用start()方法启动一个计时器。这个方法有两个参数:超时和对象将接收的事件。
			timer.Start(100)
			btn.SetText("Stop")
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
