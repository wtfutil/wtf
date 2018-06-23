package bittrex

import (
	"bytes"
	"fmt"
	"text/template"
)

func (widget *Widget) display() {
	if ok == false {
		widget.View.SetText(errorText)
		return
	}

	widget.View.SetText(summaryText(&widget.summaryList, &widget.TextColors))
}

func summaryText(list *summaryList, colors *TextColors) string {
	str := ""

	for _, baseCurrency := range list.items {
		str += fmt.Sprintf(" [%s]%s[%s] (%s)\n\n", colors.base.displayName, baseCurrency.displayName, colors.base.name, baseCurrency.name)

		resultTemplate := template.New("bittrex")

		for _, marketCurrency := range baseCurrency.markets {
			writer := new(bytes.Buffer)

			strTemplate, _ := resultTemplate.Parse(
				"  [{{.nameColor}}]{{.mName}}\n" +
					formatableText("High", "High") +
					formatableText("Low", "Low") +
					formatableText("Last", "Last") +
					formatableText("Volume", "Volume") +
					"\n" +
					formatableText("Open Buy", "OpenBuyOrders") +
					formatableText("Open Sell", "OpenSellOrders"),
			)

			strTemplate.Execute(writer, map[string]string{
				"nameColor":      colors.market.name,
				"fieldColor":     colors.market.field,
				"valueColor":     colors.market.value,
				"mName":          marketCurrency.name,
				"High":           marketCurrency.High,
				"Low":            marketCurrency.Low,
				"Last":           marketCurrency.Last,
				"Volume":         marketCurrency.Volume,
				"OpenBuyOrders":  marketCurrency.OpenBuyOrders,
				"OpenSellOrders": marketCurrency.OpenSellOrders,
			})

			str += writer.String() + "\n"
		}

	}

	return str

}

func formatableText(key, value string) string {
	return fmt.Sprintf("[{{.fieldColor}}]%12s: [{{.valueColor}}]{{.%s}}\n", key, value)
}
