package main

import (
	"fmt"
	"os"

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

// CloseEvent 关闭事件
func CloseEvent(event *gui.QCloseEvent) {
	reply := widgets.QMessageBox_Question(nil,
		"Message",
		"Are you sure to quit?",
		widgets.QMessageBox__Yes|widgets.QMessageBox__No,
		widgets.QMessageBox__Yes)

	if reply == widgets.QMessageBox__Yes {
		fmt.Println("退出")
		event.Accept()
	} else {
		fmt.Println("取消")
		event.Ignore()
	}
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化主窗口
	app := InitMainWindow()

	// 控制窗口显示在屏幕中心的方法
	center(app)

	// 修改默认关闭窗口事件
	app.ConnectCloseEvent(CloseEvent)

	// 显示主窗口
	app.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
