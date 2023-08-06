package main

import (
	"fmt"
	"os"

	"gosignal"

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
	// 连接信号
	signal := gosignal.NewGoSignal(nil)
	signal.ConnectGoSignal(getData)

	// 鼠标事件处理器重写
	app.ConnectMoveEvent(moveEvent(signal))
}

// 移动事件
func moveEvent(signal *gosignal.GoSignal) func(event *gui.QMoveEvent) {
	return func(event *gui.QMoveEvent) {
		// 发送信号： 坐标
		// event.Pos().X(),event.Pos().Y()
		_ = signal.Emit2([]interface{}{event.Pos().X(), event.Pos().Y()})
		fmt.Println("getData: ", event.Pos().X(), event.Pos().Y())
	}
}

// 获取移动事件发送的数据
func getData(args interface{}) {
	fmt.Println("getData: ", args)
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
