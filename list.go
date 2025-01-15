package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func List() *widget.Table {
	//iots, _ := fetchLatestIotData()

	list := widget.NewTableWithHeaders(func() (rows int, cols int) {
		return len(IotData), 5
	}, func() fyne.CanvasObject {
		return widget.NewLabel("                  ")
	}, func(tci widget.TableCellID, co fyne.CanvasObject) {
		co.(*widget.Label).SetText(fmt.Sprint(IotData[tci.Row].ValueOf(tci.Col)))
	})

	columns := []string{"Id", "时间", "数值", "偏移", "CID"}
	list.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row < 0 {
			template.(*widget.Label).SetText(columns[id.Col])
			//template.Resize(fyne.NewSize(100.0, 100.0))
		} else if id.Col < 0 {
			template.(*widget.Label).SetText(strconv.Itoa(id.Row + 1))
		}
	}

	// list.Bind(iots)

	list.SetColumnWidth(1, 200)
	//list.UpdateCell
	//list.Refresh()
	//list.UpdateHeader()

	list.OnSelected = func(id widget.TableCellID) {
		fmt.Println("Selected", id)
	}

	return list
}
