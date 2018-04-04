package main

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/jira"
	"github.com/senorprogrammer/wtf/opsgenie"
	"github.com/senorprogrammer/wtf/security"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/weather"
)

func loadConfig() *config.Config {
	cfg, err := config.ParseYamlFile("./config.yml")
	if err != nil {
		panic(err)
	}

	return cfg
}

func refresher(stat *status.Widget, app *tview.Application) {
	tick := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			app.Draw()
		case <-quit:
			tick.Stop()
			return
		}
	}
}

func main() {
	cfg := loadConfig()

	fmt.Printf("%v\n", cfg)

	bamboo := bamboohr.NewWidget()
	bamboo.Refresh()

	cal := gcal.NewWidget()
	cal.Refresh()

	git := git.NewWidget()
	git.Refresh()

	jira := jira.NewWidget()
	jira.Refresh()

	opsgenie := opsgenie.NewWidget()
	opsgenie.Refresh()

	sec := security.NewWidget()
	sec.Refresh()

	stat := status.NewWidget()
	stat.Refresh()

	weather := weather.NewWidget()
	weather.Refresh()

	grid := tview.NewGrid()
	grid.SetRows(9, 9, 9, 9, 9, 6, 3) // How _high_ the row is, in terminal rows
	grid.SetColumns(40, 40)           // How _wide_ the column is, in terminal columns
	grid.SetBorder(false)

	grid.AddItem(bamboo.View, 0, 0, 2, 1, 0, 0, false)
	grid.AddItem(cal.View, 2, 1, 4, 1, 0, 0, false)
	grid.AddItem(git.View, 0, 2, 3, 1, 0, 0, false)
	grid.AddItem(weather.View, 0, 1, 1, 1, 0, 0, false)
	grid.AddItem(sec.View, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(opsgenie.View, 2, 0, 3, 1, 0, 0, false)
	grid.AddItem(jira.View, 5, 0, 1, 1, 0, 0, false)
	grid.AddItem(stat.View, 6, 0, 3, 3, 0, 0, false)

	app := tview.NewApplication()

	// Loop in a routine to redraw the screen
	go refresher(stat, app)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
