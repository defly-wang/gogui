package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
	type Body struct {
		Name string
		Age  int
	}

	var bodys = []*Body{
		{Name: "name", Age: 1},
		{Name: "name1", Age: 2},
	}

	func (me *Body) value(c int) interface{} {
		switch c {
		case 0:
			return me.Name
		case 1:
			return me.Age
		}
		return ""
	}
*/
func MainWindow(app fyne.App) fyne.Window {

	win := app.NewWindow("new window")
	//a := fyne.CurrentApp().Driver().ScreenSize()
	win.Resize(fyne.NewSize(800, 600))

	clock := widget.NewLabel("")

	//fetchLatestIotData()
	list := List()
	top := Toolbar(win, list)
	context := container.NewBorder(top, clock, nil, nil, list)
	win.SetContent(context)

	updateList(list)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}

	}()

	go func() {
		for range time.Tick(time.Second * 5) {
			updateList(list)
		}
	}()

	return win

}

func updateList(list *widget.Table) {
	fetchLatestIotData()
	list.Refresh()
}

func updateTime(clock *widget.Label) {
	fmt := time.Now().Format("Time:03:04:05")
	clock.SetText(fmt)
}
