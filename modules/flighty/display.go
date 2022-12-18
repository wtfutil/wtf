package flighty

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	content := "Flight tracker!"
	setWrap := false

	return title, content, setWrap
}
