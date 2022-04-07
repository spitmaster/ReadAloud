package view

import (
	"fyne.io/fyne/v2"
)

type ReadAloud interface {
	//启动应用
	Start()
	//停止应用
	Stop()
	//打开一个新窗口
	OpenNewWindow(title string) fyne.Window
}
