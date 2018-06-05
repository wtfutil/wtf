package toplist

import "fmt"

func (widget *Widget) display() {
	str := ""
	for _, fromCurrency := range widget.list.items {
		str += fmt.Sprintf("%s (%s)\n", fromCurrency.displayName, fromCurrency.name)
		str += makeToListText(fromCurrency.to)
	}

	widget.Result = str
}

func makeToListText(toList []*tCurrency) string {
	str := ""
	for _, toCurrency := range toList {
		str += makeToText(toCurrency)
	}

	return str
}

func makeToText(toCurrency *tCurrency) string {
	str := ""
	str += fmt.Sprintf("  %s\n", toCurrency.name)
	for _, info := range toCurrency.info {
		str += makeInfoText(info)
		str += "\n\n"
	}
	return str
}

func makeInfoText(info tInfo) string {
	return fmt.Sprintf("    Exchange: %s\n", info.exchange) + fmt.Sprintf("    Volume(24h): %f-%f", info.volume24h, info.volume24hTo)
}
