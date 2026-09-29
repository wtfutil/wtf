package urlcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkValid(t *testing.T, got *urlResult) {
	assert.True(t, got.IsValid)
	assert.Less(t, got.ResultCode, 500)
	assert.Len(t, got.ResultMessage, 0)
}

func checkInvalid(t *testing.T, got *urlResult) {
	assert.False(t, got.IsValid)
	assert.GreaterOrEqual(t, got.ResultCode, 500)
	assert.Greater(t, len(got.ResultMessage), 0)
}

func Test_newUrlResult(t *testing.T) {
	type args struct {
		urlString string
	}
	type checks func(t *testing.T, res *urlResult)

	tests := []struct {
		name   string
		args   args
		checks checks
	}{
		{"good", args{"http://www.go.dev"}, checkValid},
		{"good_with_page", args{"https://go.dev/doc/install"}, checkValid},
		{"good_with_args", args{"https://mysite.com?var=1"}, checkValid},
		{"no_url", args{""}, checkInvalid},
		{"no_escape_chars", args{"http://not\nurl.com?var=1"}, checkInvalid},
		{"no_protocol", args{"go.dev"}, checkInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.checks != nil {
				tt.checks(t, newUrlResult(tt.args.urlString))
			}
		})
	}
}

func Test_urlResult_DisplayUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "no_auth",
			url:  "https://example.com/path?var=1",
			want: "https://example.com/path?var=1",
		},
		{
			name: "username_only",
			url:  "https://user@example.com/path",
			want: "https://user@example.com/path",
		},
		{
			name: "username_and_password",
			url:  "https://user:secret@example.com/path",
			want: "https://user:xxxxx@example.com/path",
		},
		{
			name: "username_and_empty_password",
			url:  "https://user:@example.com/path",
			want: "https://user:xxxxx@example.com/path",
		},
		{
			name: "encoded_credentials",
			url:  "https://user%40example.com:s%40cret@example.com/path",
			want: "https://user%40example.com:xxxxx@example.com/path",
		},
		{
			name: "invalid_url",
			url:  "http://not\nurl.com?var=1",
			want: "http://not\nurl.com?var=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := urlResult{Url: tt.url}

			assert.Equal(t, tt.want, got.DisplayUrl())
			assert.Equal(t, tt.url, got.Url)
		})
	}
}
