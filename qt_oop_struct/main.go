package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type UIMainWindow struct {
	CentralWidget *widgets.QWidget
	Menubar       *widgets.QMenuBar
	Statusbar     *widgets.QStatusBar
	PushButton1   *widgets.QPushButton
	PushButton2   *widgets.QPushButton
}

// SetupUI 初始化UI
func (w *UIMainWindow) SetupUI(mainWindow *widgets.QMainWindow) {
	mainWindow.SetObjectName("MainWindow")
	// 设置窗口的标题
	mainWindow.SetWindowTitle("Qt 教程")
	// 设置窗口的位置和大小
	mainWindow.SetGeometry2(300, 300, 300, 220)
	// 设置窗口的图标
	mainWindow.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	// 控制窗口显示在屏幕中心的方法
	w.Center(mainWindow)

	// 中心窗口组件载体
	w.CentralWidget = widgets.NewQWidget(mainWindow, core.Qt__Widget)
	// w.CentralWidget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	w.CentralWidget.SetGeometry2(0, 0, 300, 220)
	mainWindow.SetCentralWidget(w.CentralWidget)

	// 状态栏
	w.Statusbar = widgets.NewQStatusBar(mainWindow)
	w.Statusbar.SetObjectName("Statusbar")
	mainWindow.SetStatusBar(w.Statusbar)

	// 按钮
	w.PushButton1 = widgets.NewQPushButton2("PushButton1", w.CentralWidget)
	w.PushButton1.Move2(30, 50)
	w.PushButton1.ConnectClicked(w.ButtonClicked)
}

// 控制窗口显示在屏幕中心的方法
func (w *UIMainWindow) Center(mainWindow *widgets.QMainWindow) {
	// 获得窗口
	qr := mainWindow.FrameGeometry()

	//  获得屏幕中心点
	cp := widgets.NewQDesktopWidget().AvailableGeometry2(mainWindow).Center()
	// 显示到屏幕中心
	qr.MoveCenter(cp)
	mainWindow.Move(qr.TopLeft())
}

// ButtonClicked 按钮点击事件
func (w *UIMainWindow) ButtonClicked(checked bool) {
	fmt.Println("ButtonClicked", checked)
	w.Statusbar.ShowMessage("sender"+" was pressed", 0)
}

// RetranslateUi 重置UI
func (w *UIMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "QT 教程", "", -1))
	w.PushButton1.SetText(_translate("MainWindow", "PushButton", "", -1))
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化窗口
	mainWindow := widgets.NewQMainWindow(nil, 0)

	uiMain := UIMainWindow{}
	uiMain.SetupUI(mainWindow)
	uiMain.RetranslateUi(mainWindow)

	// 显示组件
	mainWindow.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
