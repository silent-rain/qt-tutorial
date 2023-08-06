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

	//创建了一个QLabel控件并为它设置了一个初始音量图像。
	label := widgets.NewQLabel(widget, core.Qt__Widget)
	//label.SetPixmap(gui.QPixmap_FromImage2(gui.NewQImage9("images/audio.png", ""), 0))

	// 图片缩放无效
	// pix := gui.NewQPixmap3("images/audio.png", "",core.Qt__AutoColor)
	// pix.Scaled(label.Size(),core.Qt__IgnoreAspectRatio, core.Qt__FastTransformation)
	// pix.Scaled2(10,10,core.Qt__IgnoreAspectRatio, core.Qt__FastTransformation)
	// label.SetPixmap(pix)

	label.SetScaledContents(true)
	label.SetPixmap(gui.NewQPixmap3("images/audio.png", "", core.Qt__AutoColor))

	label.SetGeometry2(160, 40, 50, 50)

	// 创建一个水平滑块
	sld := widgets.NewQSlider2(core.Qt__Horizontal, widget)
	sld.SetFocusPolicy(core.Qt__NoFocus)
	sld.SetGeometry2(30, 40, 100, 30)
	sld.SetValue(0)

	// 我们根据滑动条的值来设置标签的图像。
	// 当滑动条的值为0时我们为标签设置audio.ico图像。
	sld.ConnectValueChanged(func(value int) {
		if value == 0 {
			label.SetPixmap(gui.NewQPixmap3("images/audio.png", "", core.Qt__AutoColor))
		} else if value > 0 && value <= 30 {
			label.SetPixmap(gui.NewQPixmap3("images/audio-low.png", "", core.Qt__AutoColor))
		} else if value > 30 && value < 80 {
			label.SetPixmap(gui.NewQPixmap3("images/audio-mid.png", "", core.Qt__AutoColor))
		} else {
			label.SetPixmap(gui.NewQPixmap3("images/audio-high.png", "", core.Qt__AutoColor))
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
