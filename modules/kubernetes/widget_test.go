package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateTitle(t *testing.T) {
	type fields struct {
		title      string
		namespaces []string
	}
	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "No Namespaces",
			fields: fields{
				namespaces: []string{},
			},
			want: "Kube",
		},
		{
			name: "One Namespace",
			fields: fields{
				namespaces: []string{"some-namespace"},
			},
			want: "Kube - Namespace: some-namespace",
		},
		{
			name: "Multiple Namespaces",
			fields: fields{
				namespaces: []string{"ns1", "ns2"},
			},
			want: `Kube - Namespaces: ["ns1" "ns2"]`,
		},
		{
			name: "Explicit Title Set",
			fields: fields{
				namespaces: []string{},
				title:      "Test Explicit Title",
			},
			want: "Test Explicit Title",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				title:      tt.fields.title,
				namespaces: tt.fields.namespaces,
			}
			assert.Equal(t, tt.want, widget.generateTitle())
		})
	}
}
