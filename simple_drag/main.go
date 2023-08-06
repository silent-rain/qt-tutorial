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
	widget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(widget)

	//QLineEdit内置了对drag(拖动)操作的支持。
	edit := widgets.NewQLineEdit(widget)
	// 打开拖放
	edit.SetDragEnabled(true)
	edit.Move2(30, 65)
	edit.ConnectDropEvent(func(event *gui.QDropEvent) {
		fmt.Println(event.MimeData().Text())
		edit.SetText(event.MimeData().Text())
	})

	// 我们需要重新实现某些方法才能使QPushButton接受拖放操作。
	button := widgets.NewQPushButton2("Button", widget)
	button.Move2(190, 65)
	//	使该控件接受drop(放下)事件。
	button.SetAcceptDrops(true)
	button.ConnectDragEnterEvent(dragEnterEvent)
	// 通过重新实现dropEvent()方法，我们定义了在drop事件发生时的行为。
	//这里我们改变了按钮的文字。
	button.ConnectDropEvent(func(event *gui.QDropEvent) {
		fmt.Println(event.MimeData().Text())
		button.SetText(event.MimeData().Text())
	})
}

// 首先我们重新实现了dragEnterEvent()方法，并设置可接受的数据类型(在这里是普通文本)。
func dragEnterEvent(event *gui.QDragEnterEvent) {
	if event.MimeData().HasFormat("text/plain") {
		event.AcceptProposedAction()
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
