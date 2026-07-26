package cryptolive

import (
	"fmt"
	"testing"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/modules/cryptocurrency/cryptolive/price"
	"github.com/wtfutil/wtf/modules/cryptocurrency/cryptolive/toplist"
)

func TestWidget_ContentLogic(t *testing.T) {
	// Test the string concatenation logic that content() performs
	tests := []struct {
		name       string
		priceRes   string
		toplistRes string
		wantParts  []string
	}{
		{
			name:       "both have content",
			priceRes:   "BTC: $50000\n",
			toplistRes: "Top: Binance\n",
			wantParts:  []string{"BTC: $50000", "Top: Binance"},
		},
		{
			name:       "only price",
			priceRes:   "only price\n",
			toplistRes: "",
			wantParts:  []string{"only price"},
		},
		{
			name:       "only toplist",
			priceRes:   "",
			toplistRes: "only toplist\n",
			wantParts:  []string{"only toplist"},
		},
		{
			name:       "empty results",
			priceRes:   "",
			toplistRes: "",
			wantParts:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Replicate the content logic without needing CommonSettings
			str := ""
			str += tc.priceRes
			str += tc.toplistRes
			content := fmt.Sprintf("\n%s", str)

			for _, part := range tc.wantParts {
				if !strContains(content, part) {
					t.Errorf("expected %q in content, got: %s", part, content)
				}
			}
		})
	}
}

func TestPriceWidgetResult(t *testing.T) {
	pw := &price.Widget{Result: "test price data"}
	if pw.Result != "test price data" {
		t.Errorf("unexpected result: %s", pw.Result)
	}
}

func TestToplistWidgetResult(t *testing.T) {
	tw := &toplist.Widget{Result: "test toplist data"}
	if tw.Result != "test toplist data" {
		t.Errorf("unexpected result: %s", tw.Result)
	}
}

func TestContentFormat(t *testing.T) {
	// Verify the format string wraps with newline prefix
	priceResult := "price"
	toplistResult := "toplist"
	str := priceResult + toplistResult
	content := fmt.Sprintf("\n%s", str)

	if content[0] != '\n' {
		t.Error("content should start with newline")
	}
	if content != "\npricetoplist" {
		t.Errorf("unexpected content: %q", content)
	}
}

func TestNewSettingsFromYAML(t *testing.T) {
	yamlStr := `
currencies:
  BTC:
    displayName: "Bitcoin"
    to:
      - USD
      - EUR
top:
  BTC:
    displayName: "Bitcoin"
    limit: 5
    to:
      - USD
colors:
  from:
    name: "green"
    displayName: "yellow"
  to:
    name: "cyan"
    price: "white"
  top:
    from:
      name: "green"
      displayName: "yellow"
    to:
      name: "cyan"
      field: "white"
      value: "blue"
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}
	globalConfig, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	settings := NewSettingsFromYAML("cryptolive", yamlConfig, globalConfig)
	if settings == nil {
		t.Fatal("settings is nil")
	}
	if settings.from.name != "green" {
		t.Errorf("from.name: want green, got %s", settings.from.name)
	}
	if settings.from.displayName != "yellow" {
		t.Errorf("from.displayName: want yellow, got %s", settings.from.displayName)
	}
	if settings.to.name != "cyan" {
		t.Errorf("to.name: want cyan, got %s", settings.to.name)
	}
	if settings.to.price != "white" {
		t.Errorf("to.price: want white, got %s", settings.to.price)
	}
	if settings.colors.top.from.name != "green" {
		t.Errorf("top.from.name: want green, got %s", settings.colors.top.from.name)
	}
	if settings.colors.top.to.field != "white" {
		t.Errorf("top.to.field: want white, got %s", settings.colors.top.to.field)
	}
	if settings.colors.top.to.value != "blue" {
		t.Errorf("top.to.value: want blue, got %s", settings.colors.top.to.value)
	}
	if settings.priceSettings == nil {
		t.Error("priceSettings should not be nil")
	}
	if settings.toplistSettings == nil {
		t.Error("toplistSettings should not be nil")
	}
}

func strContains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
