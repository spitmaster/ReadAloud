package biz

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"read-aloud/internal/view"
)

type readAloudApp struct {
	app         fyne.App
	mainWindow  fyne.Window
	otherWindow []fyne.Window
}

func (a *readAloudApp) Start() {
	a.mainWindow.SetContent(widget.NewLabel("Hello World!"))
	a.mainWindow.Show()
	for _, ow := range a.otherWindow {
		ow.Show()
	}
	a.app.Run()
}

func (a *readAloudApp) Stop() {
	for _, w := range a.otherWindow {
		w.Close()
	}
	a.mainWindow.Close()
	a.app.Quit()
}

func (a *readAloudApp) OpenNewWindow(title string) fyne.Window {
	window := a.app.NewWindow(title)
	window.Resize(fyne.NewSize(500, 500))
	window.Show()
	a.otherWindow = append(a.otherWindow, window)
	return window
}

func FirstCreate() view.ReadAloud {
	//创建程序主体
	fyneApp := app.New()
	mainWindow := CreateMainWindow(fyneApp)
	readAloudApp := &readAloudApp{
		fyneApp,
		mainWindow,
		[]fyne.Window{},
	}
	return readAloudApp
}
