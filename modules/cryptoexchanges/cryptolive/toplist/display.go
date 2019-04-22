package toplist

import "fmt"

func (widget *Widget) display() {
	str := ""
	for _, fromCurrency := range widget.list.items {
		str += fmt.Sprintf(
			"[%s]%s [%s](%s)\n",
			widget.settings.colors.from.displayName,
			fromCurrency.displayName,
			widget.settings.colors.from.name,
			fromCurrency.name,
		)
		str += widget.makeToListText(fromCurrency.to)
	}

	widget.Result = str
}

func (widget *Widget) makeToListText(toList []*tCurrency) string {
	str := ""
	for _, toCurrency := range toList {
		str += widget.makeToText(toCurrency)
	}

	return str
}

func (widget *Widget) makeToText(toCurrency *tCurrency) string {
	str := ""
	str += fmt.Sprintf(
		"  [%s]%s\n",
		widget.settings.colors.to.name,
		toCurrency.name,
	)

	for _, info := range toCurrency.info {
		str += widget.makeInfoText(info)
		str += "\n\n"
	}
	return str
}

func (widget *Widget) makeInfoText(info tInfo) string {
	return fmt.Sprintf(
		"    [%s]Exchange: [%s]%s\n",
		widget.settings.colors.top.to.field,
		widget.settings.colors.top.to.value,
		info.exchange,
	) +
		fmt.Sprintf(
			"    [%s]Volume(24h): [%s]%f-[%s]%f",
			widget.settings.colors.top.to.field,
			widget.settings.colors.top.to.value,
			info.volume24h,
			widget.settings.colors.top.to.value,
			info.volume24hTo,
		)
}
