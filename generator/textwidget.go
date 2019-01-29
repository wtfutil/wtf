// +build ignore

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

var (
	widgetName string
)

const (
	defaultWidgetName = "NewTextWidget"
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

	tpl, _ := template.New("textwidget.tpl").Funcs(template.FuncMap{
		"Lower": strings.ToLower,
		"Title": strings.Title,
	}).ParseFiles("generator/textwidget.tpl")

	err := os.Mkdir(strings.ToLower(widgetName), os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	out, err := os.Create(fmt.Sprintf("%s/widget.go", strings.ToLower(widgetName)))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer out.Close()

	tpl.Execute(out, data)
}
