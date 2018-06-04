package bittrex

import (
	"bytes"
	"fmt"
	"text/template"
)

func (widget *Widget) display() {
	if ok == false {
		widget.View.SetText(fmt.Sprintf("%s", errorText))
		return
	}

	str := ""
	str += summaryText(&widget.summaryList, &widget.TextColors)

	widget.View.SetText(fmt.Sprintf("%s", str))
}

func summaryText(list *summaryList, colors *TextColors) string {
	str := ""

	for _, baseCurrency := range list.items {
		str += fmt.Sprintf("[%s]%s[%s](%s):\n", colors.base.displayName, baseCurrency.displayName, colors.base.name, baseCurrency.name)

		resultTemplate := template.New("bittrex")

		for _, marketCurrency := range baseCurrency.markets {
			writer := new(bytes.Buffer)
			strTemplate, _ := resultTemplate.Parse(
				"\t[{{.nameColor}}]{{.mName}}\n" +
					formatableText("High", "High") +
					formatableText("Low", "Low") +
					formatableText("Last", "Last") +
					formatableText("Volume", "Volume") +
					formatableText("OpenSellOrders", "OpenSellOrders") +
					formatableText("OpenBuyOrders", "OpenBuyOrders"),
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
				"OpenSellOrders": marketCurrency.OpenSellOrders,
				"OpenBuyOrders":  marketCurrency.OpenBuyOrders,
			})

			str += writer.String()
		}

	}

	return str

}

func formatableText(key, value string) string {
	return fmt.Sprintf("\t\t[{{.fieldColor}}]%s: [{{.valueColor}}]{{.%s}}\n", key, value)
}
