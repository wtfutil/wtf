package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
)

func main() {
	app := tview.NewApplication()

	grid := tview.NewGrid()
	grid.SetRows(10, 40)    // 10 high, 40 high
	grid.SetColumns(40, 40) // 40 wide, 40 wide
	grid.SetBorder(false)

	grid.AddItem(bambooView(), 0, 0, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func bambooView() tview.Primitive {
	items := bamboohr.Fetch()

	bamboo := tview.NewTextView()
	bamboo.SetBorder(true)
	bamboo.SetDynamicColors(true)
	bamboo.SetTitle(" 🐨 Away ")

	data := ""
	for _, item := range items {
		str := fmt.Sprintf("[green]%s[white]\n%s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
		data = data + str
	}

	fmt.Fprintf(bamboo, "%s", data)

	return bamboo
}
