package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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
