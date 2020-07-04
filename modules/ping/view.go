package ping

import "fmt"

type results []string

func getPingResult(widget *Widget, t target, logging bool) string {
	o := t.raw

	if t.ips == nil {
		if widget.settings.useEmoji {
			o = fmt.Sprintf("🔴 %s (lookup failed)", t.raw)
		} else {
			o = fmt.Sprintf("[red]%s (lookup failed)", t.raw)
		}
	} else {
		if widget.settings.showIP && t.raw != t.ips[0].String() {
			o = fmt.Sprintf("%s (%s)", t.raw, t.ips[0])
		}

		switch checkTarget(t.ips[0], widget.settings.pingTimeout, logging) {
		case msgFail:
			if widget.settings.useEmoji {
				o = fmt.Sprintf("🔴 %s", o)
			} else {
				o = fmt.Sprintf("[red]%s", o)
			}
		case msgSuccess:
			if widget.settings.useEmoji {
				o = fmt.Sprintf("🟢 %s", o)
			} else {
				o = fmt.Sprintf("[green]%s", o)
			}
		}
	}

	return o
}

func (r results) Len() int {
	return len(r)
}
func (r results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r results) Less(i, j int) bool {
	return r[j] < r[i]
}
