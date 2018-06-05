package toplist

import "fmt"

func (widget *Widget) display() {
	str := ""
	for _, fromCurrency := range widget.list.items {
		str += fmt.Sprintf("%s (%s)\n", fromCurrency.displayName, fromCurrency.name)
		for _, toCurrency := range fromCurrency.to {
			str += fmt.Sprintf("  %s\n", toCurrency.name)
			for _, info := range toCurrency.info {
				str += makeInfoRow(info)
				str += "\n\n"
			}
		}
	}

	widget.Result = str
}

func makeInfoRow(info tInfo) string {
	return fmt.Sprintf("    Exchange: %s\n", info.exchange) + fmt.Sprintf("    Volume(24h): %f-%f", info.volume24h, info.volume24hTo)
}
