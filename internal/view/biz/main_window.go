package biz

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"strconv"
)

func CreateMainWindow(fyneApp fyne.App) fyne.Window {
	mainWindow := fyneApp.NewWindow("Read Aloud")
	mainWindow.Resize(fyne.NewSize(500, 500))
	mainWindowContent(mainWindow)
	mainWindow.SetMaster() //设置主窗口, 这个窗口被关闭则程序关闭
	return mainWindow
}

func mainWindowContent(w fyne.Window) {
	//w.SetContent(widget.NewLabel("Hello World!"))

	text1 := widget.NewRichTextWithText(randomText())
	text2 := canvas.NewText(randomText(), color.Black)
	fmt.Println("第三个")
	text3 := widget.NewLabel("第三个")
	text4 := canvas.NewText(randomText(), color.Black)
	text5 := canvas.NewText(randomText(), color.Black)
	text6 := canvas.NewText(randomText(), color.Black)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3, text4, text5, text6)
	w.SetContent(grid)
}

func randomText() string {
	var bt bytes.Buffer
	bt.WriteString("                ")
	for i := 0; i < 10; i++ {
		bt.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	bt.WriteString("                ")
	return bt.String()
}
