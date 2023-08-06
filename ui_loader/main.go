package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
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

// LoadUI 加载UI文件
func LoadUI(app *widgets.QMainWindow) *widgets.QMainWindow {
	uiFile := core.NewQFile2("./untitled.ui")
	uiFile.Open(core.QIODevice__ReadOnly)
	defer uiFile.Close()
	fmt.Println("========= ", uiFile.FileName())
	uitools.NewQUiLoader(app)
	// loader := uitools.NewQUiLoader(app)
	// widgetWindow := loader.Load(uiFile, app)
	// mainWindow := (*widgets.QMainWindow)(unsafe.Pointer(widgetWindow))
	return nil
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

	// 标签
	labelObj := app.FindChild("label", core.Qt__FindChildrenRecursively)
	label := (*widgets.QLabel)(unsafe.Pointer(labelObj))

	// 默认按钮
	pushButtonObj := app.FindChild("pushButton", core.Qt__FindChildrenRecursively)
	pushButton := (*widgets.QPushButton)(unsafe.Pointer(pushButtonObj))
	pushButton.ConnectClicked(func(checked bool) {
		fmt.Println("ButtonClicked")
		label.SetText("默认按钮信息")
	})

	// 添加按钮
	btn2 := widgets.NewQPushButton2("code按钮", app)
	btn2.Move2(30, 50)
	btn2.ConnectClicked(func(checked bool) {
		fmt.Println("code按钮")
		label.SetText("code按钮信息")
	})
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)
	// core.NewQCoreApplication(len(os.Args), os.Args).SetAttribute(core.Qt__AA_ShareOpenGLContexts, false)

	// 初始化主窗口
	app := InitMainWindow()
	// LoadUI 加载UI文件
	LoadUI(app)

	// 初始化组件
	InitComponents(app)

	// 控制窗口显示在屏幕中心的方法
	center(app)

	// 显示主窗口
	app.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
