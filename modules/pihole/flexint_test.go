package pihole

import (
	"encoding/json"
	"testing"
)

func TestFlexInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    FlexInt
		wantErr bool
	}{
		{"numeric value", `42`, FlexInt(42), false},
		{"string value", `"42"`, FlexInt(42), false},
		{"negative numeric", `-5`, FlexInt(-5), false},
		{"non-numeric string", `"abc"`, FlexInt(0), true},
		{"invalid json", `{`, FlexInt(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fi FlexInt

			err := json.Unmarshal([]byte(tt.input), &fi)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("UnmarshalJSON(%s) expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("UnmarshalJSON(%s) unexpected error: %v", tt.input, err)
			}

			if fi != tt.want {
				t.Errorf("UnmarshalJSON(%s) = %v, want %v", tt.input, fi, tt.want)
			}
		})
	}
}
