package price

type list struct {
	items []*fromCurrency
}

type fromCurrency struct {
	name        string
	displayName string
	to          []*toCurrency
}

type toCurrency struct {
	name  string
	price float32
}

type cResponse map[string]float32

/* -------------------- Unexported Functions -------------------- */

func (l *list) addItem(name string, displayName string, to []*toCurrency) {
	l.items = append(l.items, &fromCurrency{
		name:        name,
		displayName: displayName,
		to:          to,
	})
}
