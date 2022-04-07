package biz

import "fyne.io/fyne/v2"

func CreateMainWindow(fyneApp fyne.App) fyne.Window {
	mainWindow := fyneApp.NewWindow("Read Aloud")
	mainWindow.Resize(fyne.NewSize(500, 500))
	mainWindow.SetMaster() //设置主窗口, 这个窗口被关闭则程序关闭
	return mainWindow
}
