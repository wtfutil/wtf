package bittrex

import (
	"bytes"
	"fmt"
	"text/template"
)

func (widget *Widget) display() {

	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	if !ok {
		return widget.CommonSettings().Title, errorText, true
	}

	list := &widget.summaryList
	str := ""

	for _, baseCurrency := range list.items {
		str += fmt.Sprintf(
			" [%s]%s[%s] (%s)\n\n",
			widget.settings.colors.base.displayName,
			baseCurrency.displayName,
			widget.settings.colors.base.name,
			baseCurrency.name,
		)

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

			err := strTemplate.Execute(writer, map[string]string{
				"nameColor":      widget.settings.colors.market.name,
				"fieldColor":     widget.settings.colors.market.field,
				"valueColor":     widget.settings.colors.market.value,
				"mName":          marketCurrency.name,
				"High":           marketCurrency.High,
				"Low":            marketCurrency.Low,
				"Last":           marketCurrency.Last,
				"Volume":         marketCurrency.Volume,
				"OpenBuyOrders":  marketCurrency.OpenBuyOrders,
				"OpenSellOrders": marketCurrency.OpenSellOrders,
			})

			if err != nil {
				str = err.Error()
			} else {
				str += writer.String() + "\n"
			}
		}

	}

	return widget.CommonSettings().Title, str, true

}

func formatableText(key, value string) string {
	return fmt.Sprintf("[{{.fieldColor}}]%12s: [{{.valueColor}}]{{.%s}}\n", key, value)
}
