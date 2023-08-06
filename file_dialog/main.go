package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

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

	// 文档框
	textEdit := widgets.NewQTextEdit(widget)
	app.SetCentralWidget(textEdit)

	menubar := app.MenuBar()
	fileMenu := menubar.AddMenu2("&File")
	// 创建一个事件和一个特定的图标和一个“退出”的标签。
	openFile := fileMenu.AddAction2(gui.NewQIcon5("images/app.ico"), "&Open")
	// 快捷键，自定义
	openFile.SetShortcut(gui.NewQKeySequence2("Ctrl+O", gui.QKeySequence__NativeText))
	// 提示语
	openFile.SetStatusTip("Open new File")
	// 事件触发
	openFile.ConnectTriggered(func(checked bool) {
		filename := widgets.NewQFileDialog(app, core.Qt__Widget).
			GetOpenFileName(
				app,
				"Open file",                        // 对话框的标题
				"/home",                            // 指定起始目录
				"*.go;;*.jpg;;*.png;;All Files(*)", // 指定过滤器，以此限制用户可以选择的文件类型。多个过滤条件用;;隔开
				"Files(*.txt)",
				widgets.QFileDialog__DontUseNativeDialog,
			)

		if filename != "" {
			// 读取了选择的文件并将文件内容显示到了TextEdit控件。
			if contents, err := ioutil.ReadFile(filename); err == nil {
				//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
				result := strings.Replace(string(contents), "\n", "", 1)
				textEdit.SetText(result)
			}
		} else {
			fmt.Println(reflect.TypeOf(filename), filename)
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
