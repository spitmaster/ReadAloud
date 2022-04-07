package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	app        fyne.App
	mainWindow fyne.Window
}

func CreateView() *View {
	//创建程序主体
	a := app.New()
	//创建一个主窗口
	w := a.NewWindow("Read Aloud")
	v := View{
		a,
		w,
	}
	return &v
}

func (v *View) Start() error {
	v.mainWindow.SetContent(widget.NewLabel("Hello World!"))
	v.mainWindow.ShowAndRun()
	return nil
}
