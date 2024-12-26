package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"

	//"fyne.io/fyne/v2/internal/theme"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Body struct {
	Name string
	Age  int
}

type Iot struct {
	Id     int64
	Time   time.Time
	Value  float64
	Offset float64
	CID    string
}

func (me *Iot) Set(c int) interface{} {
	switch c {
	case 0:
		return me.Id
	case 1:
		return me.Time
	case 2:
		return me.Value
	case 3:
		return me.Offset
	case 4:
		return me.CID
	}
	return ""
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

func updateTime(clock *widget.Label) {
	fmt := time.Now().Format("Time:03:04:05")
	clock.SetText(fmt)
}

func main() {

	//var win fyne.Window
	app := app.NewWithID(" 你好！")

	icon, err := fyne.LoadResourceFromPath("Icon-n2ew.png")
	if err != nil {
		icon = theme.HomeIcon()
	}

	app.SetIcon(icon)

	win := MainWindow(app)

	SystemTray(app, win)

	win.ShowAndRun()

}

func SystemTray(app fyne.App, win fyne.Window) {
	if desk, ok := app.(desktop.App); ok {

		show := fyne.NewMenuItem("显示窗口", func() {
			//app.Quit()
			win.Show()
		})
		show.Icon = theme.SearchIcon()

		menu := fyne.NewMenu("SysTray", show)

		desk.SetSystemTrayIcon(app.Icon())

		//desk.SetSystemTrayIcon()
		desk.SetSystemTrayMenu(menu)
		//win.SetMainMenu(fyne.NewMainMenu(menu))

	}

}

func MainWindow(app fyne.App) fyne.Window {

	win := app.NewWindow("new window")
	win.Resize(fyne.NewSize(800, 600))

	win.SetCloseIntercept(func() {
		win.Hide()
	})

	//top.Theme().Color()

	//list := widget.NewTable(nil, nil, nil)
	list := widget.NewTableWithHeaders(func() (rows int, cols int) {
		return len(bodys), 2
	}, func() fyne.CanvasObject {

		return widget.NewLabel("asdfasdf")
	}, func(tci widget.TableCellID, co fyne.CanvasObject) {
		co.(*widget.Label).SetText(fmt.Sprint(bodys[tci.Row].value(tci.Col)))
	})

	//list.

	columns := []string{"姓名", "年龄"}

	list.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row < 0 {
			template.(*widget.Label).SetText(columns[id.Col])
			//template.Resize(fyne.NewSize(100.0, 100.0))
		} else if id.Col < 0 {
			template.(*widget.Label).SetText(strconv.Itoa(id.Row + 1))
		}
	}

	clock := widget.NewLabel("")
	top := Toolbar(win)
	context := container.NewBorder(top, clock, nil, nil, list)
	win.SetContent(context)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
			//fmt.Println(win.Content().Visible())

		}

	}()

	/*
		win.SetOnClosed(func() {
			fmt.Println("exit")
			//return false
		})
	*/

	return win

}

func Toolbar(win fyne.Window) *widget.Toolbar {
	open := widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		//dialog.NewError(errors.New("发现错误，请清除！"), win).Show()
		//dialog.NewConfirm("提示", "提示信息", nil, win).Show()

		dialog.NewInformation("提示", "提示信息-------------------", win).Show()

	})
	new := widget.NewToolbarAction(theme.FolderNewIcon(), func() {
		boy := Body{Name: "ASD", Age: 10}
		bodys = append(bodys, &boy)

	})

	//close := widget.NewButton("ASD", func() {})

	exit := widget.NewToolbarAction(theme.WindowCloseIcon(), func() {
		//fmt.Println()
		dialog.NewConfirm("提示", "你是否要退出！", func(b bool) {
			fmt.Println(b)
			if b {
				win.Hide()
			}
		}, win).Show()
	})

	//open.
	//open.OnActivated = f
	return widget.NewToolbar(open, new, widget.NewToolbarSpacer(), exit)
	//top.Theme().Color()

}
