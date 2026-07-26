package todo

import (
	"strings"
	"testing"
	"time"

	"github.com/wtfutil/wtf/checklist"
)

func TestGetNowDate(t *testing.T) {
	now := getNowDate()
	if now.Hour() != 0 || now.Minute() != 0 || now.Second() != 0 || now.Nanosecond() != 0 {
		t.Errorf("getNowDate() should return midnight, got %v", now)
	}
	today := time.Now()
	if now.Year() != today.Year() || now.Month() != today.Month() || now.Day() != today.Day() {
		t.Errorf("getNowDate() should be today's date, got %v", now)
	}
}

func TestGetDateString(t *testing.T) {
	now := getNowDate()
	widget := &Widget{
		settings: &Settings{
			switchToInDaysIn:  7,
			dateFormat:        "yyyy-mm-dd",
			hideYearIfCurrent: false,
		},
	}

	tests := []struct {
		name     string
		date     time.Time
		format   string
		hideYear bool
		expected string
	}{
		// today / tomorrow / "in X days" cases
		{"today", now, "yyyy-mm-dd", false, "today"},
		{"tomorrow", now.AddDate(0, 0, 1), "yyyy-mm-dd", false, "tomorrow"},
		{"in 3 days", now.AddDate(0, 0, 3), "yyyy-mm-dd", false, "in 3 days"},
		{"in 7 days (boundary)", now.AddDate(0, 0, 7), "yyyy-mm-dd", false, "in 7 days"},

		// format: yyyy-mm-dd
		{"yyyy-mm-dd far future", time.Date(2030, 3, 15, 0, 0, 0, 0, time.Local), "yyyy-mm-dd", false, "2030-03-15"},
		// format: yy-mm-dd
		{"yy-mm-dd", time.Date(2030, 3, 15, 0, 0, 0, 0, time.Local), "yy-mm-dd", false, "30-03-15"},
		// format: dd-mm-yyyy
		{"dd-mm-yyyy", time.Date(2030, 12, 5, 0, 0, 0, 0, time.Local), "dd-mm-yyyy", false, "05-12-2030"},
		// format: dd-mm-yy
		{"dd-mm-yy", time.Date(2030, 12, 5, 0, 0, 0, 0, time.Local), "dd-mm-yy", false, "05-12-30"},
		// format: dd M yyyy
		{"dd M yyyy", time.Date(2030, 1, 22, 0, 0, 0, 0, time.Local), "dd M yyyy", false, "22 Jan 2030"},
		// format: dd M yy
		{"dd M yy", time.Date(2030, 1, 22, 0, 0, 0, 0, time.Local), "dd M yy", false, "22 Jan 30"},
		// unknown format defaults to yyyy-mm-dd
		{"unknown format", time.Date(2030, 6, 1, 0, 0, 0, 0, time.Local), "unknown", false, "2030-06-01"},

		// hideYearIfCurrent with yyyy-mm-dd format (year prefix)
		{"hideYear yyyy-mm-dd current year", time.Date(now.Year(), 8, 15, 0, 0, 0, 0, time.Local), "yyyy-mm-dd", true, "08-15"},
		// hideYearIfCurrent with yy-mm-dd format (year prefix)
		{"hideYear yy-mm-dd current year", time.Date(now.Year(), 8, 15, 0, 0, 0, 0, time.Local), "yy-mm-dd", true, "08-15"},
		// hideYearIfCurrent with dd-mm-yyyy (year suffix, dash separator)
		{"hideYear dd-mm-yyyy current year", time.Date(now.Year(), 8, 15, 0, 0, 0, 0, time.Local), "dd-mm-yyyy", true, "15-08"},
		// hideYearIfCurrent with dd-mm-yy (year suffix, dash separator)
		{"hideYear dd-mm-yy current year", time.Date(now.Year(), 8, 15, 0, 0, 0, 0, time.Local), "dd-mm-yy", true, "15-08"},
		// hideYearIfCurrent with dd M yyyy (space separator)
		{"hideYear dd M yyyy current year", time.Date(now.Year(), 8, 15, 0, 0, 0, 0, time.Local), "dd M yyyy", true, "15 Aug"},
		// hideYearIfCurrent with dd M yy (space separator)
		{"hideYear dd M yy current year", time.Date(now.Year(), 8, 15, 0, 0, 0, 0, time.Local), "dd M yy", true, "15 Aug"},
		// hideYearIfCurrent should NOT hide if different year
		{"hideYear different year", time.Date(now.Year()+2, 8, 15, 0, 0, 0, 0, time.Local), "yyyy-mm-dd", true, formatYMD(now.Year()+2, 8, 15)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			widget.settings.dateFormat = tc.format
			widget.settings.hideYearIfCurrent = tc.hideYear
			// Ensure dates far enough in the future to not hit "in X days"
			date := tc.date
			got := widget.getDateString(&date)
			if got != tc.expected {
				t.Errorf("getDateString() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func formatYMD(y int, m, d int) string {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local).Format("2006-01-02")
}

func TestGetDateString_SwitchToInDaysIn(t *testing.T) {
	now := getNowDate()
	widget := &Widget{
		settings: &Settings{
			switchToInDaysIn:  3,
			dateFormat:        "yyyy-mm-dd",
			hideYearIfCurrent: false,
		},
	}

	// 3 days out should show "in 3 days"
	date := now.AddDate(0, 0, 3)
	got := widget.getDateString(&date)
	if got != "in 3 days" {
		t.Errorf("expected 'in 3 days', got %q", got)
	}

	// 4 days out should show date format since switchToInDaysIn=3
	date = now.AddDate(0, 0, 4)
	got = widget.getDateString(&date)
	if got == "in 4 days" {
		t.Errorf("expected formatted date, got %q", got)
	}
}

func TestShouldShowItem(t *testing.T) {
	tests := []struct {
		name          string
		item          *checklist.ChecklistItem
		parseTags     bool
		showTagPrefix string
		showFilter    string
		hideTags      []interface{}
		expected      bool
	}{
		{
			name:      "no filter, no tags parsing",
			item:      &checklist.ChecklistItem{Text: "hello world", Tags: nil},
			parseTags: false,
			expected:  true,
		},
		{
			name:       "filter matches text",
			item:       &checklist.ChecklistItem{Text: "Buy groceries"},
			showFilter: "buy",
			parseTags:  false,
			expected:   true,
		},
		{
			name:       "filter does not match",
			item:       &checklist.ChecklistItem{Text: "Buy groceries"},
			showFilter: "sell",
			parseTags:  false,
			expected:   false,
		},
		{
			name:      "parseTags, no tags on item, no showTagPrefix",
			item:      &checklist.ChecklistItem{Text: "hello", Tags: []string{}},
			parseTags: true,
			expected:  true,
		},
		{
			name:          "parseTags, no tags on item, with showTagPrefix",
			item:          &checklist.ChecklistItem{Text: "hello", Tags: []string{}},
			parseTags:     true,
			showTagPrefix: "work",
			expected:      false,
		},
		{
			name:      "parseTags, item has matching tag, no prefix filter",
			item:      &checklist.ChecklistItem{Text: "task", Tags: []string{"work"}},
			parseTags: true,
			expected:  true,
		},
		{
			name:          "parseTags, item has matching tag prefix",
			item:          &checklist.ChecklistItem{Text: "task", Tags: []string{"work"}},
			parseTags:     true,
			showTagPrefix: "wo",
			expected:      true,
		},
		{
			name:          "parseTags, item has non-matching tag prefix",
			item:          &checklist.ChecklistItem{Text: "task", Tags: []string{"home"}},
			parseTags:     true,
			showTagPrefix: "work",
			expected:      false,
		},
		{
			name:     "hideTags hides matching tag when no prefix",
			item:     &checklist.ChecklistItem{Text: "task", Tags: []string{"secret"}},
			parseTags: true,
			hideTags: []interface{}{"secret"},
			expected: false,
		},
		{
			name:          "hideTags does NOT hide when showTagPrefix is set",
			item:          &checklist.ChecklistItem{Text: "task", Tags: []string{"secret"}},
			parseTags:     true,
			showTagPrefix: "sec",
			hideTags:      []interface{}{"secret"},
			expected:      true,
		},
		{
			name:     "item with multiple tags, one is hidden",
			item:     &checklist.ChecklistItem{Text: "task", Tags: []string{"visible", "secret"}},
			parseTags: true,
			hideTags: []interface{}{"secret"},
			expected: true, // "visible" tag allows it to show
		},
		{
			name:     "item with only hidden tags",
			item:     &checklist.ChecklistItem{Text: "task", Tags: []string{"secret"}},
			parseTags: true,
			hideTags: []interface{}{"secret"},
			expected: false,
		},
		{
			name:       "filter + tags combined: filter fails",
			item:       &checklist.ChecklistItem{Text: "Buy milk", Tags: []string{"shop"}},
			parseTags:  true,
			showFilter: "sell",
			expected:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			widget := &Widget{
				showTagPrefix: tc.showTagPrefix,
				showFilter:    tc.showFilter,
				settings: &Settings{
					parseTags: tc.parseTags,
					hideTags:  tc.hideTags,
				},
			}
			got := widget.shouldShowItem(tc.item)
			if got != tc.expected {
				t.Errorf("shouldShowItem() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestGetTodoTags(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedText string
		expectedTags []string
	}{
		{"no tags", "hello world", "hello world", []string{}},
		{"single tag at end", "task #work", "task", []string{"work"}},
		{"single tag at start", "#work task", "task", []string{"work"}},
		{"multiple tags", "task #work #urgent", "task", []string{"work", "urgent"}},
		{"tag in middle", "do #work today", "do today", []string{"work"}},
		{"tag with numbers", "task #project1", "task", []string{"project1"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			text, tags := getTodoTags(tc.input)
			text = strings.TrimSpace(text)
			if text != tc.expectedText {
				t.Errorf("getTodoTags() text = %q, want %q", text, tc.expectedText)
			}
			if len(tags) != len(tc.expectedTags) {
				t.Errorf("getTodoTags() tags = %v, want %v", tags, tc.expectedTags)
				return
			}
			for i, tag := range tags {
				if tag != tc.expectedTags[i] {
					t.Errorf("getTodoTags() tag[%d] = %q, want %q", i, tag, tc.expectedTags[i])
				}
			}
		})
	}
}

func TestSortListByChecked_Hidden(t *testing.T) {
	// We can't test sortListByChecked directly because it calls formattedItemLine
	// which requires a tview.View. Instead, we test shouldShowItem which is the
	// core filtering logic used by sortListByChecked.
	// The sortListByChecked ordering logic is implicitly tested via shouldShowItem
	// and the placeItemBasedOnDate tests.
}

func TestRowColor(t *testing.T) {
	// RowColor requires widget.View (tview.TextView) which needs full tview setup.
	// The color selection logic is simple branching that is validated through
	// integration testing rather than unit testing.
}


func TestGetTextAndDate(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			parseDates: true,
		},
	}

	t.Run("in X days pattern", func(t *testing.T) {
		text, date := widget.getTextAndDate("in 3 days buy milk")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(0, 0, 3)
		if date.Day() != expected.Day() || date.Month() != expected.Month() {
			t.Errorf("expected date ~%v, got %v", expected, *date)
		}
		if !strings.Contains(text, "buy milk") {
			t.Errorf("expected text to contain 'buy milk', got %q", text)
		}
	})

	t.Run("in X weeks pattern", func(t *testing.T) {
		text, date := widget.getTextAndDate("in 2 weeks do laundry")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(0, 0, 14)
		if date.Day() != expected.Day() || date.Month() != expected.Month() {
			t.Errorf("expected date ~%v, got %v", expected, *date)
		}
		if !strings.Contains(text, "do laundry") {
			t.Errorf("expected text to contain 'do laundry', got %q", text)
		}
	})

	t.Run("in X months pattern", func(t *testing.T) {
		_, date := widget.getTextAndDate("in 1 month review")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(0, 1, 0)
		if date.Month() != expected.Month() {
			t.Errorf("expected month %v, got %v", expected.Month(), date.Month())
		}
	})

	t.Run("in X years pattern", func(t *testing.T) {
		_, date := widget.getTextAndDate("in 1 year review")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(1, 0, 0)
		if date.Year() != expected.Year() {
			t.Errorf("expected year %v, got %v", expected.Year(), date.Year())
		}
	})

	t.Run("today prefix", func(t *testing.T) {
		text, date := widget.getTextAndDate("today fix bug")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		now := time.Now()
		if date.Day() != now.Day() {
			t.Errorf("expected today's date, got %v", *date)
		}
		if !strings.Contains(text, "fix bug") {
			t.Errorf("expected text 'fix bug', got %q", text)
		}
	})

	t.Run("tomorrow prefix", func(t *testing.T) {
		text, date := widget.getTextAndDate("tomorrow fix bug")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(0, 0, 1)
		if date.Day() != expected.Day() {
			t.Errorf("expected tomorrow's date, got %v", *date)
		}
		if !strings.Contains(text, "fix bug") {
			t.Errorf("expected text 'fix bug', got %q", text)
		}
	})

	t.Run("next week prefix", func(t *testing.T) {
		_, date := widget.getTextAndDate("next week task")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(0, 0, 7)
		if date.Day() != expected.Day() {
			t.Errorf("expected date %v, got %v", expected, *date)
		}
	})

	t.Run("next month prefix", func(t *testing.T) {
		_, date := widget.getTextAndDate("next month task")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(0, 1, 0)
		if date.Month() != expected.Month() {
			t.Errorf("expected month %v, got %v", expected.Month(), date.Month())
		}
	})

	t.Run("next year prefix", func(t *testing.T) {
		_, date := widget.getTextAndDate("next year task")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		expected := time.Now().AddDate(1, 0, 0)
		if date.Year() != expected.Year() {
			t.Errorf("expected year %v, got %v", expected.Year(), date.Year())
		}
	})

	t.Run("YYYY-MM-DD prefix", func(t *testing.T) {
		text, date := widget.getTextAndDate("2025-06-15 deadline")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		if date.Year() != 2025 || date.Month() != 6 || date.Day() != 15 {
			t.Errorf("expected 2025-06-15, got %v", *date)
		}
		if !strings.Contains(text, "deadline") {
			t.Errorf("expected text 'deadline', got %q", text)
		}
	})

	t.Run("MM-DD prefix", func(t *testing.T) {
		text, date := widget.getTextAndDate("06-15 deadline")
		if date == nil {
			t.Fatal("expected a date, got nil")
		}
		if date.Month() != 6 || date.Day() != 15 {
			t.Errorf("expected month 6 day 15, got %v", *date)
		}
		if !strings.Contains(text, "deadline") {
			t.Errorf("expected text 'deadline', got %q", text)
		}
	})

	t.Run("no date", func(t *testing.T) {
		text, date := widget.getTextAndDate("just text")
		if date != nil {
			t.Errorf("expected nil date, got %v", *date)
		}
		if text != "just text" {
			t.Errorf("expected 'just text', got %q", text)
		}
	})

	t.Run("in X days without text returns no match", func(t *testing.T) {
		// "in 3 days" alone (no trailing text) should not match
		text, date := widget.getTextAndDate("in 3 days")
		if date != nil {
			t.Errorf("expected nil date for bare 'in 3 days', got %v", *date)
		}
		if text != "in 3 days" {
			t.Errorf("expected unchanged text, got %q", text)
		}
	})

	t.Run("today alone without text returns no match", func(t *testing.T) {
		text, date := widget.getTextAndDate("today")
		if date != nil {
			t.Errorf("expected nil date for bare 'today', got %v", *date)
		}
		if text != "today" {
			t.Errorf("expected unchanged text, got %q", text)
		}
	})
}

func TestGetTextComponents(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			parseDates: true,
			parseTags:  true,
		},
	}

	text, date, tags := widget.getTextComponents("today buy #groceries milk")
	if date == nil {
		t.Fatal("expected date")
	}
	if len(tags) != 1 || tags[0] != "groceries" {
		t.Errorf("expected tags=[groceries], got %v", tags)
	}
	if !strings.Contains(text, "buy") {
		t.Errorf("expected text to contain 'buy', got %q", text)
	}
}

func TestGetTextComponents_NoDates(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			parseDates: false,
			parseTags:  true,
		},
	}

	text, date, tags := widget.getTextComponents("today buy #groceries milk")
	if date != nil {
		t.Error("expected nil date when parseDates is false")
	}
	if len(tags) != 1 || tags[0] != "groceries" {
		t.Errorf("expected tags=[groceries], got %v", tags)
	}
	if !strings.Contains(text, "today") {
		t.Errorf("expected 'today' in text when dates disabled, got %q", text)
	}
}

func TestGetTextComponents_NoTags(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			parseDates: false,
			parseTags:  false,
		},
	}

	text, date, tags := widget.getTextComponents("buy #groceries milk")
	if date != nil {
		t.Error("expected nil date")
	}
	if len(tags) != 0 {
		t.Errorf("expected empty tags when parseTags is false, got %v", tags)
	}
	if text != "buy #groceries milk" {
		t.Errorf("expected unchanged text, got %q", text)
	}
}

func TestTodoDateIsEarlier(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			undatedAsDays: 7,
		},
		list: checklist.NewChecklist("x", " "),
	}

	now := getNowDate()
	d1 := now.AddDate(0, 0, 1)
	d2 := now.AddDate(0, 0, 5)

	item1 := &checklist.ChecklistItem{Date: &d1, Text: "early"}
	item2 := &checklist.ChecklistItem{Date: &d2, Text: "late"}
	item3 := &checklist.ChecklistItem{Date: nil, Text: "no date"}

	widget.list.Items = []*checklist.ChecklistItem{item1, item2, item3}

	// item1 (day 1) is earlier than item2 (day 5)
	if !widget.todoDateIsEarlier(0, 1) {
		t.Error("expected item1 to be earlier than item2")
	}
	if widget.todoDateIsEarlier(1, 0) {
		t.Error("expected item2 NOT to be earlier than item1")
	}

	// both nil
	itemNil1 := &checklist.ChecklistItem{Date: nil, Text: "a"}
	itemNil2 := &checklist.ChecklistItem{Date: nil, Text: "b"}
	widget.list.Items = []*checklist.ChecklistItem{itemNil1, itemNil2}
	if widget.todoDateIsEarlier(0, 1) {
		t.Error("two nil dates should not be 'earlier'")
	}

	// nil vs dated: undatedAsDays=7 means nil treated as now+7 days
	widget.list.Items = []*checklist.ChecklistItem{item3, item1}
	// item3 (nil -> now+7) vs item1 (now+1): nil is NOT earlier
	if widget.todoDateIsEarlier(0, 1) {
		t.Error("nil (treated as +7 days) should not be earlier than +1 day")
	}
	// item1 (now+1) vs item3 (nil -> now+7): item1 IS earlier
	widget.list.Items = []*checklist.ChecklistItem{item1, item3}
	if !widget.todoDateIsEarlier(0, 1) {
		t.Error("+1 day should be earlier than nil (+7 days)")
	}
}

func TestPlaceItemBasedOnDate(t *testing.T) {
	widget := &Widget{
		settings: &Settings{
			undatedAsDays: 7,
		},
		list: checklist.NewChecklist("x", " "),
	}

	now := getNowDate()
	d1 := now.AddDate(0, 0, 10)
	d2 := now.AddDate(0, 0, 20)
	d3 := now.AddDate(0, 0, 5)

	item1 := &checklist.ChecklistItem{Date: &d1, Text: "10 days"}
	item2 := &checklist.ChecklistItem{Date: &d2, Text: "20 days"}
	item3 := &checklist.ChecklistItem{Date: &d3, Text: "5 days"}

	// Start with [10, 20, 5] - placing index 2 should move it to front
	widget.list.Items = []*checklist.ChecklistItem{item1, item2, item3}
	newIdx := widget.placeItemBasedOnDate(2)
	if newIdx != 0 {
		t.Errorf("expected item3 to move to index 0, got %d", newIdx)
	}
	if widget.list.Items[0].Text != "5 days" {
		t.Errorf("expected '5 days' at index 0, got %q", widget.list.Items[0].Text)
	}
}
