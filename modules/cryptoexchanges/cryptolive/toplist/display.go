package toplist

import "fmt"

func (widget *Widget) display() {
	str := ""
	for _, fromCurrency := range widget.list.items {
		str += fmt.Sprintf(
			"[%s]%s [%s](%s)\n",
			widget.colors.from.displayName,
			fromCurrency.displayName,
			widget.colors.from.name,
			fromCurrency.name,
		)
		str += makeToListText(fromCurrency.to, widget.colors)
	}

	widget.Result = str
}

func makeToListText(toList []*tCurrency, colors textColors) string {
	str := ""
	for _, toCurrency := range toList {
		str += makeToText(toCurrency, colors)
	}

	return str
}

func makeToText(toCurrency *tCurrency, colors textColors) string {
	str := ""
	str += fmt.Sprintf("  [%s]%s\n", colors.to.name, toCurrency.name)
	for _, info := range toCurrency.info {
		str += makeInfoText(info, colors)
		str += "\n\n"
	}
	return str
}

func makeInfoText(info tInfo, colors textColors) string {
	return fmt.Sprintf(
		"    [%s]Exchange: [%s]%s\n",
		colors.to.field,
		colors.to.value,
		info.exchange,
	) +
		fmt.Sprintf(
			"    [%s]Volume(24h): [%s]%f-[%s]%f",
			colors.to.field,
			colors.to.value,
			info.volume24h,
			colors.to.value,
			info.volume24hTo,
		)
}
