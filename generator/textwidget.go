//go:build ignore

// This package takes care of generates for empty widgets. Each generator is named after the
// type of widget it generate, so textwidget.go will generate the skeleton for a new TextWidget
// using the textwidget.tpl template.
// The TextWidget generator needs one environment variable, called WTF_WIDGET_NAME, which will
// be the name of the TextWidget it generates. If the variable hasn't been set, the generator
// will use "NewTextWidget". On Linux and macOS the command can be run as
// 'WTF_WIDGET_NAME=MyNewWidget go generate -run=text'.
package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const (
	defaultWidgetName = "NewTextWidget"
	widgetMaker       = "app/widget_maker.go"
)

func main() {
	widgetName, present := os.LookupEnv("WTF_WIDGET_NAME")
	if !present {
		widgetName = defaultWidgetName
	}

	data := struct {
		Name string
	}{
		widgetName,
	}

	createModuleDirectory(data)

	generateWidgetFile(data)
	generateSettingsFile(data)
	fmt.Println("Don't forget to register your module in file", widgetMaker)
}

/* -------------------- Unexported Functions -------------------- */

func createModuleDirectory(data struct{ Name string }) {
	err := os.MkdirAll(strings.ToLower(fmt.Sprintf("modules/%s", data.Name)), os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func generateWidgetFile(data struct{ Name string }) {
	tpl, _ := template.New("textwidget.tpl").Funcs(template.FuncMap{
		"Lower": strings.ToLower,
	}).ParseFiles("generator/textwidget.tpl")

	out, err := os.Create(fmt.Sprintf("modules/%s/widget.go", strings.ToLower(data.Name)))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer out.Close()

	tpl.Execute(out, data)
}

func generateSettingsFile(data struct{ Name string }) {
	tpl, _ := template.New("settings.tpl").Funcs(template.FuncMap{
		"Lower": strings.ToLower,
	}).ParseFiles("generator/settings.tpl")

	out, err := os.Create(fmt.Sprintf("modules/%s/settings.go", strings.ToLower(data.Name)))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer out.Close()

	tpl.Execute(out, data)
}
